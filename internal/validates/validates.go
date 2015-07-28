package validates

import (
	. "github.com/mitsuse/matrix-go"
)

const (
	NON_FEATURE_MATRIX_PANIC = "NON_FEATURE_MATRIX_PANIC"
)

func ShouldBeFeature(matrix Matrix) {
	if matrix.Rows() == 1 {
		return
	}

	panic(NON_FEATURE_MATRIX_PANIC)
}
