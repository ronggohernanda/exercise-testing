package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

    "unit-test"
)

var _ = Describe("Core", func() {
	Describe("Test Multiplication", func() {
		It("Should Return Correct Number", func() {
			// pengujian menggunakan 'Expect' dari Gomega
			var result = main.Multiplication(5, 5)
			var expected int = 25

			Expect(result).To(Equal(expected))
		})
	})

	Describe("Test Division", func() {
		It("Should Return Correct Number", func() {
			// pengujian menggunakan 'Expect' dari Gomega
			var result = main.Division(100, 5)
			var expected int = 20

			Expect(result).To(Equal(expected))
		})
	})

	Describe("Test Modulus", func() {
		It("Should Return Correct Number", func() {
			// pengujian menggunakan 'Expect' dari Gomega
			var result = main.Modulus(10, 3)
			var expected int = 1

			Expect(result).To(Equal(expected))
		})
	})
})
