package brazil_test

import (
	"testing"

	. "github.com/flavioltonon/go-brazil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseCNPJ(t *testing.T) {
	Convey("Given a string named s", t, func() {
		var s string

		Convey("If s is empty", func() {
			s = ""

			Convey("And the function ParseCNPJ is called using it as an argument", func() {
				cnpj, err := ParseCNPJ(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the CNPJ struct number should be empty", func() {
						So(cnpj.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a CNPJ number with an invalid first digit", func() {
			s = "11222333000171"

			Convey("And the function ParseCNPJ is called using it as an argument", func() {
				cnpj, err := ParseCNPJ(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the CNPJ struct number should be empty", func() {
						So(cnpj.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a CNPJ number with an invalid second digit", func() {
			s = "11222333000182"

			Convey("And the function ParseCNPJ is called using it as an argument", func() {
				cnpj, err := ParseCNPJ(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the CNPJ struct number should be empty", func() {
						So(cnpj.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a valid CNPJ number", func() {
			s = "11222333000181"

			Convey("And the function ParseCNPJ is called using it as an argument", func() {
				cnpj, err := ParseCNPJ(s)

				Convey("It should not return an error", func() {
					So(err, ShouldEqual, nil)

					Convey("And the CNPJ struct number should exist", func() {
						So(cnpj.Number(false), ShouldEqual, "11222333000181")
						So(cnpj.Number(true), ShouldEqual, "11.222.333/0001-81")
					})
				})
			})
		})
	})
}

func TestRandomCNPJNumber(t *testing.T) {
	Convey("Given the function RandomCNPJNumber", t, func() {
		Convey("If its mask argument equals true", func() {
			number := RandomCNPJNumber(true)

			Convey("It should return a valid CNPJ number", func() {
				cnpj, err := ParseCNPJ(number)

				So(err, ShouldEqual, nil)
				So(cnpj.Number(false), ShouldNotEqual, "")
				So(cnpj.Number(true), ShouldNotEqual, "")
			})
		})

		Convey("If its mask argument equals false", func() {
			number := RandomCNPJNumber(false)

			Convey("It should return a valid CNPJ number", func() {
				cnpj, err := ParseCNPJ(number)

				So(err, ShouldEqual, nil)
				So(cnpj.Number(false), ShouldNotEqual, "")
				So(cnpj.Number(true), ShouldNotEqual, "")
			})
		})
	})
}
