package brazil_test

import (
	"testing"

	. "github.com/flavioltonon/go-brazil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParsePIS(t *testing.T) {
	Convey("Given a string named s", t, func() {
		var s string

		Convey("If s is empty", func() {
			s = ""

			Convey("And the function ParsePIS is called using it as an argument", func() {
				pis, err := ParsePIS(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the PIS struct number should be empty", func() {
						So(pis.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a false-positive PIS number", func() {
			s = "00000000000"

			Convey("And the function ParsePIS is called using it as an argument", func() {
				pis, err := ParsePIS(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the PIS struct number should be empty", func() {
						So(pis.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a PIS number with an invalid digit", func() {
			s = "31873846876"

			Convey("And the function ParsePIS is called using it as an argument", func() {
				pis, err := ParsePIS(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the PIS struct number should be empty", func() {
						So(pis.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a valid PIS number", func() {
			s = "31873846877"

			Convey("And the function ParsePIS is called using it as an argument", func() {
				pis, err := ParsePIS(s)

				Convey("It should not return an error", func() {
					So(err, ShouldEqual, nil)

					Convey("And the PIS struct number should exist", func() {
						So(pis.Number(false), ShouldEqual, "31873846877")
						So(pis.Number(true), ShouldEqual, "318.73846.87-7")
					})
				})
			})
		})
	})
}

func TestRandomPISNumber(t *testing.T) {
	Convey("Given the function RandomPISNumber", t, func() {
		Convey("If its mask argument equals true", func() {
			number := RandomPISNumber(true)

			Convey("It should return a valid PIS number", func() {
				pis, err := ParsePIS(number)

				So(pis.Number(false), ShouldNotEqual, "")
				So(pis.Number(true), ShouldNotEqual, "")
				So(err, ShouldEqual, nil)
			})
		})

		Convey("If its mask argument equals false", func() {
			number := RandomPISNumber(false)

			Convey("It should return a valid PIS number", func() {
				pis, err := ParsePIS(number)

				So(pis.Number(false), ShouldNotEqual, "")
				So(pis.Number(true), ShouldNotEqual, "")
				So(err, ShouldEqual, nil)
			})
		})
	})
}
