# go-qpdf

[![GoDoc](https://godoc.org/github.com/mrsaints/go-qpdf/qpdf?status.svg)](https://godoc.org/github.com/mrsaints/go-qpdf/qpdf)

Simple Go bindings for [qpdf](https://github.com/qpdf/qpdf) C / C++ API; mostly for linearization.

I would not recommend using this on production. I only worked on it to experiment with [cgo](https://golang.org/cmd/cgo/), and I do not plan on maintaining it very often. Contributions are nevertheless, welcomed.


## Dependencies

To build, and run the package, you must have `libqpdf` installed.

On Debian systems, this can be achieved using
`apt-get install libqpdf-dev`.


## Usage

1. Download, and install `go-qpdf/qpdf`:

    ```shell
    go get github.com/MrSaints/go-qpdf/qpdf
    ```

2. Import the package into your code:

    ```go
    import "github.com/MrSaints/go-qpdf/qpdf"
    ```

View the [GoDoc][], [examples][] or [code][] for more information.


[GoDoc]: https://godoc.org/github.com/mrsaints/go-qpdf/qpdf
[examples]: examples/
[code]: qpdf/