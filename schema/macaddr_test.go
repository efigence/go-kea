package schema

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMacaddr_MarshalJSON(t *testing.T) {
	mac, err := ParseMAC("11:22:33:44:55:66")
	out, err2 := mac.MarshalJSON()

	Convey("mac type to json", t, func() {
		So(err, ShouldBeNil)
		So(err2, ShouldBeNil)
		So(string(out), ShouldEqual, `"11:22:33:44:55:66"`)
	})
}

func TestMacaddr_UnmarshalJSON(t *testing.T) {
	var mac Macaddr
	err := mac.UnmarshalJSON([]byte(`"22:33:44:55:66:77"`))
	Convey("mac json to type", t, func() {
		So(err, ShouldBeNil)
		So(mac.String(), ShouldEqual, `22:33:44:55:66:77`)

	})
}
