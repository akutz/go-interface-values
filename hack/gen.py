#!/usr/bin/env python3

"""
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
"""

"""generate

Generates several of the go source files.
"""

import subprocess

_HEADER = """/*
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

"""

_INSTRUMENTED_TYPES = [
    {
        "type": "int",
        "nonz": "nonZeroRandInt(_int_size)",
    },
    {
        "type": "int8",
        "nonz": "int8(nonZeroRandInt(8))",
    },
    {
        "type": "int16",
        "nonz": "int16(nonZeroRandInt(16))",
    },
    {
        "type": "int32",
        "nonz": "int32(nonZeroRandInt(32))",
    },
    {
        "type": "int64",
        "nonz": "int64(nonZeroRandInt(64))",
        "nwln": 2,
    },
    {
        "type": "uint",
        "nonz": "uint(nonZeroRandInt(_int_size))",
    },
    {
        "type": "uint8",
        "nonz": "uint8(nonZeroRandInt(8))",
    },
    {
        "type": "uint16",
        "nonz": "uint16(nonZeroRandInt(16))",
    },
    {
        "type": "uint32",
        "nonz": "uint32(nonZeroRandInt(32))",
    },
    {
        "type": "uint64",
        "nonz": "uint64(nonZeroRandInt(64))",
        "nwln": 2,
    },
    {
        "type": "float32",
        "nonz": "float32(nonZeroRandInt(32))",
    },
    {
        "type": "float64",
        "nonz": "float64(nonZeroRandInt(64))",
        "nwln": 2,
    },
    {
        "type": "complex64",
        "nonz": "complex(float32(nonZeroRandInt(32)), float32(nonZeroRandInt(32)))",
    },
    {
        "type": "complex128",
        "nonz": "complex(float64(nonZeroRandInt(64)), float64(nonZeroRandInt(64)))",
        "nwln": 2,
    },
    {
        "type": "byte",
        "nonz": "byte(nonZeroRandInt(8))",
    },
    {
        "type": "bool",
        "nonz": "nonConstBoolTrue()",
    },
    {
        "type": "rune",
        "nonz": "rune(nonZeroRandInt(32))",
    },
    {
        "type": "string",
        "nonz": "nonZeroString(50)",
        "nwln": 2,
    },
    {
        "name": "struct_int32_int32",
        "type": "struct{ a, b int32 }",
        "nonz": "{ a: int32(nonZeroRandInt(32)), b: int32(nonZeroRandInt(32)) }",
    },
    {
        "name": "struct_int32_int64",
        "type": "struct { a int32; b int64 }",
        "nonz": "{ a: int32(nonZeroRandInt(32)), b: int64(nonZeroRandInt(64)) }",
    },
    {
        "name": "struct_array_bytes_7",
        "type": "struct{ a [7]byte }",
        "nonz": "{ a: [7]byte{byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8)), byte(nonZeroRandInt(8)) }}",
    },
    {
        "name": "struct_byte_7",
        "type": "struct{ a, b, c, d, e, f, g byte }",
        "nonz": "{ a: byte(nonZeroRandInt(8)), b: byte(nonZeroRandInt(8)), c: byte(nonZeroRandInt(8)), d: byte(nonZeroRandInt(8)), e: byte(nonZeroRandInt(8)), f: byte(nonZeroRandInt(8)), g: byte(nonZeroRandInt(8)) }",
        "nwln": 2,
    },
]

_TYPES_TEST_PATH = "types_test.go"
_PRINT_TEST_PATH = "print_test.go"
_MEM_TEST_PATH = "mem_test.go"

_is_struct = lambda t: t["type"].startswith("struct")
_non_struct_types = [t for t in _INSTRUMENTED_TYPES if not _is_struct(t)]
_struct_types = [t for t in _INSTRUMENTED_TYPES if _is_struct(t)]


def go_fmt(p):
    # Format the file.
    subprocess.run(
        ["go", "fmt", p],
        capture_output=True,
        check=True,
    )


