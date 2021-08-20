package types

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestToV4(t *testing.T) {
	var i IPv4
	err := i.Scan(int64(2078987375))
	Convey("int to IP", t, func() {
		So(err, ShouldBeNil)
		So(i.String(), ShouldEqual, "123.234.212.111")
	})
	var i2 IPv4
	i2.UnmarshalText([]byte("123.234.212.111"))
	v, err := i2.Value()
	Convey("IP to int", t, func() {
		So(err, ShouldBeNil)
		So(v, ShouldEqual, 2078987375)
	})
}
