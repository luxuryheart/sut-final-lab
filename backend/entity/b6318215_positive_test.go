package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	customer := Customer{
		Name:       "nut",
		Email:      "B6318215@gmail.com",
		CustomerID: "M0123456",
	}

	ok, err := govalidator.ValidateStruct(customer)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
}
