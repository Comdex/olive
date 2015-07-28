package validates

import (
	"testing"

	"github.com/mitsuse/matrix-go/dense"
)

func TestMatrixShouldBeFeatureSucceedsForFeature(t *testing.T) {
	test := dense.Zeros(1, 1)

	defer func() {
		if p := recover(); p != nil {
			t.Fatal("Feature validation should not cause any panic for feature matirx.")
		}
	}()
	ShouldBeFeature(test)
}

func TestMatrixShouldBeFeaturePanicsForNonFeature(t *testing.T) {
	test := dense.Zeros(4, 1)

	defer func() {
		if p := recover(); p == NON_FEATURE_MATRIX_PANIC {
			return
		}

		t.Fatal("Feature should be a matrix which the number of rows equals to 1.")
	}()
	ShouldBeFeature(test)
}
