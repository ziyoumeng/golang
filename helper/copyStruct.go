package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/ziyoumeng/golang/helper/idl"
	"github.com/ziyoumeng/golang/helper/model"
	"reflect"
)

func main() {
	var tmp int64=1
	a := idl.Order{
		Rule: idl.ExchangeRule{
			Items: []*idl.VoucherItem{
				&idl.VoucherItem{
					PropID:  1,
					PropNum: 1,
				},
			},
		},
		UID:       &tmp,
		PrizeName: "test",
	}

	var b = new(model.ExchangeOrder)
	if err := CopyData(b, a, true); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n %+v \n", *a.Rule.Items[0], b.Rule.Items[0])
	fmt.Println(*a.UID, *b.UID)
	*b.UID = 2
	fmt.Println(*a.UID, *b.UID)
	fmt.Println(a.PrizeName, b.PrizeName)


}

//CopyData dst必须指针类型;需要拷贝的字段dst和src必须同名;深拷貝;src的tag的json:"-"会跳过
func CopyData(dst interface{}, src interface{}, isOkWhenSrcFieldLack bool) error {
	return copyData(reflect.ValueOf(dst), reflect.ValueOf(src), isOkWhenSrcFieldLack)
}

func newPtr(value reflect.Value) error {
	if !value.IsValid() {
		return errors.New("value can't be zero Value")
	}

	if value.Type().Kind() == reflect.Ptr || value.Type().Kind() == reflect.Interface {
		if !value.Elem().IsValid() {
			if !value.CanSet() {
				return errors.New("value can't set")
			}
			value.Set(reflect.New(value.Type().Elem()))
		} else {
			newPtr(value.Elem())
		}
	}
	return nil
}

func getElemValue(value reflect.Value) (reflect.Value, error) {
	if !value.IsValid() {
		return reflect.Value{}, errors.New("value can't be zero Value")
	}
	if value.Type().Kind() == reflect.Ptr || value.Type().Kind() == reflect.Interface {
		return getElemValue(value.Elem())
	}
	return value, nil
}

func isStruct(data reflect.Value) bool {
	return data.Type().Kind() == reflect.Struct
}

func copyData(dstValue, srcValue reflect.Value, isOkWhenSrcFieldLack bool) error {
	err := newPtr(dstValue)
	if err != nil {
		return errors.Wrap(err, "dstValue newPtr fail")
	}

	dstValue, err = getElemValue(dstValue)
	if err != nil {
		return errors.Wrap(err, "dstValue getElemValue fail")
	}

	srcValue, err = getElemValue(srcValue)
	if err != nil {
		return errors.Wrap(err, "srcValue getElemValue fail")
	}

	dstKind, srcKind := dstValue.Type().Kind(), srcValue.Type().Kind()
	if !isCompile(dstKind, srcKind) {
		return errors.Errorf("dstKind %s is not compile with srcKind %s)", dstKind, srcKind)
	}

	if !dstValue.CanSet() {
		return errors.New("dstValue can't set")
	}

	dstKind = getCompileKind(dstKind)
	switch dstKind {
	case reflect.Int:
		if dstValue.OverflowInt(srcValue.Int()) {
			return errors.Errorf("dstValue overflowInt with value = %d", srcValue.Int())
		}
		dstValue.SetInt(srcValue.Int())
	case reflect.Uint:
		if dstValue.OverflowUint(srcValue.Uint()) {
			return errors.Errorf("dstValue overflowUint with value = %d", srcValue.Int())
		}
		dstValue.SetUint(srcValue.Uint())
	case reflect.Float64:
		if dstValue.OverflowFloat(srcValue.Float()) {
			return errors.Errorf("dstValue overflowFlow with value = %d", srcValue.Int())
		}
		dstValue.SetFloat(srcValue.Float())
	case reflect.Struct:
		if err := copyStruct(dstValue, srcValue, isOkWhenSrcFieldLack); err != nil {
			return errors.Wrap(err, "structCopy fail")
		}
	case reflect.Slice:
		if srcValue.Len() == 0 {
			return nil
		}
		dstValue.Set(reflect.MakeSlice(dstValue.Type(), srcValue.Len(), srcValue.Cap()))
		for i := 0; i < srcValue.Len(); i++ {
			if err := copyData(dstValue.Index(i), srcValue.Index(i), isOkWhenSrcFieldLack); err != nil {
				return errors.Wrapf(err, "copy slice data fail with index = %d", i)
			}
		}
	default:
		if !srcValue.Type().AssignableTo(dstValue.Type()) {
			return errors.New("srcValue can't AssignableTo dstValue")
		}
		dstValue.Set(srcValue)
	}
	return nil
}

func copyStruct(dstValue, srcValue reflect.Value, isOkWhenSrcFieldLack bool) error {
	for i := 0; i < dstValue.NumField(); i++ {
		dstField := dstValue.Type().Field(i)
		if dstField.Tag.Get("json") == "-" {
			continue
		}

		srcFieldValue := srcValue.FieldByName(dstField.Name)
		if !srcFieldValue.IsValid() {
			if isOkWhenSrcFieldLack {
				continue
			} else {
				return errors.Errorf("srcValue field %s not exist", dstField.Name)
			}

		}

		if err := copyData(dstValue.Field(i), srcFieldValue, isOkWhenSrcFieldLack); err != nil {
			return errors.Wrapf(err, "field (name=%s) copyData fail", dstField.Name)
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

	if isFloat(dstKind) && isUint(srcKind) {
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

func isFloat(k reflect.Kind) bool {
	return k == reflect.Float64 || k == reflect.Float32
}
func getCompileKind(k reflect.Kind) reflect.Kind {
	if isInt(k) {
		return reflect.Int
	}
	if isUint(k) {
		return reflect.Uint
	}
	if isFloat(k) {
		return reflect.Float64
	}
	return k
}
