package brazil_test

import (
	"testing"

	. "github.com/flavioltonon/go-brazil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseTituloEleitoral(t *testing.T) {
	Convey("Given a string named s", t, func() {
		var s string

		Convey("If s is empty", func() {
			s = ""

			Convey("And the function ParseTituloEleitoral is called using it as an argument", func() {
				titulo, err := ParseTituloEleitoral(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the TituloEleitoral struct number should be empty", func() {
						So(titulo.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a false-positive TituloEleitoral number", func() {
			s = "00000000000"

			Convey("And the function ParseTituloEleitoral is called using it as an argument", func() {
				titulo, err := ParseTituloEleitoral(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the TituloEleitoral struct number should be empty", func() {
						So(titulo.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a TituloEleitoral number with an invalid first digit", func() {
			s = "027442350185"

			Convey("And the function ParseTituloEleitoral is called using it as an argument", func() {
				titulo, err := ParseTituloEleitoral(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the TituloEleitoral struct number should be empty", func() {
						So(titulo.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a TituloEleitoral number with an invalid second digit", func() {
			s = "027442350174"

			Convey("And the function ParseTituloEleitoral is called using it as an argument", func() {
				titulo, err := ParseTituloEleitoral(s)

				Convey("It should return an error", func() {
					So(err, ShouldNotEqual, nil)

					Convey("And the TituloEleitoral struct number should be empty", func() {
						So(titulo.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a valid TituloEleitoral number", func() {
			s = "027442350175"

			Convey("And the function ParseTituloEleitoral is called using it as an argument", func() {
				titulo, err := ParseTituloEleitoral(s)

				Convey("It should not return an error", func() {
					So(err, ShouldEqual, nil)

					Convey("And the TituloEleitoral struct number should exist", func() {
						So(titulo.Number(false), ShouldEqual, "027442350175")
						So(titulo.Number(true), ShouldEqual, "0274 4235 0175")
					})
				})
			})
		})
	})
}

func TestRandomTituloEleitoralNumber(t *testing.T) {
	Convey("Given the function RandomTituloEleitoralNumber", t, func() {
		Convey("If its mask argument equals true", func() {
			number := RandomTituloEleitoralNumber(true)

			Convey("It should return a valid TituloEleitoral number", func() {
				titulo, err := ParseTituloEleitoral(number)

				So(titulo.Number(false), ShouldNotEqual, "")
				So(titulo.Number(true), ShouldNotEqual, "")
				So(err, ShouldEqual, nil)
			})
		})

		Convey("If its mask argument equals false", func() {
			number := RandomTituloEleitoralNumber(false)

			Convey("It should return a valid TituloEleitoral number", func() {
				titulo, err := ParseTituloEleitoral(number)

				So(titulo.Number(false), ShouldNotEqual, "")
				So(titulo.Number(true), ShouldNotEqual, "")
				So(err, ShouldEqual, nil)
			})
		})
	})
}
