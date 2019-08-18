package main

import (
	"reflect"
	"testing"
)

func TestPaddingZero(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 3, 4},
		{10, 11, 12, 13, 14},
		{20, 21, 0, 23, 24},
		{30, 31, 32, 33, 34},
	}

	expected := [][]int{
		{0, 0, 0, 0, 0},
		{0, 11, 0, 13, 14},
		{0, 0, 0, 0, 0},
		{0, 31, 0, 33, 34},
	}

	actual := paddingZero(matrix)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected is %v but actual is %v", expected, actual)
	}
}

func TestSearchZeroPoints(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 3, 4},
		{10, 11, 12, 13, 14},
		{20, 21, 0, 23, 24},
		{30, 31, 32, 33, 34},
	}

	expected := []point{
		point{x: 0, y: 0},
		point{x: 2, y: 2},
	}

	actual := searchZeroPoints(matrix)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected is %v but actual is %v", expected, actual)
	}
}
