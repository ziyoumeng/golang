package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCopyData(t *testing.T) {
	Convey("build-in type",t, func() {
		var src = 1
		Convey("dst is referenced", func() {
			Convey("dst is pointer", func() {
				var dst *int
				err := CopyData(&dst, src, true)
				So(err,ShouldBeNil)
				So(*dst,ShouldEqual,src)
			})
			Convey("dst is not pointer", func() {
				var dst *int
				err := CopyData(&dst, src, true)
				So(err,ShouldBeNil)
				So(*dst,ShouldEqual,src)
			})
		})
		Convey("dst is not referenced",  func() {
			Convey("dst is pointer", func() {
				var dst *int
				err := CopyData(dst, src, true)
				So(err,ShouldNotBeNil)
			})
			Convey("dst is not pointer", func() {
				var dst *int
				err := CopyData(dst, src, true)
				So(err,ShouldNotBeNil)
			})
		})
	} )

}

func BenchmarkCopyData_simple(b *testing.B) {
	var dst int
	var src = 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CopyData(&dst, src, true)
	}
}

func Benchmark_assign(b *testing.B) {
	var dst int
	var src = 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst = src
	}
	src = dst + 1
}
