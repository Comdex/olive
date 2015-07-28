package validates

import (
	"testing"

	. "github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/dense"
)

type weightUpdateTest struct {
	old    Matrix
	update Matrix
}

func TestShouldBeOneOrMoreClassesSucceedsForPositiveSizedClasses(t *testing.T) {
	test := 4

	defer func() {
		if p := recover(); p != nil {
			t.Fatal("classes-size validation should not cause any panic for positve size.")
		}
	}()
	ShouldBeOneOrMoreClasses(test)
}

func TestShouldBeOneOrMoreClassesPanicsForNonPositiveSizedClasses(t *testing.T) {
	test := 0

	defer func() {
		if p := recover(); p == INVALID_CLASS_SIZE {
			return
		}

		t.Fatal("The size of classes should be one or more.")
	}()
	ShouldBeOneOrMoreClasses(test)
}

func TestShouldBeFeatureSucceedsForFeature(t *testing.T) {
	test := dense.Zeros(1, 1)

	defer func() {
		if p := recover(); p != nil {
			t.Fatal("Feature validation should not cause any panic for feature matirx.")
		}
	}()
	ShouldBeFeature(test)
}

func TestShouldBeFeaturePanicsForNonFeature(t *testing.T) {
	test := dense.Zeros(4, 1)

	defer func() {
		if p := recover(); p == NON_FEATURE_MATRIX_PANIC {
			return
		}

		t.Fatal("Feature should be a matrix which the number of rows equals to 1.")
	}()
	ShouldBeFeature(test)
}

func TestShouldBeCompatibleWeightSucceedsForSameShapeMatrices(t *testing.T) {
	test := weightUpdateTest{
		old:    dense.Zeros(4, 8),
		update: dense.Zeros(4, 8),
	}

	defer func() {
		if p := recover(); p == nil {
			return
		}

		t.Fatal(
			"The validation should not cause any panic for matrices with same shape.",
		)
	}()
	ShouldBeCompatibleWeights(test.old, test.update)
}

func TestShouldBeCompatibleWeightPanicsForMatricsWithDifferentClassSize(t *testing.T) {
	test := weightUpdateTest{
		old:    dense.Zeros(4, 8),
		update: dense.Zeros(2, 8),
	}

	defer func() {
		if p := recover(); p == INCOMPATIBLE_WEIGHTS_PANIC {
			return
		}

		t.Fatal("The new weights should have same size of classes as the old one.")
	}()
	ShouldBeCompatibleWeights(test.old, test.update)
}

func TestShouldBeCompatibleWeightPanicsForMatricsWithDifferentDimmensions(t *testing.T) {
	test := weightUpdateTest{
		old:    dense.Zeros(4, 8),
		update: dense.Zeros(4, 10),
	}

	defer func() {
		if p := recover(); p == INCOMPATIBLE_WEIGHTS_PANIC {
			return
		}

		t.Fatal("The new weights should have same dimmensions as the old one.")
	}()
	ShouldBeCompatibleWeights(test.old, test.update)
}
