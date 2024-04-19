package main

import "testing"

func TestMultiplication(t *testing.T) {
	var result = Multiplication(5, 5)
	var expected int = 25

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestDivision(t *testing.T) {
	var result = Division(100, 5)
	var expected int = 20

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestModulus(t *testing.T) {
	var result = Modulus(10, 3)
	var expected int = 1

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
