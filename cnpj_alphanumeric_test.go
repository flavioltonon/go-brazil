package brazil_test

import (
	"testing"

	. "github.com/flavioltonon/go-brazil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseAlphanumericCNPJ(t *testing.T) {
	Convey("Given a string named s", t, func() {
		var s string

		Convey("If s is empty", func() {
			s = ""

			Convey("And the function ParseAlphanumericCNPJ is called using it as an argument", func() {
				cnpj, err := ParseAlphanumericCNPJ(s)

				Convey("It should return an error", func() {
					So(err, ShouldEqual, ErrIncorrectLenghtCNPJNumber)

					Convey("And the CNPJ struct number should be empty", func() {
						So(cnpj.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a valid alphanumeric CNPJ number", func() {
			s = "12ABC345000188"

			Convey("And the function ParseAlphanumericCNPJ is called using it as an argument", func() {
				cnpj, err := ParseAlphanumericCNPJ(s)

				Convey("It should not return an error", func() {
					So(err, ShouldEqual, nil)

					Convey("And the CNPJ struct number should exist", func() {
						So(cnpj.Number(false), ShouldEqual, "12ABC345000188")
						So(cnpj.Number(true), ShouldEqual, "12.ABC.345/0001-88")
					})
				})
			})
		})

		Convey("If s is a valid masked alphanumeric CNPJ number", func() {
			s = "12.ABC.345/0001-88"

			Convey("And the function ParseAlphanumericCNPJ is called using it as an argument", func() {
				cnpj, err := ParseAlphanumericCNPJ(s)

				Convey("It should not return an error", func() {
					So(err, ShouldEqual, nil)

					Convey("And the CNPJ struct number should exist", func() {
						So(cnpj.Number(false), ShouldEqual, "12ABC345000188")
						So(cnpj.Number(true), ShouldEqual, "12.ABC.345/0001-88")
					})
				})
			})
		})

		Convey("If s is a valid alphanumeric CNPJ number with lowercase letters", func() {
			s = "12abc345000188"

			Convey("And the function ParseAlphanumericCNPJ is called using it as an argument", func() {
				cnpj, err := ParseAlphanumericCNPJ(s)

				Convey("It should not return an error", func() {
					So(err, ShouldEqual, nil)

					Convey("And the CNPJ struct number should be normalized to uppercase", func() {
						So(cnpj.Number(false), ShouldEqual, "12ABC345000188")
					})
				})
			})
		})

		Convey("If s is a valid legacy numeric CNPJ number", func() {
			s = "12345678000195"

			Convey("And the function ParseAlphanumericCNPJ is called using it as an argument", func() {
				cnpj, err := ParseAlphanumericCNPJ(s)

				Convey("It should not return an error", func() {
					So(err, ShouldEqual, nil)

					Convey("And the CNPJ struct number should exist", func() {
						So(cnpj.Number(false), ShouldEqual, "12345678000195")
						So(cnpj.Number(true), ShouldEqual, "12.345.678/0001-95")
					})
				})
			})
		})

		Convey("If s is an alphanumeric CNPJ number with invalid check digits", func() {
			s = "12ABC345000100"

			Convey("And the function ParseAlphanumericCNPJ is called using it as an argument", func() {
				cnpj, err := ParseAlphanumericCNPJ(s)

				Convey("It should return an error", func() {
					So(err, ShouldEqual, ErrInvalidCNPJFirstDigit)

					Convey("And the CNPJ struct number should be empty", func() {
						So(cnpj.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a CNPJ number with all characters equal", func() {
			s = "00000000000000"

			Convey("And the function ParseAlphanumericCNPJ is called using it as an argument", func() {
				cnpj, err := ParseAlphanumericCNPJ(s)

				Convey("It should return an error", func() {
					So(err, ShouldEqual, ErrRepeatedCNPJNumber)

					Convey("And the CNPJ struct number should be empty", func() {
						So(cnpj.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a CNPJ number with an illegal character", func() {
			s = "12@BC345000188"

			Convey("And the function ParseAlphanumericCNPJ is called using it as an argument", func() {
				cnpj, err := ParseAlphanumericCNPJ(s)

				Convey("It should return an error", func() {
					So(err, ShouldEqual, ErrInvalidCNPJCharacter)

					Convey("And the CNPJ struct number should be empty", func() {
						So(cnpj.Number(false), ShouldEqual, "")
					})
				})
			})
		})

		Convey("If s is a CNPJ number with a letter in the check digit positions", func() {
			s = "12ABC3450001A8"

			Convey("And the function ParseAlphanumericCNPJ is called using it as an argument", func() {
				cnpj, err := ParseAlphanumericCNPJ(s)

				Convey("It should return an error", func() {
					So(err, ShouldEqual, ErrInvalidCNPJCharacter)

					Convey("And the CNPJ struct number should be empty", func() {
						So(cnpj.Number(false), ShouldEqual, "")
					})
				})
			})
		})
	})
}

func TestRandomAlphanumericCNPJNumber(t *testing.T) {
	Convey("Given the function RandomAlphanumericCNPJNumber", t, func() {
		Convey("If its mask argument equals true", func() {
			number := RandomAlphanumericCNPJNumber(true)

			Convey("It should return a valid alphanumeric CNPJ number", func() {
				cnpj, err := ParseAlphanumericCNPJ(number)

				So(err, ShouldEqual, nil)
				So(cnpj.Number(false), ShouldNotEqual, "")
				So(cnpj.Number(true), ShouldNotEqual, "")
			})
		})

		Convey("If its mask argument equals false", func() {
			number := RandomAlphanumericCNPJNumber(false)

			Convey("It should return a valid alphanumeric CNPJ number", func() {
				cnpj, err := ParseAlphanumericCNPJ(number)

				So(err, ShouldEqual, nil)
				So(cnpj.Number(false), ShouldNotEqual, "")
				So(cnpj.Number(true), ShouldNotEqual, "")
			})
		})
	})
}
