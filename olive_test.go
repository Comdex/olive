package olive

import (
	"testing"

	"github.com/mitsuse/matrix-go/dense"
	"github.com/mitsuse/olive/classifier"
)

func TestInstanceFailsForNonFeatureMatirx(t *testing.T) {
	m := dense.New(2, 3)(
		0, 1, 2,
		1, 0, -1,
	)

	label := 4

	defer func() {
		if p := recover(); p == classifier.NON_FEATURE_MATRIX_PANIC {
			return
		}

		t.Fatal("NewInstance should use \"validate.ShouldBeFeature\".")
	}()
	NewInstance(m, label)
}

func TestInstanceFeature(t *testing.T) {
	feature := dense.New(1, 6)(0, 1, 2, 1, 0, -1)
	label := 4
	instance := NewInstance(feature, label)

	m := dense.New(1, 6)(0, 1, 2, 1, 0, -1)

	if instance.Feature().Equal(m) {
		return
	}

	t.Fatal("The feature vector should not be modified on creating instance.")
}

func TestInstanceLabel(t *testing.T) {
	feature := dense.New(1, 6)(0, 1, 2, 1, 0, -1)
	label := 4
	instance := NewInstance(feature, label)

	if instance.Label() == label {
		return
	}

	t.Fatal("The label should not be modified on creating instance.")
}
