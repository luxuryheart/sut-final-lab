package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositiveCustomer(t *testing.T) {
	g := NewGomegaWithT(t)

	

	customer := Customer{
		Name:       "nut",
		Email:      "nut@gmail.com",
		CustomerID: "B6318215",
	}

	ok, err := govalidator.ValidateStruct(customer)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("ไม่ตรงตามรูปแบบ"))
}