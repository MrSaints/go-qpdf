// Package qpdf provides simple Go bindings for qpdf C / C++ API.
// Due to the limitations of the C API, it is only capable of doing
// basic transformations on PDF files; mostly linearization.
// Refer to examples/ for usage information.
// For more information: https://github.com/qpdf/qpdf/
package qpdf

/*
#include <stdlib.h>
#include <qpdf/qpdf-c.h>
#cgo pkg-config: libqpdf
*/
import "C"

import (
	"errors"
	"unsafe"
)

// QPDF contains the qpdf_data object.
// It should never initialised manually (i.e. new(QPDF) or QPDF{}).
// It should be initialised using the Init method.
type QPDF struct {
	// qpdfdata contains the qpdf_data object.
	// The state of all operations / transformations on a source PDF
	// is stored in this qpdf_data object.
	// An instance of qpdf_data is not thread-safe. It must not be accessed
	// simultaneously by multiple threads.
	qpdfdata C.qpdf_data
}

// Version is an implementation of qpdf_get_qpdf_version.
// It returns the current version of the qpdf software / library as a string.
func Version() string {
	return C.GoString(C.qpdf_get_qpdf_version())
}

// Init is an implementation of qpdf_init.
// It returns a QPDF object containing an initialised qpdf_data pointer.
// This method must be called before any other QPDF operations.
func Init() *QPDF {
	qpdf := new(QPDF)
	qpdf.qpdfdata = C.qpdf_init()
	return qpdf
}

// GC is an implementation of qpdf_cleanup.
// All dynamic memory - except the qpdf_data object - is managed by the
// qpdf library.
// This method may be used to free up the dynamically allocated qpdf_data
// pointer. After calling this method, any pointers to the qpdf_data object
// will no longer be valid.
func (q *QPDF) GC() {
	C.qpdf_cleanup(&q.qpdfdata)
}

// HasError is an implementation of qpdf_has_error.
// It checks if there are any outstanding / unhandled errors in qpdf_data,
// and if so, it will return true.
func (q *QPDF) HasError() bool {
	if C.qpdf_has_error(q.qpdfdata) == C.QPDF_TRUE {
		return true
	}
	return false
}

// LastError is an implementation of qpdf_get_error, and
// qpdf_get_error_full_text.
// It attempts to get the last error condition, and if there is an error,
// it will pass the condition to qpdf_get_error_full_text to retrieve
// the full / user-friendly error message from the qpdf library.
// Upon calling this method, HasError will return false until the next
// error condition occurs (i.e. it drains the errors).
func (q *QPDF) LastError() error {
	err := C.qpdf_get_error(q.qpdfdata)
	full := C.qpdf_get_error_full_text(q.qpdfdata, err)
	return errors.New(C.GoString(full))
}

// Open is an implementation of qpdf_read.
// It reads, and parses the source PDF file into qpdf_data.
// Methods to manipulate the data can be called after the file is opened
// with this method.
func (q *QPDF) Open(fn string) error {
	c_fn := C.CString(fn)
	defer C.free(unsafe.Pointer(c_fn))
	err := C.qpdf_read(q.qpdfdata, c_fn, nil)
	if err != C.QPDF_SUCCESS {
		return q.LastError()
	}
	return nil
}

// SetOutput is an implementation of qpdf_init_write.
// It initialises qpdf_data for write operations, and it prepares the target
// file for PDF data to be written.
func (q *QPDF) SetOutput(fn string) error {
	c_fn := C.CString(fn)
	defer C.free(unsafe.Pointer(c_fn))
	err := C.qpdf_init_write(q.qpdfdata, c_fn)
	if err != C.QPDF_SUCCESS {
		return q.LastError()
	}
	return nil
}

// Write is an implementation of qpdf_write.
// It carries out the actual write operation using data from the qpdf_data
// object. The output is written to the target file defined in SetOutput,
// assuming no errors occured.
func (q *QPDF) Write() error {
	err := C.qpdf_write(q.qpdfdata)
	if err != C.QPDF_SUCCESS {
		return q.LastError()
	}
	return nil
}

// Linearize is an implementation of qpdf_set_linearization.
// It does not actually linearize the qpdf_data object immediately.
// Instead, enables linearization mode. The actual transformation occurs
// during qpdf_write.
func (q *QPDF) Linearize() {
	C.qpdf_set_linearization(q.qpdfdata, C.QPDF_TRUE)
}
