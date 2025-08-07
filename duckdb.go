// file: duckdb.go
package duckdb

// These CGO directives are the heart of this package.
// They are only active on the "windows" OS and "amd64" architecture.
// For other systems, they are ignored.

/*
#cgo windows,amd64 CXXFLAGS: -I./include -std=c++11
#cgo windows,amd64 LDFLAGS: -L./lib -lduckdb -lstdc++
#include "duckdb.hpp"
*/
import "C"

import "unsafe"

// LibraryVersion returns the version of the linked DuckDB library.
// This is a great way to test if the linking was successful.
func LibraryVersion() string {
    // C.DuckDB::LibraryVersion() returns a C++ string (const char*).
    c_version := C.DuckDB::LibraryVersion()
    // We need to convert it to a Go string.
    return C.GoString(c_version)
}

// You can add more Go functions here to wrap other DuckDB C++ APIs.