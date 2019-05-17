package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/ziyoumeng/golang/helper/idl"
	"github.com/ziyoumeng/golang/helper/model"
	"reflect"
)

type A struct {
	Name string
	Age  int
	Son  struct {
		Name   string
		School string
	}
	TPtr  *int
	Slice []int
	//One   int
	Int8  int8
	Int64 int64
}

type B struct {
	Name string `json:"na"`
	Age  int
	Son  struct {
		Name   string
		School string
	}
	TPtr  *int
	Slice []int
	Int64 int8
	Int8  int64
	//One   string
}

func main() {
	/*	var i = 10
		a := A{
			Name: "lf",
			Age:  120,
			Son: struct {
				Name   string
				School string
			}{Name: "lfson", School: "sch"},
			TPtr:  &i,
			Slice: []int{1, 2, 3},
			Int64:64,
			Int8:8,
		}
		var b B*/

	a := idl.Order{
		Rule: idl.ExchangeRule{
			Items: []*idl.VoucherItem{
				&idl.VoucherItem{
					PropID:  1,
					PropNum: 1,
				},
			},
		},
		UID:1,
	}
	var b = new(model.ExchangeOrder)
	if err := CopyData(b, a); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n %+v \n", a, b)

}


//CopyData 使用限制，不能包含指针
func CopyData(dst interface{}, src interface{}) error {
	return copyData(reflect.ValueOf(dst), reflect.ValueOf(src))
}

func getElemValue(value reflect.Value) (reflect.Value, error) {
	if !value.IsValid() {
		return reflect.Value{}, errors.New("value can't be zero Value")
	}
	if value.Type().Kind() == reflect.Ptr || value.Type().Kind() == reflect.Interface {
		value = value.Elem()
	} else {
		return value, nil
	}
	return getElemValue(value)
}

func isStruct(data reflect.Value) bool {
	return data.Type().Kind() == reflect.Struct
}

func copyData(dstValue, srcValue reflect.Value) error {
	var err error
	dstValue, err = getElemValue(dstValue)
	if err != nil {
		return errors.Wrap(err, "dstValue getElemValue fail")
	}

	srcValue, err = getElemValue(srcValue)
	if err != nil {
		return errors.Wrap(err, "srcValue getElemValue fail")
	}

	if !dstValue.CanSet() {
		return errors.New("dstValue can't set")
	}
	dstKind, srcKind := dstValue.Type().Kind(), srcValue.Type().Kind()
	if !isCompile(dstKind, srcKind) {
		return errors.Errorf("dstKind %s is not compile with srcKind %s)", dstKind, srcKind)
	}
	srcKind = getCompileKind(srcKind)
	switch srcKind {
	case reflect.Int:
		if dstValue.OverflowInt(srcValue.Int()){
			return errors.Errorf("dstValue overflowInt with value = %d",srcValue.Int())
		}
	case reflect.Uint:
		if dstValue.OverflowUint(srcValue.Uint()){
			return errors.Errorf("dstValue overflowUint with value = %d",srcValue.Int())
		}
	case reflect.Float64:
		if dstValue.OverflowFloat(srcValue.Float()){
			return errors.Errorf("dstValue overflowFlow with value = %d",srcValue.Int())
		}
	case reflect.Struct:
		if err := copyStruct(dstValue, srcValue); err != nil {
			return errors.Wrap(err, "structCopy fail")
		}
	case reflect.Slice:
		if srcValue.Len() == 0 {
			return nil
		}
		dstValue.Set(reflect.MakeSlice(dstValue.Type(), srcValue.Len(), srcValue.Cap()))
		for i := 0; i < srcValue.Len(); i++ {
			fmt.Println(dstValue.Index(i), srcValue.Index(i))
			if err := copyData(dstValue.Index(i), srcValue.Index(i)); err != nil {
				return errors.Wrapf(err, "copy slice data fail with index = %d", i)
			}
		}
	default:
		dstValue.Set(srcValue)
	}
	return nil
}

