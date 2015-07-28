package classifier

import (
	"testing"
)

type constructionTest struct {
	classSize  int
	dimensions int
}

func TestNewSucceeds(t *testing.T) {
	test := constructionTest{
		classSize:  4,
		dimensions: 8,
	}

	defer func() {
		if p := recover(); p == nil {
			return
		}

		t.Fatal("New should not panic for one or more classes and positive dimmensions.")
	}()
	New(test.classSize, test.dimensions)
}

func TestNewPanicsForNonPositiveClassSize(t *testing.T) {
	test := constructionTest{
		classSize:  0,
		dimensions: 8,
	}

	defer func() {
		if p := recover(); p == nil {
			t.Fatal("The size of claasses should be one or more.")
		}
	}()
	New(test.classSize, test.dimensions)
}