def gen_types():
    def _print(
        f,
        it,
        is_zero=False,
        is_struct=False,
        is_wrapped=False,
        is_benchmark_global=False,
    ):
        t = it["type"]

        # Define the variable' name.
        f.write("\t_")
        if is_wrapped:
            f.write("struct_")
        f.write(t if "name" not in it else it["name"])
        if not is_zero:
            f.write("_n")
        if is_benchmark_global:
            f.write("_benchmark")
        f.write(" ")

        # The variables type declaration depends on whether this is a non-zero
        # value and if the type is a struct.
        if not is_zero and not is_benchmark_global and (is_struct or is_wrapped):
            f.write(" = ")

        # Define the variable' type.
        if is_wrapped:
            f.write("struct{ a ")
        f.write(t)
        if is_wrapped:
            f.write(" }")

        if not is_zero:
            # Define the variable's value.
            if not is_struct and not is_wrapped:
                f.write(" = ")

            if is_wrapped:
                f.write("{ a: ")

            f.write(it["nonz"])

            if is_wrapped:
                f.write(" }")

        f.write("\n" * (1 if "nwln" not in it else it["nwln"]))

    with open(_TYPES_TEST_PATH, "w") as f:
        f.write(_HEADER)
        f.write("var (\n")

        # Define the variables for the zero values for the non-struct types.
        for it in _non_struct_types:
            _print(f, it, is_zero=True)

        # Define the variables for the zero values for the non-struct types wrapped
        # by a struct.
        for it in _non_struct_types:
            _print(f, it, is_zero=True, is_wrapped=True)

        # Define the variables for the zero-values for the struct types.
        for it in _struct_types:
            _print(f, it, is_zero=True, is_struct=True)

        # Define the variables for the non-zero values for the non-struct types.
        for it in _non_struct_types:
            _print(f, it)

        # Define the variables for the non-zero values for the non-struct types
        # wrapped by a struct.
        for it in _non_struct_types:
            _print(f, it, is_wrapped=True)

        # Define the variables for the non-zero values for the struct types.
        for it in _struct_types:
            _print(f, it, is_struct=True)

        # Define the variables for the zero values for the non-struct types
        # used by the typed benchmarks.
        for it in _non_struct_types:
            _print(f, it, is_zero=True, is_benchmark_global=True)

        # Define the variables for the zero values for the non-struct types
        # wrapped by a struct and used by the typed benchmarks.
        for it in _non_struct_types:
            _print(f, it, is_zero=True, is_wrapped=True, is_benchmark_global=True)

        # Define the variables for the zero-values for the struct types
        # used by the typed benchmarks.
        for it in _struct_types:
            _print(f, it, is_zero=True, is_struct=True, is_benchmark_global=True)

        f.write(")\n\n")

    go_fmt(_TYPES_TEST_PATH)


def gen_print():
    s = 'fmt.Fprintf(w, s, "{}", _{}, unsafe.Sizeof(_{}), unsafe.Sizeof(&_{}), unsafe.Sizeof(interface{{}}(_{})), unsafe.Sizeof(interface{{}}(&_{})))\n'

    def _print(f, it, is_wrapped=False):
        t = it["type"] if "name" not in it else it["name"]
        if is_wrapped:
            t = "struct_" + t
        f.write(s.format(t, t, t, t, t, t))

    with open(_PRINT_TEST_PATH, "w") as f:
        f.write(_HEADER)
        f.write(
            """
        import (
            "fmt"
            "os"
            "testing"
            "text/tabwriter"
            "unsafe"
        )
        """
        )

        f.write(
            """
        func TestPrintSizes(t *testing.T) {
            w := tabwriter.NewWriter(os.Stdout, 4, 0, 1, ' ', tabwriter.Debug)
            fmt.Fprintf(w, "| T\\t real(T)\\t size(T)\\t size(*T)\\t size(any(T))\\t size(any(*T)) |\\n")
            fmt.Fprintf(w, "|:---:|:---:|:---:|:---:|:---:|:---:|\\n")
            const s = "| %s\\t `%T`\\t %d\\t %d\\t %d\\t %d |\\n"
        """
        )

        # Print the non-struct types.
        for it in _non_struct_types:
            _print(f, it)

        # Print the non-struct types wrapped by a struct.
        for it in _non_struct_types:
            _print(f, it, is_wrapped=True)

        # Print the struct types.
        for it in _struct_types:
            _print(f, it)

        f.write("w.Flush()\n}\n")

    go_fmt(_PRINT_TEST_PATH)


def gen_mem():

    s = """
    b.Run("{0}", func(b *testing.B) {{
        b.Logf("real(T)=%T", _{0})
        b.Run("0", func(b *testing.B) {{
            b.Run("h", func(b *testing.B) {{
                for j := 0; j < b.N; j++ {{
                    _i = _{0}
                }}
            }})
            b.Run("s", func(b *testing.B) {{
                for j := 0; j < b.N; j++ {{
                    _{0}_benchmark = _{0}
                }}
            }})
        }})
        b.Run("n", func(b *testing.B) {{
            b.Run("h", func(b *testing.B) {{
                for j := 0; j < b.N; j++ {{
                    _i = _{0}_n
                }}
            }})
            b.Run("s", func(b *testing.B) {{
                for j := 0; j < b.N; j++ {{
                    _{0}_benchmark = _{0}_n
                }}
            }})
        }})
    }})
    """

    def _print(f, it, is_wrapped=False):
        t = it["type"] if "name" not in it else it["name"]
        if is_wrapped:
            t = "struct_" + t
        f.write(s.format(t))

    with open(_MEM_TEST_PATH, "w") as f:
        f.write(_HEADER)
        f.write('\nimport "testing"\n')
        f.write("func BenchmarkMem(b *testing.B) {\n")

        # Benchmarks for the non-struct types.
        for it in _non_struct_types:
            _print(f, it)

        # Benchmarks for the non-struct types wrapped by a struct.
        for it in _non_struct_types:
            _print(f, it, is_wrapped=True)

        # Benchmarks for the struct types.
        for it in _struct_types:
            _print(f, it)

        f.write("}\n")

    go_fmt(_MEM_TEST_PATH)


gen_types()
gen_print()
gen_mem()
