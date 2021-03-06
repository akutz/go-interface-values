# Overall impact

The previous page reviews which types are subject to optimizations Go uses when storing values in interfaces, but did not actually detail the impact.  Well, there are a few criteria reviewed on the previous page, but here it is again for convenience:

* Zero values, including `0`, `nil`, and an empty string `""` all qualify.
* Any value that is type which is a single byte wide, such as `bool`, `int8`, and `uint8`.
* Any integer type with a value that is in the inclusive range of 0-255.
* In some cases if `T` is subject to an optimization, so too will `struct{a T}`.

In order to produce a more comprehensive dataset for when this behavior _could_ occur, we can run the following command:

```bash
docker run -it --rm go-interface-values:latest \
  bash -c 'go test -v -count 1 -benchtime 1000x \
  -run Mem -benchmem -bench BenchmarkMem ./tests/mem | \
  python3 hack/b2md.py --no-echo'
```

The output will be a markdown table that prints the:

* friendly name of the type
* actual type as formatted with `%T`
* number of bytes allocated on the heap to:
    * copy a zero value to another variable of the same type
    * store a zero value in an interface
    * copy a non-zero value to another variable of the same type
    * store a non-zero value in an interface

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

The above table clearly aligns with the findings in this repository, that the following values stored in interfaces do not result in memory allocated on the heap:

* zero values, ex. `0`, `nil`, and `""`
* single byte-wide types, ex. `byte`, `bool`, `int8`, and `uint8`
* non-zero, numeric values between and including 0-255 for types:
    * `int`, `int8`, `int16`, `int32`, `int64`
    * `uint`, `uint8`, `uint16`, `uint32`, `iint64`
    * ~~`float32`, `float64`~~

        ---

        It is unclear to me why the types `float32` and `float64` are not subject to this optimization, but tests show they are not. The functions `convT32` and `convT64` are used for these types when storing values in interfaces, and those functions have the optimization logic.

        I intend to raise this in Gopher Slack and with a GitHub issue, and I will link those here.

        ---

* zero values for `struct{a T}` where `T` is:
    * `int`, `int16`, `int32`, `int64`
    * `uint`, `uint16`, `uint32`, `uint64`
    * `float32`, `float64`
    * `rune` (which has an underlying type of `int32`)
    * `string`

---

Next: [Lessons learned](../05-lessons-learned/)
