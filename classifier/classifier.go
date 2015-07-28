/*
Package "classifier" provides an implementation of multi-class linear classifier.
*/
package classifier

import (
	"math"

	. "github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/dense"
	"github.com/mitsuse/olive/internal/validates"
)

// Classifier is an implementation of multi-class linear classifier.
type Classifier struct {
	weights Matrix
}

func New(classSize, dimensions int) *Classifier {
	c := &Classifier{
		// TODO: Replace the weights with the immutable version of dense matrix.
		weights: dense.Zeros(classSize, dimensions),
	}

	return c
}

func (c *Classifier) Classify(feature Matrix) (label int) {
	validates.ShouldBeFeature(feature)

	scores := c.weights.Multiply(feature.Transpose())

	score := math.Inf(-1)

	// TODO: Use "(Matrix).Max()" instead of finding the max element with "(Matrix).All()".
	for cursor := scores.All(); cursor.HasNext(); {
		if s, l, _ := cursor.Get(); score < s || label < 0 {
			score = s
			label = l
		}
	}

	return label
}
