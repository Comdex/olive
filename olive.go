/*
Package "olive" provides online algorithms for machine learning.
*/
package olive

import (
	. "github.com/mitsuse/matrix-go"
	"github.com/mitsuse/olive/classifier"
	"github.com/mitsuse/olive/internal/validates"
)

type Learner interface {
	Learn(c *classifier.Classifier, instances []*Instance) *classifier.Classifier
}

type Instance struct {
	feature Matrix
	label   int
}

func NewInstance(feature Matrix, label int) *Instance {
	validates.ShouldBeFeature(feature)

	i := &Instance{
		feature: feature,
		label:   label,
	}

	return i
}

func (i *Instance) Feature() Matrix {
	return i.feature
}

func (i *Instance) Label() int {
	return i.label
}
