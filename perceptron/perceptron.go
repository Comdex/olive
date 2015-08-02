package perceptron

import (
	"github.com/mitsuse/olive"
	"github.com/mitsuse/olive/classifier"
)

type Perceptron struct {
	iterations int
}

func New(iterations int) olive.Learner {
	p := &Perceptron{
		iterations: iterations,
	}

	return p
}

func (p *Perceptron) Learn(
	c *classifier.Classifier, instances []*olive.Instance,
) *classifier.Classifier {
	for i := 0; i < p.iterations; i++ {
		for _, instance := range instances {
			class := c.Classify(instance.Feature())

			if class != instance.Class() {
				// TODO: Update the weights.
			}
		}
	}

	return c
}
