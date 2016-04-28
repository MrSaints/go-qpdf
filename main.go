package main

/*
#include <qpdf/qpdf-c.h>
#cgo pkg-config: libqpdf
*/
import "C"
import "fmt"
import "errors"

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

func (q *QPDF) GC() {
	C.qpdf_cleanup(&q.qpdfdata)
}

func (q *QPDF) HasError() bool {
	if C.qpdf_has_error(q.qpdfdata) == C.QPDF_TRUE {
		return true
	}
	return false
}

func (q *QPDF) LastError() error {
	err := C.qpdf_get_error(q.qpdfdata)
	full := C.qpdf_get_error_full_text(q.qpdfdata, err)
	return errors.New(C.GoString(full))
}

func (q *QPDF) Open(fn string) error {
	err := C.qpdf_read(q.qpdfdata, C.CString(fn), nil)
	if err != C.QPDF_SUCCESS {
		return q.LastError()
	}
	return nil
}

func (q *QPDF) SetOutput(fn string) error {
	err := C.qpdf_init_write(q.qpdfdata, C.CString(fn))
	if err != C.QPDF_SUCCESS {
		return q.LastError()
	}
	return nil
}

func (q *QPDF) Write() error {
	err := C.qpdf_write(q.qpdfdata)
	if err != C.QPDF_SUCCESS {
		return q.LastError()
	}
	return nil
}

func (q *QPDF) Linearize() {
	C.qpdf_set_linearization(q.qpdfdata, C.QPDF_TRUE)
}

func main() {
	fmt.Printf("QPDF Version: %s\n", Version())

	qpdf := Init()
	defer qpdf.GC()

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

	if qpdf.HasError() {
		panic(qpdf.LastError())
	}

	fmt.Println("Linearized PDF!")
}
