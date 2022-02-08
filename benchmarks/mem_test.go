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

import "testing"

func BenchmarkMem(b *testing.B) {

	b.Run("int", func(b *testing.B) {
		b.Logf("real(T)=%T", _int)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _int
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_int_benchmark = _int
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _int_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_int_benchmark = _int_n
				}
			})
		})
	})

	b.Run("int8", func(b *testing.B) {
		b.Logf("real(T)=%T", _int8)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _int8
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_int8_benchmark = _int8
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _int8_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_int8_benchmark = _int8_n
				}
			})
		})
	})

	b.Run("int16", func(b *testing.B) {
		b.Logf("real(T)=%T", _int16)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _int16
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_int16_benchmark = _int16
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _int16_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_int16_benchmark = _int16_n
				}
			})
		})
	})

	b.Run("int32", func(b *testing.B) {
		b.Logf("real(T)=%T", _int32)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _int32
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_int32_benchmark = _int32
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _int32_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_int32_benchmark = _int32_n
				}
			})
		})
	})

	b.Run("int64", func(b *testing.B) {
		b.Logf("real(T)=%T", _int64)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _int64
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_int64_benchmark = _int64
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _int64_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_int64_benchmark = _int64_n
				}
			})
		})
	})

	b.Run("uint", func(b *testing.B) {
		b.Logf("real(T)=%T", _uint)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _uint
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_uint_benchmark = _uint
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _uint_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_uint_benchmark = _uint_n
				}
			})
		})
	})

	b.Run("uint8", func(b *testing.B) {
		b.Logf("real(T)=%T", _uint8)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _uint8
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_uint8_benchmark = _uint8
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _uint8_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_uint8_benchmark = _uint8_n
				}
			})
		})
	})

	b.Run("uint16", func(b *testing.B) {
		b.Logf("real(T)=%T", _uint16)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _uint16
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_uint16_benchmark = _uint16
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _uint16_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_uint16_benchmark = _uint16_n
				}
			})
		})
	})

	b.Run("uint32", func(b *testing.B) {
		b.Logf("real(T)=%T", _uint32)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _uint32
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_uint32_benchmark = _uint32
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _uint32_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_uint32_benchmark = _uint32_n
				}
			})
		})
	})

	b.Run("uint64", func(b *testing.B) {
		b.Logf("real(T)=%T", _uint64)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _uint64
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_uint64_benchmark = _uint64
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _uint64_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_uint64_benchmark = _uint64_n
				}
			})
		})
	})

	b.Run("float32", func(b *testing.B) {
		b.Logf("real(T)=%T", _float32)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _float32
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_float32_benchmark = _float32
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _float32_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_float32_benchmark = _float32_n
				}
			})
		})
	})

	b.Run("float64", func(b *testing.B) {
		b.Logf("real(T)=%T", _float64)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _float64
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_float64_benchmark = _float64
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _float64_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_float64_benchmark = _float64_n
				}
			})
		})
	})

	b.Run("complex64", func(b *testing.B) {
		b.Logf("real(T)=%T", _complex64)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _complex64
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_complex64_benchmark = _complex64
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _complex64_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_complex64_benchmark = _complex64_n
				}
			})
		})
	})

	b.Run("complex128", func(b *testing.B) {
		b.Logf("real(T)=%T", _complex128)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _complex128
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_complex128_benchmark = _complex128
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _complex128_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_complex128_benchmark = _complex128_n
				}
			})
		})
	})

	b.Run("byte", func(b *testing.B) {
		b.Logf("real(T)=%T", _byte)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _byte
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_byte_benchmark = _byte
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _byte_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_byte_benchmark = _byte_n
				}
			})
		})
	})

	b.Run("bool", func(b *testing.B) {
		b.Logf("real(T)=%T", _bool)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _bool
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_bool_benchmark = _bool
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _bool_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_bool_benchmark = _bool_n
				}
			})
		})
	})

	b.Run("rune", func(b *testing.B) {
		b.Logf("real(T)=%T", _rune)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _rune
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_rune_benchmark = _rune
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _rune_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_rune_benchmark = _rune_n
				}
			})
		})
	})

	b.Run("string", func(b *testing.B) {
		b.Logf("real(T)=%T", _string)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _string
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_string_benchmark = _string
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _string_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_string_benchmark = _string_n
				}
			})
		})
	})

	b.Run("struct_int", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_int)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int_benchmark = _struct_int
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int_benchmark = _struct_int_n
				}
			})
		})
	})

	b.Run("struct_int8", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_int8)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int8
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int8_benchmark = _struct_int8
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int8_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int8_benchmark = _struct_int8_n
				}
			})
		})
	})

	b.Run("struct_int16", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_int16)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int16
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int16_benchmark = _struct_int16
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int16_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int16_benchmark = _struct_int16_n
				}
			})
		})
	})

	b.Run("struct_int32", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_int32)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int32
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int32_benchmark = _struct_int32
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int32_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int32_benchmark = _struct_int32_n
				}
			})
		})
	})

	b.Run("struct_int64", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_int64)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int64
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int64_benchmark = _struct_int64
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int64_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int64_benchmark = _struct_int64_n
				}
			})
		})
	})

	b.Run("struct_uint", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_uint)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_uint
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_uint_benchmark = _struct_uint
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_uint_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_uint_benchmark = _struct_uint_n
				}
			})
		})
	})

	b.Run("struct_uint8", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_uint8)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_uint8
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_uint8_benchmark = _struct_uint8
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_uint8_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_uint8_benchmark = _struct_uint8_n
				}
			})
		})
	})

	b.Run("struct_uint16", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_uint16)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_uint16
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_uint16_benchmark = _struct_uint16
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_uint16_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_uint16_benchmark = _struct_uint16_n
				}
			})
		})
	})

	b.Run("struct_uint32", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_uint32)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_uint32
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_uint32_benchmark = _struct_uint32
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_uint32_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_uint32_benchmark = _struct_uint32_n
				}
			})
		})
	})

	b.Run("struct_uint64", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_uint64)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_uint64
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_uint64_benchmark = _struct_uint64
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_uint64_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_uint64_benchmark = _struct_uint64_n
				}
			})
		})
	})

	b.Run("struct_float32", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_float32)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_float32
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_float32_benchmark = _struct_float32
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_float32_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_float32_benchmark = _struct_float32_n
				}
			})
		})
	})

	b.Run("struct_float64", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_float64)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_float64
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_float64_benchmark = _struct_float64
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_float64_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_float64_benchmark = _struct_float64_n
				}
			})
		})
	})

	b.Run("struct_complex64", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_complex64)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_complex64
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_complex64_benchmark = _struct_complex64
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_complex64_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_complex64_benchmark = _struct_complex64_n
				}
			})
		})
	})

	b.Run("struct_complex128", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_complex128)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_complex128
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_complex128_benchmark = _struct_complex128
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_complex128_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_complex128_benchmark = _struct_complex128_n
				}
			})
		})
	})

	b.Run("struct_byte", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_byte)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_byte
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_byte_benchmark = _struct_byte
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_byte_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_byte_benchmark = _struct_byte_n
				}
			})
		})
	})

	b.Run("struct_bool", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_bool)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_bool
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_bool_benchmark = _struct_bool
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_bool_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_bool_benchmark = _struct_bool_n
				}
			})
		})
	})

	b.Run("struct_rune", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_rune)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_rune
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_rune_benchmark = _struct_rune
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_rune_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_rune_benchmark = _struct_rune_n
				}
			})
		})
	})

	b.Run("struct_string", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_string)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_string
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_string_benchmark = _struct_string
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_string_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_string_benchmark = _struct_string_n
				}
			})
		})
	})

	b.Run("struct_int32_int32", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_int32_int32)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int32_int32
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int32_int32_benchmark = _struct_int32_int32
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int32_int32_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int32_int32_benchmark = _struct_int32_int32_n
				}
			})
		})
	})

	b.Run("struct_int32_int64", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_int32_int64)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int32_int64
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int32_int64_benchmark = _struct_int32_int64
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_int32_int64_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_int32_int64_benchmark = _struct_int32_int64_n
				}
			})
		})
	})

	b.Run("struct_array_bytes_7", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_array_bytes_7)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_array_bytes_7
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_array_bytes_7_benchmark = _struct_array_bytes_7
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_array_bytes_7_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_array_bytes_7_benchmark = _struct_array_bytes_7_n
				}
			})
		})
	})

	b.Run("struct_byte_7", func(b *testing.B) {
		b.Logf("real(T)=%T", _struct_byte_7)
		b.Run("0", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_byte_7
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_byte_7_benchmark = _struct_byte_7
				}
			})
		})
		b.Run("n", func(b *testing.B) {
			b.Run("h", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_i = _struct_byte_7_n
				}
			})
			b.Run("s", func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					_struct_byte_7_benchmark = _struct_byte_7_n
				}
			})
		})
	})
}
