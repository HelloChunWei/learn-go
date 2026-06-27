package martix

import (
	"testing"
)

func TestReverseMatrix(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	actual := [][]int{
		{3, 6, 9},
		{2, 5, 8},
		{1, 4, 7},
	}

	result := reverseMartix(input)
	n := len(input)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if result[i][j] != actual[i][j] {
				t.Errorf("actual: %d, result: %d", actual[i][j], result[i][j])
				break
			}
		}
	}

	input2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	actual2 := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	result2 := reverseMartix2(input2)
	n2 := len(input2)
	for i := 0; i < n2; i++ {
		for j := 0; j < n2; j++ {
			if result2[i][j] != actual2[i][j] {
				t.Errorf("actual: %d, result: %d", actual2[i][j], result2[i][j])
				break
			}
		}
	}
}