func copyStruct(dstValue ,srcValue reflect.Value) error {
	for i := 0; i < srcValue.NumField(); i++ {
		srcFieldValue := srcValue.Field(i)
		srcFieldName := srcValue.Type().Field(i).Name
		if srcFieldName == "XXX_unrecognized" {
			continue
		}
		dstFieldValue := dstValue.FieldByName(srcFieldName)
		/*if dstFieldValue.Kind() == reflect.Ptr && dstFieldValue.IsNil() {
			tp := dstFieldValue.Type()
			dstFieldValue.Set(reflect.New(tp).Elem())
		}*/

		if err := copyData(dstFieldValue, srcFieldValue); err != nil {
			return errors.Wrapf(err, "field (name=%s) copyData fail", srcFieldName)
		}
	}
	return nil
}

func isCompile(dstKind, srcKind reflect.Kind) bool {
	if isInt(dstKind) && isInt(srcKind) {
		return true
	}
	if isUint(dstKind) && isUint(srcKind) {
		return true
	}

	if isFloat(dstKind) && isUint(srcKind){
		return true
	}

	return dstKind == srcKind
}

func isInt(k reflect.Kind) bool {
	return k >= reflect.Int && k <= reflect.Int64
}

func isUint(k reflect.Kind) bool {
	return k >= reflect.Uint && k <= reflect.Uint64
}

func isFloat(k reflect.Kind) bool{
	return k == reflect.Float64 || k== reflect.Float32
}
func getCompileKind(k reflect.Kind) reflect.Kind {
	if isInt(k) {
		return reflect.Int
	}
	if isUint(k) {
		return reflect.Uint
	}
	if isFloat(k){
		return reflect.Float64
	}
	return k
}

/*func typeMatch(dstField, srcField reflect.Value) error {
	kind := srcField.Type().Kind()
	//dstKind := dstField.Type().Kind()

	//if !isCompile(dstKind, kind) {
	//	return errors.Errorf("dest %s is not compile with src %s)", dstKind, kind)
	//}
	kind = getCompileKind(kind)
	switch kind {
	case reflect.String:
		dstField.SetString(srcField.String())
	case reflect.Int:
		dstField.SetInt(srcField.Int())
	case reflect.Uint:
		dstField.SetUint(srcField.Uint())
	case reflect.Struct:
		if err := structCopy(dstField, srcField); err != nil {
			return errors.Wrap(err, "structCopy fail")
		}
	case reflect.Ptr:
		if dstField.Kind() == reflect.Ptr && dstField.IsNil() {
			tp := dstField.Type()
			dstField.Set(reflect.New(tp).Elem())
			dstField = dstField.Elem()
		}

		if err := typeMatch(dstField, srcField.Elem()); err != nil {
			return errors.Wrap(err, "type match fail")
		}
	case reflect.Slice:
		if srcField.Len() == 0 {
			return nil
		}
		dstField.Set(reflect.MakeSlice(dstField.Type(), srcField.Len(), srcField.Cap()))
		for i := 0; i < srcField.Len(); i++ {
			fmt.Println(dstField.Index(i), srcField.Index(i))
			if err := typeMatch(dstField.Index(i), srcField.Index(i)); err != nil {
				return errors.Wrap(err, "type match fail")
			}
		}
	default:
		return errors.Errorf("not support type %s", kind)
	}
	return nil
}*/

/*func test(c interface{}) {
	setEmptyPtr(reflect.ValueOf(c), reflect.TypeOf(c))
}

func setEmptyPtr(val reflect.Value, tp reflect.Type) (reflect.Value, bool) {
	if !val.IsValid() {
		fmt.Println("elem is zero Value")
		return reflect.Value{}, false
	}

	if  tp.Kind() == reflect.Ptr {//tp.Kind() == reflect.Interface
		if !val.Elem().CanSet() {
			fmt.Println("val can't set")
			return reflect.Value{}, false
		}
		if !val.Elem().IsValid(){
			val.Set(reflect.New(tp.Elem()))
		}
		return setEmptyPtr(val.Elem(), tp.Elem())
	}
	return reflect.Value{}, true
}*/
