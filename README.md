# Daycount

[![GoDoc](https://godoc.org/github.com/edgelaboratories/daycount?status.png)](http://godoc.org/github.com/edgelaboratories/daycount)
[![Build Status](https://api.travis-ci.org/edgelaboratories/daycount.svg?branch=master)](https://travis-ci.org/edgelaboratories/daycount)
[![Go Report Card](https://goreportcard.com/badge/github.com/edgelaboratories/daycount)](https://goreportcard.com/report/github.com/edgelaboratories/daycount)

## Description

Package `daycount` provides daycounting methods.

## Installation

    go get -u github.com/edgelaboratories/daycount

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/edgelaboratories/daycount"
)

func main() {
	from := date.New(2018, time.January, 1)
	to := date.New(2018, time.July, 31)
	yf := daycount.YearFractionDiff(from, to, daycount.ActualActual)
	fmt.Printf("year fraction between the two dates is %0.2f\n", yf)
}
```
