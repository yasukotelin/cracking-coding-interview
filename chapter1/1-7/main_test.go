package main

import (
	"reflect"
	"testing"
)

func TestCreateMatrix(t *testing.T) {
	expected := [][]int{
		[]int{0, 1, 2, 3, 4},
		[]int{10, 11, 12, 13, 14},
		[]int{20, 21, 22, 23, 24},
		[]int{30, 31, 32, 33, 34},
		[]int{40, 41, 42, 43, 44},
	}

	actual := createMatrix(5)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected is %v but actual is %v", expected, actual)
	}
}

func TestRotate90(t *testing.T) {

	matrix := createMatrix(5)

	expected := [][]int{
		[]int{40, 30, 20, 10, 0},
		[]int{41, 31, 21, 11, 1},
		[]int{42, 32, 22, 12, 2},
		[]int{43, 33, 23, 13, 3},
		[]int{44, 34, 24, 14, 4},
	}

	rotate90(matrix)

	if !reflect.DeepEqual(matrix, expected) {
		t.Fatalf("expected is %v but actual is %v", expected, matrix)
	}
}
