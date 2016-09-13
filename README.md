# go-yara

[![GoDoc](https://godoc.org/github.com/hillu/go-yara?status.svg)](https://godoc.org/github.com/hillu/go-yara)

Go bindings for [YARA](http://plusvic.github.io/yara/), staying as
close as sensible to the library's C-API while taking inspiration from
the `yara-python` implementation.

YARA 3.4.0 or higher is required for full functionality. If you need
to build with YARA 3.3.0, please build with the `yara3.3` build tag.
(The `compat-yara-3.3` branch has been removed.)

## Installation

### Unix

On a Unix system with libyara properly installed, this should work,
provided that `GOPATH` is set:

    $ go get github.com/hillu/go-yara
    $ go install github.com/hillu/go-yara

Depending on what location libyara and its headers have been
installed, proper `CFLAGS` and `LDFLAGS` may have to be added to
`cgo.go` or be specified via environment variables (`CGO_CFLAGS` and
`CGO_LDFLAGS`).

Linker errors buried in the CGO output such as

    undefined reference to `yr_compiler_add_file'

probably mean that the linker is looking at an old YARA version.

### Windows

I have not yet built go-yara *on* Windows, only used the MinGW-w64
provided on Debian. The YARA library was built like this:

    $ ./configure --host=i686-w64-mingw32 --disable-magic --disable-cuckoo --without-crypto
    [...]
    $ make
    $ make install prefix=/path/to/i686-w64-mingw32

Compiling and linking `go-yara` against this library was achieved like
this:

    $ CC=i686-w64-mingw32-gcc \
      CGO_ENABLED=1 \
      GOOS=windows GOARCH=386 \
      CGO_CFLAGS=-I/path/to/i686-w64-mingw32/include \
      CGO_LDFLAGS=-L/path/to/i686-w64-mingw32/lib \
      go install --ldflags '-extldflags "-static"' github.com/hillu/go-yara

## License

BSD 2-clause, see LICENSE file in the source distribution.

## Author

Hilko Bengen <bengen@hilluzination.de>


## ChenQ
windows下使用动态库；需要将libyara32拷贝到运行目录下