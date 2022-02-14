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

package mem_test

var (
	_int   int
	_int8  int8
	_int16 int16
	_int32 int32
	_int64 int64

	_uint   uint
	_uint8  uint8
	_uint16 uint16
	_uint32 uint32
	_uint64 uint64

	_float32 float32
	_float64 float64

	_complex64  complex64
	_complex128 complex128

	_byte   byte
	_bool   bool
	_rune   rune
	_string string

	_struct_int   struct{ a int }
	_struct_int8  struct{ a int8 }
	_struct_int16 struct{ a int16 }
	_struct_int32 struct{ a int32 }
	_struct_int64 struct{ a int64 }

	_struct_uint   struct{ a uint }
	_struct_uint8  struct{ a uint8 }
	_struct_uint16 struct{ a uint16 }
	_struct_uint32 struct{ a uint32 }
	_struct_uint64 struct{ a uint64 }

	_struct_float32 struct{ a float32 }
	_struct_float64 struct{ a float64 }

	_struct_complex64  struct{ a complex64 }
	_struct_complex128 struct{ a complex128 }

	_struct_byte   struct{ a byte }
	_struct_bool   struct{ a bool }
	_struct_rune   struct{ a rune }
	_struct_string struct{ a string }

	_struct_int32_int32 struct{ a, b int32 }
	_struct_int32_int64 struct {
		a int32
		b int64
	}
	_struct_array_bytes_7 struct{ a [7]byte }
	_struct_byte_7        struct{ a, b, c, d, e, f, g byte }

	_int_n   int   = nonZeroRandInt(_int_size)
	_int8_n  int8  = int8(nonZeroRandInt(8))
	_int16_n int16 = int16(nonZeroRandInt(16))
	_int32_n int32 = int32(nonZeroRandInt(32))
	_int64_n int64 = int64(nonZeroRandInt(64))

	_uint_n   uint   = uint(nonZeroRandInt(_int_size))
	_uint8_n  uint8  = uint8(nonZeroRandInt(8))
	_uint16_n uint16 = uint16(nonZeroRandInt(16))
	_uint32_n uint32 = uint32(nonZeroRandInt(32))
	_uint64_n uint64 = uint64(nonZeroRandInt(64))

	_float32_n float32 = float32(nonZeroRandInt(32))
	_float64_n float64 = float64(nonZeroRandInt(64))

	_complex64_n  complex64  = complex(float32(nonZeroRandInt(32)), float32(nonZeroRandInt(32)))
	_complex128_n complex128 = complex(float64(nonZeroRandInt(64)), float64(nonZeroRandInt(64)))

	_byte_n   byte   = byte(nonZeroRandInt(8))
	_bool_n   bool   = nonConstBoolTrue()
	_rune_n   rune   = rune(nonZeroRandInt(32))
	_string_n string = nonZeroString(50)

	_struct_int_n   = struct{ a int }{a: nonZeroRandInt(_int_size)}
	_struct_int8_n  = struct{ a int8 }{a: int8(nonZeroRandInt(8))}
	_struct_int16_n = struct{ a int16 }{a: int16(nonZeroRandInt(16))}
	_struct_int32_n = struct{ a int32 }{a: int32(nonZeroRandInt(32))}
	_struct_int64_n = struct{ a int64 }{a: int64(nonZeroRandInt(64))}

	_struct_uint_n   = struct{ a uint }{a: uint(nonZeroRandInt(_int_size))}
	_struct_uint8_n  = struct{ a uint8 }{a: uint8(nonZeroRandInt(8))}
	_struct_uint16_n = struct{ a uint16 }{a: uint16(nonZeroRandInt(16))}
	_struct_uint32_n = struct{ a uint32 }{a: uint32(nonZeroRandInt(32))}
	_struct_uint64_n = struct{ a uint64 }{a: uint64(nonZeroRandInt(64))}

	_struct_float32_n = struct{ a float32 }{a: float32(nonZeroRandInt(32))}
	_struct_float64_n = struct{ a float64 }{a: float64(nonZeroRandInt(64))}

	_struct_complex64_n  = struct{ a complex64 }{a: complex(float32(nonZeroRandInt(32)), float32(nonZeroRandInt(32)))}
	_struct_complex128_n = struct{ a complex128 }{a: complex(float64(nonZeroRandInt(64)), float64(nonZeroRandInt(64)))}

	_struct_byte_n   = struct{ a byte }{a: byte(nonZeroRandInt(8))}
	_struct_bool_n   = struct{ a bool }{a: nonConstBoolTrue()}
	_struct_rune_n   = struct{ a rune }{a: rune(nonZeroRandInt(32))}
	_struct_string_n = struct{ a string }{a: nonZeroString(50)}

	_struct_int32_int32_n = struct{ a, b int32 }{a: int32(nonZeroRandInt(32)), b: int32(nonZeroRandInt(32))}
	_struct_int32_int64_n = struct {
		a int32
		b int64
	}{a: int32(nonZeroRandInt(32)), b: int64(nonZeroRandInt(64))}
	_struct_array_bytes_7_n = struct{ a [7]byte }{a: [7]byte{byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8))}}
	_struct_byte_7_n        = struct{ a, b, c, d, e, f, g byte }{a: byte(nonZeroRandInt(8)), b: byte(nonZeroRandInt(8)), c: byte(nonZeroRandInt(8)), d: byte(nonZeroRandInt(8)), e: byte(nonZeroRandInt(8)), f: byte(nonZeroRandInt(8)), g: byte(nonZeroRandInt(8))}

	_int_benchmark   int
	_int8_benchmark  int8
	_int16_benchmark int16
	_int32_benchmark int32
	_int64_benchmark int64

	_uint_benchmark   uint
	_uint8_benchmark  uint8
	_uint16_benchmark uint16
	_uint32_benchmark uint32
	_uint64_benchmark uint64

	_float32_benchmark float32
	_float64_benchmark float64

	_complex64_benchmark  complex64
	_complex128_benchmark complex128

	_byte_benchmark   byte
	_bool_benchmark   bool
	_rune_benchmark   rune
	_string_benchmark string

	_struct_int_benchmark   struct{ a int }
	_struct_int8_benchmark  struct{ a int8 }
	_struct_int16_benchmark struct{ a int16 }
	_struct_int32_benchmark struct{ a int32 }
	_struct_int64_benchmark struct{ a int64 }

	_struct_uint_benchmark   struct{ a uint }
	_struct_uint8_benchmark  struct{ a uint8 }
	_struct_uint16_benchmark struct{ a uint16 }
	_struct_uint32_benchmark struct{ a uint32 }
	_struct_uint64_benchmark struct{ a uint64 }

	_struct_float32_benchmark struct{ a float32 }
	_struct_float64_benchmark struct{ a float64 }

	_struct_complex64_benchmark  struct{ a complex64 }
	_struct_complex128_benchmark struct{ a complex128 }

	_struct_byte_benchmark   struct{ a byte }
	_struct_bool_benchmark   struct{ a bool }
	_struct_rune_benchmark   struct{ a rune }
	_struct_string_benchmark struct{ a string }

	_struct_int32_int32_benchmark struct{ a, b int32 }
	_struct_int32_int64_benchmark struct {
		a int32
		b int64
	}
	_struct_array_bytes_7_benchmark struct{ a [7]byte }
	_struct_byte_7_benchmark        struct{ a, b, c, d, e, f, g byte }
)
