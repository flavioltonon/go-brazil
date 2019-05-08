package brazil_test

import (
	"testing"

	. "github.com/flavioltonon/go-brazil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestValidate(t *testing.T) {
	Convey("Given a documentType named s, a string named n and a bool named m", t, func() {
		var s DocumentType
		var n string
		var m bool

		Convey("If s is empty", func() {
			s = ""

			Convey("And the function Validate is called using it as an argument", func() {
				number, err := Validate(s, n, m)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the number should be empty", func() {
						So(number, ShouldEqual, "")
					})
				})
			})
		})
	})
}
