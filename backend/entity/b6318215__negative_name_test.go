package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegative(t *testing.T) {
	g := NewGomegaWithT(t)

	customer := Customer{
		Name:       "",
		Email:      "nut@gmail.com",
		CustomerID: "M0123456",
	}

	ok, err := govalidator.ValidateStruct(customer)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("ไม่ได้กรอกข้อมูลชื่อ"))
}