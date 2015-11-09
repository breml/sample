[![Build Status](https://travis-ci.org/breml/sample.svg)](https://travis-ci.org/breml/sample) [![Coverage Status](https://coveralls.io/repos/breml/sample/badge.svg?branch=master&service=github)](https://coveralls.io/github/breml/sample?branch=master) [![Go Report Card](http://goreportcard.com/badge/breml/sample)](http://goreportcard.com/report/breml/sample)  
[![GoDoc](https://godoc.org/github.com/breml/sample?status.svg)](https://godoc.org/github.com/breml/sample) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

# Package sample

Go package sample implements sampling algorithms for 1 in n sampling for a random value (probe):

* Modulo, using modulo-operation
* PowerOf2, using bitwise AND-operation, only usable if the sampling rate is a power of 2
* LowerThan, checking if the probe is lower than a pre calculated boundary (maximum value for probe divided by sampling rate)
* Reciprocal, using a multiplication by the reciprocal value of the sampling rate ([Details](https://breml.github.io/blog/2015/10/22/dividable-without-remainder/))
* Decrement, implementation variant, where the random value is only calculated after a successful sampling

## Install

	go get github.com/breml/sample

## Go generate

To generate the reciprocal_uint files use:
    
    go generate *.go

## Documentation

[GoDoc](https://godoc.org/github.com/breml/sample)
