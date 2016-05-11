package util

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSrc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "util")
}

var _ = Describe("util", func() {

	BeforeEach(func() {
		//write your initial code here
	})

	Describe("Que envio array de string para que se concatenen.", func() {

		Context("envio un array de string.", func() {

			It("El servicio debe devolver un string con los datos concatenados.", func() {
				response := Concatenate([]string{"esto", "=", "esto"})
				Expect("esto=esto").To(Equal(response)) //Validación
			})

		})

		/*Context("dígito los datos token invalido de sesión, contraseña nueva y confirmación igual  a la nueva contraseña.", func() {

			It("El servicio deberá mostrar un mensaje de error.", func() {
				Expect(1).To(Equal(1)) //Validación
			})

		})*/

	})

})
