# Daycount

[![GoDoc](https://godoc.org/github.com/edgelaboratories/daycount?status.png)](http://godoc.org/github.com/edgelaboratories/daycount)
![Build Status](https://github.com/edgelaboratories/daycount/workflows/Test/badge.svg)
![GolangCI Lint](https://github.com/edgelaboratories/daycount/workflows/golangci/badge.svg)

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
