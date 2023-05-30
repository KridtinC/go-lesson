package calculator_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"unit-test/src/calculator"
)

var _ = Describe("Pow", func() {
	It("should calculate pow successful", func() {
		got, err := calculator.PositivePow(2, 2)
		want := 4.0
		Expect(err).To(BeNil())
		Expect(got).To(Equal(want))
	})
})
