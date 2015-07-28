# Olive

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)][godoc]
[![Version](https://img.shields.io/github/tag/mitsuse/olive.svg?style=flat-square)][release]
[![Wercker](http://img.shields.io/wercker/ci/55b573c7f32b86a9292fcdec.svg?style=flat-square)][wercker]
[![Coverage](https://img.shields.io/codecov/c/github/mitsuse/olive/develop.svg?style=flat-square)][coverage]

[license]: LICENSE.txt
[godoc]: http://godoc.org/github.com/mitsuse/olive
[release]: https://github.com/mitsuse/olive/releases
[wercker]: https://app.wercker.com/project/bykey/3edc447f6194909364e46e1de66254a3
[coverage]: https://codecov.io/github/mitsuse/olive

Online algorithms for machine learning implemented in [Golang][golang].

[golang]: http://golang.org/


## Installation

For installation, execute the following command:

```
$ go get github.com/mitsuse/olive
```

### Requirements

Olive uses the following libraries:

- [`mitsuse/matrix-go (v0.1.0)`][repo-matrix-go]

[repo-matrix-go]: https://github.com/mitsuse/matrix-go


## Features

### Algorithms

Not provided yet.


### Classifier

Olive provides an implementation of multi-class linear classifier.


#### Creation

Call [`classifier.New`][doc-classifier.New] with the size of classes and the dimensions of features.

[doc-classifier.New]: http://godoc.org/github.com/mitsuse/olive/classifier/#New

```go
classSize, dimensions := 4, 8

c := classifier.New(classSize, dimensions)
```


#### Classification

Apply [`(*Classifier).Classify`][doc-Classifier.Classify] to a feature matrix by classifying it.
The feature matrix is typed as [`Matrix`][doc-Matrix] of mitsuse/matrix-go.

[doc-Classifier.Classify]: http://godoc.org/github.com/mitsuse/olive/classifier/#Classifier.Classify
[doc-Matrix]: http://godoc.org/github.com/mitsuse/matrix-go/#Matrix

```go
features := dense.New(classSize, dimensions)(
    0, 0, 1, 1, 0.5, 0.1, -2, 3, 0,
)

class := c.Classify(features)
```


## License

Please read [LICENSE.txt](LICENSE.txt).
