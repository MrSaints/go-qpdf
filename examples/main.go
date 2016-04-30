package main

import (
	"fmt"
	"github.com/mrsaints/go-qpdf/qpdf"
)

func main() {
	fmt.Printf("QPDF Version: %s\n", qpdf.Version())

	q := qpdf.Init()
	defer q.GC()

	err := q.Open("test.pdf")
	if err != nil {
		panic(err)
	}

	err = q.SetOutput("test-linearized.pdf")
	if err != nil {
		panic(err)
	}

	q.Linearize()

	err = q.Write()
	if err != nil {
		panic(err)
	}

	if q.HasError() {
		panic(q.LastError())
	}

	fmt.Println("Linearized PDF!")
}
