/*
Copyright 2022

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//
// !! Generated code -- do not modify !!
//

package benchmarks_test

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
	"unsafe"
)

func TestPrintSizes(t *testing.T) {
	w := tabwriter.NewWriter(os.Stdout, 4, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "| T\t real(T)\t size(T)\t size(*T)\t size(any(T))\t size(any(*T)) |\n")
	fmt.Fprintf(w, "|:---:|:---:|:---:|:---:|:---:|:---:|\n")
	const s = "| %s\t `%T`\t %d\t %d\t %d\t %d |\n"
	fmt.Fprintf(w, s, "int", _int, unsafe.Sizeof(_int), unsafe.Sizeof(&_int), unsafe.Sizeof(interface{}(_int)), unsafe.Sizeof(interface{}(&_int)))
	fmt.Fprintf(w, s, "int8", _int8, unsafe.Sizeof(_int8), unsafe.Sizeof(&_int8), unsafe.Sizeof(interface{}(_int8)), unsafe.Sizeof(interface{}(&_int8)))
	fmt.Fprintf(w, s, "int16", _int16, unsafe.Sizeof(_int16), unsafe.Sizeof(&_int16), unsafe.Sizeof(interface{}(_int16)), unsafe.Sizeof(interface{}(&_int16)))
	fmt.Fprintf(w, s, "int32", _int32, unsafe.Sizeof(_int32), unsafe.Sizeof(&_int32), unsafe.Sizeof(interface{}(_int32)), unsafe.Sizeof(interface{}(&_int32)))
	fmt.Fprintf(w, s, "int64", _int64, unsafe.Sizeof(_int64), unsafe.Sizeof(&_int64), unsafe.Sizeof(interface{}(_int64)), unsafe.Sizeof(interface{}(&_int64)))
	fmt.Fprintf(w, s, "uint", _uint, unsafe.Sizeof(_uint), unsafe.Sizeof(&_uint), unsafe.Sizeof(interface{}(_uint)), unsafe.Sizeof(interface{}(&_uint)))
	fmt.Fprintf(w, s, "uint8", _uint8, unsafe.Sizeof(_uint8), unsafe.Sizeof(&_uint8), unsafe.Sizeof(interface{}(_uint8)), unsafe.Sizeof(interface{}(&_uint8)))
	fmt.Fprintf(w, s, "uint16", _uint16, unsafe.Sizeof(_uint16), unsafe.Sizeof(&_uint16), unsafe.Sizeof(interface{}(_uint16)), unsafe.Sizeof(interface{}(&_uint16)))
	fmt.Fprintf(w, s, "uint32", _uint32, unsafe.Sizeof(_uint32), unsafe.Sizeof(&_uint32), unsafe.Sizeof(interface{}(_uint32)), unsafe.Sizeof(interface{}(&_uint32)))
	fmt.Fprintf(w, s, "uint64", _uint64, unsafe.Sizeof(_uint64), unsafe.Sizeof(&_uint64), unsafe.Sizeof(interface{}(_uint64)), unsafe.Sizeof(interface{}(&_uint64)))
	fmt.Fprintf(w, s, "float32", _float32, unsafe.Sizeof(_float32), unsafe.Sizeof(&_float32), unsafe.Sizeof(interface{}(_float32)), unsafe.Sizeof(interface{}(&_float32)))
	fmt.Fprintf(w, s, "float64", _float64, unsafe.Sizeof(_float64), unsafe.Sizeof(&_float64), unsafe.Sizeof(interface{}(_float64)), unsafe.Sizeof(interface{}(&_float64)))
	fmt.Fprintf(w, s, "complex64", _complex64, unsafe.Sizeof(_complex64), unsafe.Sizeof(&_complex64), unsafe.Sizeof(interface{}(_complex64)), unsafe.Sizeof(interface{}(&_complex64)))
	fmt.Fprintf(w, s, "complex128", _complex128, unsafe.Sizeof(_complex128), unsafe.Sizeof(&_complex128), unsafe.Sizeof(interface{}(_complex128)), unsafe.Sizeof(interface{}(&_complex128)))
	fmt.Fprintf(w, s, "byte", _byte, unsafe.Sizeof(_byte), unsafe.Sizeof(&_byte), unsafe.Sizeof(interface{}(_byte)), unsafe.Sizeof(interface{}(&_byte)))
	fmt.Fprintf(w, s, "bool", _bool, unsafe.Sizeof(_bool), unsafe.Sizeof(&_bool), unsafe.Sizeof(interface{}(_bool)), unsafe.Sizeof(interface{}(&_bool)))
	fmt.Fprintf(w, s, "rune", _rune, unsafe.Sizeof(_rune), unsafe.Sizeof(&_rune), unsafe.Sizeof(interface{}(_rune)), unsafe.Sizeof(interface{}(&_rune)))
	fmt.Fprintf(w, s, "string", _string, unsafe.Sizeof(_string), unsafe.Sizeof(&_string), unsafe.Sizeof(interface{}(_string)), unsafe.Sizeof(interface{}(&_string)))
	fmt.Fprintf(w, s, "struct_int", _struct_int, unsafe.Sizeof(_struct_int), unsafe.Sizeof(&_struct_int), unsafe.Sizeof(interface{}(_struct_int)), unsafe.Sizeof(interface{}(&_struct_int)))
	fmt.Fprintf(w, s, "struct_int8", _struct_int8, unsafe.Sizeof(_struct_int8), unsafe.Sizeof(&_struct_int8), unsafe.Sizeof(interface{}(_struct_int8)), unsafe.Sizeof(interface{}(&_struct_int8)))
	fmt.Fprintf(w, s, "struct_int16", _struct_int16, unsafe.Sizeof(_struct_int16), unsafe.Sizeof(&_struct_int16), unsafe.Sizeof(interface{}(_struct_int16)), unsafe.Sizeof(interface{}(&_struct_int16)))
	fmt.Fprintf(w, s, "struct_int32", _struct_int32, unsafe.Sizeof(_struct_int32), unsafe.Sizeof(&_struct_int32), unsafe.Sizeof(interface{}(_struct_int32)), unsafe.Sizeof(interface{}(&_struct_int32)))
	fmt.Fprintf(w, s, "struct_int64", _struct_int64, unsafe.Sizeof(_struct_int64), unsafe.Sizeof(&_struct_int64), unsafe.Sizeof(interface{}(_struct_int64)), unsafe.Sizeof(interface{}(&_struct_int64)))
	fmt.Fprintf(w, s, "struct_uint", _struct_uint, unsafe.Sizeof(_struct_uint), unsafe.Sizeof(&_struct_uint), unsafe.Sizeof(interface{}(_struct_uint)), unsafe.Sizeof(interface{}(&_struct_uint)))
	fmt.Fprintf(w, s, "struct_uint8", _struct_uint8, unsafe.Sizeof(_struct_uint8), unsafe.Sizeof(&_struct_uint8), unsafe.Sizeof(interface{}(_struct_uint8)), unsafe.Sizeof(interface{}(&_struct_uint8)))
	fmt.Fprintf(w, s, "struct_uint16", _struct_uint16, unsafe.Sizeof(_struct_uint16), unsafe.Sizeof(&_struct_uint16), unsafe.Sizeof(interface{}(_struct_uint16)), unsafe.Sizeof(interface{}(&_struct_uint16)))
	fmt.Fprintf(w, s, "struct_uint32", _struct_uint32, unsafe.Sizeof(_struct_uint32), unsafe.Sizeof(&_struct_uint32), unsafe.Sizeof(interface{}(_struct_uint32)), unsafe.Sizeof(interface{}(&_struct_uint32)))
	fmt.Fprintf(w, s, "struct_uint64", _struct_uint64, unsafe.Sizeof(_struct_uint64), unsafe.Sizeof(&_struct_uint64), unsafe.Sizeof(interface{}(_struct_uint64)), unsafe.Sizeof(interface{}(&_struct_uint64)))
	fmt.Fprintf(w, s, "struct_float32", _struct_float32, unsafe.Sizeof(_struct_float32), unsafe.Sizeof(&_struct_float32), unsafe.Sizeof(interface{}(_struct_float32)), unsafe.Sizeof(interface{}(&_struct_float32)))
	fmt.Fprintf(w, s, "struct_float64", _struct_float64, unsafe.Sizeof(_struct_float64), unsafe.Sizeof(&_struct_float64), unsafe.Sizeof(interface{}(_struct_float64)), unsafe.Sizeof(interface{}(&_struct_float64)))
	fmt.Fprintf(w, s, "struct_complex64", _struct_complex64, unsafe.Sizeof(_struct_complex64), unsafe.Sizeof(&_struct_complex64), unsafe.Sizeof(interface{}(_struct_complex64)), unsafe.Sizeof(interface{}(&_struct_complex64)))
	fmt.Fprintf(w, s, "struct_complex128", _struct_complex128, unsafe.Sizeof(_struct_complex128), unsafe.Sizeof(&_struct_complex128), unsafe.Sizeof(interface{}(_struct_complex128)), unsafe.Sizeof(interface{}(&_struct_complex128)))
	fmt.Fprintf(w, s, "struct_byte", _struct_byte, unsafe.Sizeof(_struct_byte), unsafe.Sizeof(&_struct_byte), unsafe.Sizeof(interface{}(_struct_byte)), unsafe.Sizeof(interface{}(&_struct_byte)))
	fmt.Fprintf(w, s, "struct_bool", _struct_bool, unsafe.Sizeof(_struct_bool), unsafe.Sizeof(&_struct_bool), unsafe.Sizeof(interface{}(_struct_bool)), unsafe.Sizeof(interface{}(&_struct_bool)))
	fmt.Fprintf(w, s, "struct_rune", _struct_rune, unsafe.Sizeof(_struct_rune), unsafe.Sizeof(&_struct_rune), unsafe.Sizeof(interface{}(_struct_rune)), unsafe.Sizeof(interface{}(&_struct_rune)))
	fmt.Fprintf(w, s, "struct_string", _struct_string, unsafe.Sizeof(_struct_string), unsafe.Sizeof(&_struct_string), unsafe.Sizeof(interface{}(_struct_string)), unsafe.Sizeof(interface{}(&_struct_string)))
	fmt.Fprintf(w, s, "struct_int32_int32", _struct_int32_int32, unsafe.Sizeof(_struct_int32_int32), unsafe.Sizeof(&_struct_int32_int32), unsafe.Sizeof(interface{}(_struct_int32_int32)), unsafe.Sizeof(interface{}(&_struct_int32_int32)))
	fmt.Fprintf(w, s, "struct_int32_int64", _struct_int32_int64, unsafe.Sizeof(_struct_int32_int64), unsafe.Sizeof(&_struct_int32_int64), unsafe.Sizeof(interface{}(_struct_int32_int64)), unsafe.Sizeof(interface{}(&_struct_int32_int64)))
	fmt.Fprintf(w, s, "struct_array_bytes_7", _struct_array_bytes_7, unsafe.Sizeof(_struct_array_bytes_7), unsafe.Sizeof(&_struct_array_bytes_7), unsafe.Sizeof(interface{}(_struct_array_bytes_7)), unsafe.Sizeof(interface{}(&_struct_array_bytes_7)))
	fmt.Fprintf(w, s, "struct_byte_7", _struct_byte_7, unsafe.Sizeof(_struct_byte_7), unsafe.Sizeof(&_struct_byte_7), unsafe.Sizeof(interface{}(_struct_byte_7)), unsafe.Sizeof(interface{}(&_struct_byte_7)))
	w.Flush()
}
