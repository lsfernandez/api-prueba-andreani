package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/customer-experience/api-prueba/models"
	"github.com/customer-experience/api-prueba/validation"
)

var _ = Describe("Api", func() {
	Describe("Validando que el nombre es correcto", func() {
		Context("Nombre que es correcto", func() {
			It("Tiene que devolver true", func() {
				_, resultado := validation.ValidarNombre("Lionel Messi")
				Expect(resultado).To(Equal(true))
			})
		})
		Context("Nombre que es incorrecto, se escriben solo dos iniciales", func() {
			It("Tiene que devolver false", func() {
				_, resultado := validation.ValidarNombre("L M")
				Expect(resultado).To(Equal(false))
			})
		})
		Context("Nombre que es incorrecto, tiene caracteres especiales", func() {
			It("Tiene que devolver false", func() {
				_, resultado := validation.ValidarNombre("Lionel_ Messi@")
				Expect(resultado).To(Equal(false))
			})
		})
		Context("Nombre que es incorrecto, tiene numeros", func() {
			It("Tiene que devolver false", func() {
				_, resultado := validation.ValidarNombre("L10n3l M3ss1")
				Expect(resultado).To(Equal(false))
			})
		})
		Context("Nombre que es incorrecto, tiene solo espacios", func() {
			It("Tiene que devolver false", func() {
				_, resultado := validation.ValidarNombre("           ")
				Expect(resultado).To(Equal(false))
			})
		})
	})

	Describe("Validando que se limpian los espacios innecesarios", func() {
		Context("Nombre escrito con espacios al principio", func() {
			It("Tiene que devolver el nombre sin los espacios innecesarios", func() {
				Expect(validation.LimpiarEspacios("     Lionel Messi")).To(Equal("Lionel Messi"))
			})
		})
		Context("Nombre escrito con espacios al final", func() {
			It("Tiene que devolver el nombre sin los espacios innecesarios", func() {
				Expect(validation.LimpiarEspacios("Lionel Messi    ")).To(Equal("Lionel Messi"))
			})
		})
		Context("Nombre escrito con espacios en medio", func() {
			It("Tiene que devolver el nombre sin los espacios innecesarios", func() {
				Expect(validation.LimpiarEspacios("Lionel      Messi")).To(Equal("Lionel Messi"))
			})
		})
		Context("Nombre escrito con espacios al principio y final", func() {
			It("Tiene que devolver el nombre sin los espacios innecesarios", func() {
				Expect(validation.LimpiarEspacios("   Lionel Messi     ")).To(Equal("Lionel Messi"))
			})
		})
		Context("Nombre escrito con espacios al principio, final y en medio", func() {
			It("Tiene que devolver el nombre sin los espacios innecesarios", func() {
				Expect(validation.LimpiarEspacios("    Lionel        Messi   ")).To(Equal("Lionel Messi"))
			})
		})
		Context("Solo espacios", func() {
			It("Tiene que devolver un string vacio", func() {
				Expect(validation.LimpiarEspacios("                  ")).To(Equal(""))
			})
		})
	})

	Describe("Validando que el color es correcto", func() {
		Context("Color correcto escrito en español", func() {
			It("Tiene que devolver el color traducido en ingles", func() {
				colortraducido, _ := models.ConsultaColores("azul")
				Expect(colortraducido).To(Equal("blue"))
			})
		})
		Context("Color incorrecto porque esta escrito en ingles", func() {
			It("Tiene que devolver false", func() {
				_, resultado := models.ConsultaColores("black")
				Expect(resultado).To(Equal(false))
			})
		})
		Context("Color incorrecto porque esta escrito solo numeros", func() {
			It("Tiene que devolver false", func() {
				_, resultado := models.ConsultaColores("1234")
				Expect(resultado).To(Equal(false))
			})
		})
	})
})
