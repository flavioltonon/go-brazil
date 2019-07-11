package brazil_test

import (
	"testing"

	. "flavioltonon/go-brazil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseSUS(t *testing.T) {
	Convey("Given a string named s", t, func() {
		var s string

		Convey("If s is empty", func() {
			s = ""

			Convey("And the function ParseSUS is called using it as an argument", func() {
				sus, err := ParseSUS(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the SUS struct number should be empty", func() {
						So(sus.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is an invalid SUS number", func() {
			s = "213946989760001"

			Convey("And the function ParseSUS is called using it as an argument", func() {
				sus, err := ParseSUS(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the SUS struct number should be empty", func() {
						So(sus.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a valid SUS number", func() {
			s = "213946989760008"

			Convey("And the function ParseSUS is called using it as an argument", func() {
				sus, err := ParseSUS(s)

				Convey("It should not return an error", func() {
					So(err, ShouldEqual, nil)

					Convey("And the SUS struct number should exist", func() {
						So(sus.Number(false), ShouldEqual, "213946989760008")
						So(sus.Number(true), ShouldEqual, "213 9469 8976 0008")
					})
				})
			})
		})
	})
}

func TestRandomSUSNumber(t *testing.T) {
	Convey("Given the function RandomSUSNumber", t, func() {
		Convey("If its mask argument equals true", func() {
			number := RandomSUSNumber(true)

			Convey("It should return a valid SUS number", func() {
				sus, err := ParseSUS(number)

				So(sus.Number(false), ShouldNotEqual, "")
				So(sus.Number(true), ShouldNotEqual, "")
				So(err, ShouldEqual, nil)
			})
		})

		Convey("If its mask argument equals false", func() {
			number := RandomSUSNumber(false)

			Convey("It should return a valid SUS number", func() {
				sus, err := ParseSUS(number)

				So(sus.Number(false), ShouldNotEqual, "")
				So(sus.Number(true), ShouldNotEqual, "")
				So(err, ShouldEqual, nil)
			})
		})
	})
}
