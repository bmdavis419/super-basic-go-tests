package main

import "testing"

func TestSumEmpty(t *testing.T) {
	input := []int{}
	sum := Sum(input)
	if sum != 0 {
		t.Errorf("Sum([]int{}) = %d; want 0", sum)
	}
}

func TestSumOne(t *testing.T) {
	input := []int{1}
	sum := Sum(input)
	if sum != 1 {
		t.Errorf("Sum([]int{1}) = %d; want 1", sum)
	}
}

func TestSumNormal(t *testing.T) {
	input := []int{1, 2, 3}
	sum := Sum(input)
	if sum != 6 {
		t.Errorf("Sum([]int{1, 2, 3}) = %d; want 6", sum)
	}
}
