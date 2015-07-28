package validates

import (
	"testing"

	"github.com/mitsuse/matrix-go/dense"
)

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
