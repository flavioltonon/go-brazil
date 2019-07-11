package brazil_test

import (
	"testing"

	. "flavioltonon/go-brazil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseCertidao(t *testing.T) {
	Convey("Given a string named s", t, func() {
		var s string

		Convey("If s is empty", func() {
			s = ""

			Convey("And the function ParseCertidao is called using it as an argument", func() {
				certidao, err := ParseCertidao(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the Certidao struct number should be empty", func() {
						So(certidao.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a Certidao number with an invalid year", func() {
			s = "104539 01 55 2050 1 00012 021 0000123-21"

			Convey("And the function ParseCertidao is called using it as an argument", func() {
				certidao, err := ParseCertidao(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the Certidao struct number should be empty", func() {
						So(certidao.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a Certidao number with an invalid first digit", func() {
			s = "104539 01 55 2013 1 00012 021 0000123-31"

			Convey("And the function ParseCertidao is called using it as an argument", func() {
				certidao, err := ParseCertidao(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the Certidao struct number should be empty", func() {
						So(certidao.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a Certidao number with an invalid second digit", func() {
			s = "104539 01 55 2013 1 00012 021 0000123-27"

			Convey("And the function ParseCertidao is called using it as an argument", func() {
				certidao, err := ParseCertidao(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the Certidao struct number should be empty", func() {
						So(certidao.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a valid Certidao number", func() {
			s = "104539 01 55 2013 1 00012 021 0000123-21"

			Convey("And the function ParseCertidao is called using it as an argument", func() {
				certidao, err := ParseCertidao(s)

				Convey("It should not return an error", func() {
					So(err, ShouldEqual, nil)

					Convey("And the Certidao struct number should exist", func() {
						So(certidao.Number(false), ShouldEqual, "10453901552013100012021000012321")
						So(certidao.Number(true), ShouldEqual, "104539 01 55 2013 1 00012 021 0000123-21")
					})
				})
			})
		})
	})
}

func TestRandomCertidaoNumber(t *testing.T) {
	Convey("Given the function RandomCertidaoNumber", t, func() {
		Convey("If its mask argument equals true", func() {
			number := RandomCertidaoNumber(true, CertidaoKindNone)

			Convey("It should return a valid Certidao number", func() {
				certidao, err := ParseCertidao(number)

				So(err, ShouldEqual, nil)
				So(certidao.Number(false), ShouldNotEqual, "")
				So(certidao.Number(true), ShouldNotEqual, "")
			})
		})

		Convey("If its mask argument equals false", func() {
			number := RandomCertidaoNumber(false, CertidaoKindNone)

			Convey("It should return a valid Certidao number", func() {
				certidao, err := ParseCertidao(number)

				So(err, ShouldEqual, nil)
				So(certidao.Number(false), ShouldNotEqual, "")
				So(certidao.Number(true), ShouldNotEqual, "")
			})
		})
	})
}
