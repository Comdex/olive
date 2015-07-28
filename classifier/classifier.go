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

const (
	INVALID_CLASS_SIZE         = validates.INVALID_CLASS_SIZE
	NON_FEATURE_MATRIX_PANIC   = validates.NON_FEATURE_MATRIX_PANIC
	INCOMPATIBLE_WEIGHTS_PANIC = validates.INCOMPATIBLE_WEIGHTS_PANIC
)

// Classifier is an implementation of multi-class linear classifier.
type Classifier struct {
	weights Matrix
}

// Create a new classifier with the given number of classes and dimensions of features.
// The generate classifier has a weight matrix with zeros.
func New(classSize, dimensions int) *Classifier {
	validates.ShouldBeOneOrMoreClasses(classSize)

	c := &Classifier{
		// TODO: Replace the weights with the immutable version of dense matrix.
		weights: dense.Zeros(classSize, dimensions),
	}

	return c
}

// Return the size of classes.
func (c *Classifier) ClassSize() int {
	return c.weights.Rows()
}

// Return the dimensions of features.
func (c *Classifier) Dimensions() int {
	return c.weights.Columns()
}

// Update the weight matrix.
// The new weight matrix should have same shape as the old one.
func (c *Classifier) Update(weights Matrix) *Classifier {
	validates.ShouldBeCompatibleWeights(c.weights, weights)

	c.weights = weights

	return c
}

// Classify the given feature matrix into one of the classes.
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
