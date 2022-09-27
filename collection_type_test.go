package wlc

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoginTraceParam_Add(t *testing.T) {
	loginTraceParam := LoginTraceParam{
		Collections: make([]*LoginTrace, 0),
	}
	for i := 0; i < 128; i++ {
		loginTraceParam.Add(&LoginTrace{})
		Convey("TestLoginTraceParam_Add", t, func() {
			So(loginTraceParam.Collections[i].No, ShouldEqual, i+1)
		})
	}
}
