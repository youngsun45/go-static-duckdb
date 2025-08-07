// file: duckdb.go
package duckdb

/*
#cgo windows,amd64 CXXFLAGS: -I./include -std=c++11
#cgo windows,amd64 LDFLAGS: -L./lib -lduckdb -lstdc++

// We must include the C++ header here.
#include "duckdb.hpp"

// ---- C WRAPPER FUNCTION ----
// Go cannot call C++ `DuckDB::LibraryVersion()` directly.
// So, we create a C-compatible function that wraps the C++ call.
// The `extern "C"` block tells the C++ compiler to create a function
// with a simple C name, which cgo can easily find.
extern "C" const char* duckdb_get_library_version() {
    return DuckDB::LibraryVersion();
}
// ----------------------------
*/
import "C"

// LibraryVersion returns the version of the linked DuckDB library.
// It works by calling our C wrapper function.
func LibraryVersion() string {
	// Call our C wrapper function, which is exposed to Go via the C pseudo-package.
	c_version := C.duckdb_get_library_version()

	// Convert the C string (const char*) to a Go string.
	return C.GoString(c_version)
}
