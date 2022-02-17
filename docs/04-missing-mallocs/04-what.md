# Overall impact



| Type | `%T` | Bytes to store zero value | ....in an interface | Bytes to store non-zero, random value | ...in an interface |
|:----:|:----:|:-------------------------:|:-------------------:|:-------------------------------------:|:------------------:|
| int | `int` | 0 | 0 | 0 | 8 |
| int8 | `int8` | 0 | 0 | 0 | 0 |
| int16 | `int16` | 0 | 0 | 0 | 2 |
| int32 | `int32` | 0 | 0 | 0 | 4 |
| int64 | `int64` | 0 | 0 | 0 | 8 |
| uint | `uint` | 0 | 0 | 0 | 8 |
| uint8 | `uint8` | 0 | 0 | 0 | 0 |
| uint16 | `uint16` | 0 | 0 | 0 | 2 |
| uint32 | `uint32` | 0 | 0 | 0 | 4 |
| uint64 | `uint64` | 0 | 0 | 0 | 8 |
| float32 | `float32` | 0 | 0 | 0 | 4 |
| float64 | `float64` | 0 | 0 | 0 | 8 |
| complex64 | `complex64` | 0 | 8 | 0 | 8 |
| complex128 | `complex128` | 0 | 16 | 0 | 16 |
| byte | `uint8` | 0 | 0 | 0 | 0 |
| bool | `bool` | 0 | 0 | 0 | 0 |
| rune | `int32` | 0 | 0 | 0 | 4 |
| string | `string` | 0 | 0 | 0 | 16 |
| struct_int | `struct { a int }` | 0 | 0 | 0 | 8 |
| struct_int8 | `struct { a int8 }` | 0 | 1 | 0 | 1 |
| struct_int16 | `struct { a int16 }` | 0 | 0 | 0 | 2 |
| struct_int32 | `struct { a int32 }` | 0 | 0 | 0 | 4 |
| struct_int64 | `struct { a int64 }` | 0 | 0 | 0 | 8 |
| struct_uint | `struct { a uint }` | 0 | 0 | 0 | 8 |
| struct_uint8 | `struct { a uint8 }` | 0 | 1 | 0 | 1 |
| struct_uint16 | `struct { a uint16 }` | 0 | 0 | 0 | 2 |
| struct_uint32 | `struct { a uint32 }` | 0 | 0 | 0 | 4 |
| struct_uint64 | `struct { a uint64 }` | 0 | 0 | 0 | 8 |
| struct_float32 | `struct { a float32 }` | 0 | 0 | 0 | 4 |
| struct_float64 | `struct { a float64 }` | 0 | 0 | 0 | 8 |
| struct_complex64 | `struct { a complex64 }` | 0 | 8 | 0 | 8 |
| struct_complex128 | `struct { a complex128 }` | 0 | 16 | 0 | 16 |
| struct_byte | `struct { a uint8 }` | 0 | 1 | 0 | 1 |
| struct_bool | `struct { a bool }` | 0 | 1 | 0 | 1 |
| struct_rune | `struct { a int32 }` | 0 | 0 | 0 | 4 |
| struct_string | `struct { a string }` | 0 | 0 | 0 | 16 |
| struct_int32_int32 | `struct { a int32; b int32 }` | 0 | 8 | 0 | 8 |
| struct_int32_int64 | `struct { a int32; b int64 }` | 0 | 16 | 0 | 16 |
| struct_array_bytes_7 | `struct { a [7]uint8 }` | 0 | 8 | 0 | 8 |
| struct_byte_7 | `struct { a uint8; b uint8; c uint8; d uint8; e uint8; f uint8; g uint8 }` | 0 | 8 | 0 | 8 |

---

Next: [Lessons learned](../05-lessons-learned/)
