package main

/*
#include <qpdf/qpdf-c.h>
#cgo pkg-config: libqpdf
*/
import "C"
import "fmt"

type QPDF struct {
	qpdfdata C.qpdf_data
}

func Version() string {
	return C.GoString(C.qpdf_get_qpdf_version())
}

func Init() *QPDF {
	qpdf := new(QPDF)
	qpdf.qpdfdata = C.qpdf_init()
	return qpdf
}

func (q *QPDF) Open(fn string) error {
	err := C.qpdf_read(q.qpdfdata, C.CString(fn), nil)
	if err != 0 {
		return fmt.Errorf("error reading input PDF file, code: %+v", err)
	}
	return nil
}

func (q *QPDF) SetOutput(fn string) error {
	err := C.qpdf_init_write(q.qpdfdata, C.CString(fn))
	if err != C.QPDF_SUCCESS {
		return fmt.Errorf("error creating output PDF file, code: %+v", err)
	}
	return nil
}

func (q *QPDF) Write() error {
	err := C.qpdf_write(q.qpdfdata)
	if err != C.QPDF_SUCCESS {
		return fmt.Errorf("error creating output PDF file, code: %+v", err)
	}
	return nil
}

func (q *QPDF) Linearize() {
	C.qpdf_set_linearization(q.qpdfdata, C.QPDF_TRUE)
}

func main() {
	fmt.Printf("QPDF Version: %s\n", Version())

	qpdf := Init()
	err := qpdf.Open("test.pdf")
	if err != nil {
		panic(err)
	}

	err = qpdf.SetOutput("test-linearized.pdf")
	if err != nil {
		panic(err)
	}

	qpdf.Linearize()

	err = qpdf.Write()
	if err != nil {
		panic(err)
	}

	fmt.Println("Linearized PDF!")
}
