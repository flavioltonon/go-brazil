package brazil_test

import (
	"testing"

	. "github.com/flavioltonon/go-brazil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseCPF(t *testing.T) {
	Convey("Given a string named s", t, func() {
		var s string

		Convey("If s is empty", func() {
			s = ""

			Convey("And the function ParseCPF is called using it as an argument", func() {
				cpf, err := ParseCPF(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the CPF struct number should be empty", func() {
						So(cpf.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a false-positive CPF number", func() {
			s = "11111111111"

			Convey("And the function ParseCPF is called using it as an argument", func() {
				cpf, err := ParseCPF(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the CPF struct number should be empty", func() {
						So(cpf.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a CPF number with an invalid first digit", func() {
			s = "05143026065"

			Convey("And the function ParseCPF is called using it as an argument", func() {
				cpf, err := ParseCPF(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the CPF struct number should be empty", func() {
						So(cpf.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a CPF number with an invalid second digit", func() {
			s = "05143026074"

			Convey("And the function ParseCPF is called using it as an argument", func() {
				cpf, err := ParseCPF(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the CPF struct number should be empty", func() {
						So(cpf.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a valid CPF number", func() {
			s = "05143026075"

			Convey("And the function ParseCPF is called using it as an argument", func() {
				cpf, err := ParseCPF(s)

				Convey("It should not return an error", func() {
					So(err, ShouldEqual, nil)

					Convey("And the CPF struct number should exist", func() {
						So(cpf.Number(false), ShouldEqual, "05143026075")
						So(cpf.Number(true), ShouldEqual, "051.430.260-75")
					})
				})
			})
		})
	})
}

func TestRandomCPFNumber(t *testing.T) {
	Convey("Given the function RandomCPFNumber", t, func() {
		Convey("If its mask argument equals true", func() {
			number := RandomCPFNumber(true)

			Convey("It should return a valid CPF number", func() {
				cpf, err := ParseCPF(number)

				So(cpf.Number(false), ShouldNotEqual, "")
				So(cpf.Number(true), ShouldNotEqual, "")
				So(err, ShouldEqual, nil)
			})
		})

		Convey("If its mask argument equals false", func() {
			number := RandomCPFNumber(false)

			Convey("It should return a valid CPF number", func() {
				cpf, err := ParseCPF(number)

				So(cpf.Number(false), ShouldNotEqual, "")
				So(cpf.Number(true), ShouldNotEqual, "")
				So(err, ShouldEqual, nil)
			})
		})
	})
}
