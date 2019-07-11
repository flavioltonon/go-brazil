package brazil_test

import (
	"testing"

	. "github.com/flavioltonon/go-brazil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseMobile(t *testing.T) {
	Convey("Given a string named s", t, func() {
		var s string

		Convey("If s has less than 13 numbers", func() {
			s = "123456789"

			Convey("And the function ParseMobile is called using it as an argument", func() {
				mobile, err := ParseMobile(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the mobile struct numbers should be empty", func() {
						So(mobile.CountryCode(false), ShouldEqual, "")
						So(mobile.AreaCode(false), ShouldEqual, "")
						So(mobile.Number(false), ShouldEqual, "")
						So(mobile.FullNumber(false), ShouldEqual, "")
						So(mobile.CountryCode(true), ShouldEqual, "")
						So(mobile.AreaCode(true), ShouldEqual, "")
						So(mobile.Number(true), ShouldEqual, "")
						So(mobile.FullNumber(true), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s contains an invalid country code", func() {
			s = "5411987654321"

			Convey("And the function ParseMobile is called using it as an argument", func() {
				mobile, err := ParseMobile(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the mobile struct numbers should be empty", func() {
						So(mobile.CountryCode(false), ShouldEqual, "")
						So(mobile.AreaCode(false), ShouldEqual, "")
						So(mobile.Number(false), ShouldEqual, "")
						So(mobile.FullNumber(false), ShouldEqual, "")
						So(mobile.CountryCode(true), ShouldEqual, "")
						So(mobile.AreaCode(true), ShouldEqual, "")
						So(mobile.Number(true), ShouldEqual, "")
						So(mobile.FullNumber(true), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s contains an invalid area code", func() {
			s = "5501987654321"

			Convey("And the function ParseMobile is called using it as an argument", func() {
				mobile, err := ParseMobile(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the mobile struct numbers should be empty", func() {
						So(mobile.CountryCode(false), ShouldEqual, "")
						So(mobile.AreaCode(false), ShouldEqual, "")
						So(mobile.Number(false), ShouldEqual, "")
						So(mobile.FullNumber(false), ShouldEqual, "")
						So(mobile.CountryCode(true), ShouldEqual, "")
						So(mobile.AreaCode(true), ShouldEqual, "")
						So(mobile.Number(true), ShouldEqual, "")
						So(mobile.FullNumber(true), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s contains an invalid mobile number", func() {
			s = "5511287654321"

			Convey("And the function ParseMobile is called using it as an argument", func() {
				mobile, err := ParseMobile(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the mobile struct numbers should be empty", func() {
						So(mobile.CountryCode(false), ShouldEqual, "")
						So(mobile.AreaCode(false), ShouldEqual, "")
						So(mobile.Number(false), ShouldEqual, "")
						So(mobile.FullNumber(false), ShouldEqual, "")
						So(mobile.CountryCode(true), ShouldEqual, "")
						So(mobile.AreaCode(true), ShouldEqual, "")
						So(mobile.Number(true), ShouldEqual, "")
						So(mobile.FullNumber(true), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a valid mobile number", func() {
			s = "5511987654321"

			Convey("And the function ParseMobile is called using it as an argument", func() {
				mobile, err := ParseMobile(s)

				Convey("It should not return an error", func() {
					So(err, ShouldEqual, nil)

					Convey("And the mobile struct numbers should exist", func() {
						So(mobile.CountryCode(false), ShouldEqual, "55")
						So(mobile.AreaCode(false), ShouldEqual, "11")
						So(mobile.Number(false), ShouldEqual, "987654321")
						So(mobile.FullNumber(false), ShouldEqual, "5511987654321")

						So(mobile.CountryCode(true), ShouldEqual, "+55")
						So(mobile.AreaCode(true), ShouldEqual, "(11)")
						So(mobile.Number(true), ShouldEqual, "98765-4321")
						So(mobile.FullNumber(true), ShouldEqual, "+55(11)98765-4321")
					})
				})
			})
		})
	})
}

func TestRandomMobileFullNumber(t *testing.T) {
	Convey("Given the function RandomMobileFullNumber", t, func() {
		Convey("If its mask argument equals true", func() {
			number := RandomMobileFullNumber(true)

			Convey("It should return a valid mobile number", func() {
				mobile, err := ParseMobile(number)

				So(mobile.CountryCode(false), ShouldNotEqual, "")
				So(mobile.AreaCode(false), ShouldNotEqual, "")
				So(mobile.Number(false), ShouldNotEqual, "")
				So(mobile.FullNumber(false), ShouldNotEqual, "")

				So(mobile.CountryCode(true), ShouldNotEqual, "")
				So(mobile.AreaCode(true), ShouldNotEqual, "")
				So(mobile.Number(true), ShouldNotEqual, "")
				So(mobile.FullNumber(true), ShouldNotEqual, "")

				So(err, ShouldEqual, nil)
			})
		})

		Convey("If its mask argument equals false", func() {
			number := RandomMobileFullNumber(false)

			Convey("It should return a valid mobile number", func() {
				mobile, err := ParseMobile(number)

				So(mobile.CountryCode(false), ShouldNotEqual, "")
				So(mobile.AreaCode(false), ShouldNotEqual, "")
				So(mobile.Number(false), ShouldNotEqual, "")
				So(mobile.FullNumber(false), ShouldNotEqual, "")

				So(mobile.CountryCode(true), ShouldNotEqual, "")
				So(mobile.AreaCode(true), ShouldNotEqual, "")
				So(mobile.Number(true), ShouldNotEqual, "")
				So(mobile.FullNumber(true), ShouldNotEqual, "")

				So(err, ShouldEqual, nil)
			})
		})
	})
}
