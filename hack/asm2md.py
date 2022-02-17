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

"""asm2md

Transforms Go compiler assembly into a markdown table.

    asm2md.py FILE
    cat FILE | asm2md.py
"""

import argparse
import platform
import re
import sys

parser = argparse.ArgumentParser(
    description="asm2md", formatter_class=argparse.ArgumentDefaultsHelpFormatter
)

parser.add_argument(
    "--no-echo",
    dest="no_echo",
    action="store_true",
    help="do not echo stdin to stdout",
)

parser.add_argument(
    "file",
    metavar="FILE",
    nargs="?",
    help="path to an input file",
)

args = parser.parse_args()
if not args:
    parser.print_help()
    sys.exit(1)


if args.file:
    echo = False
    f = open(args.file, "r")
else:
    echo = not args.no_echo
    f = sys.stdin

rx = re.compile(
    r"\((bench_test\.go:(\d+)\))\s(\w?MOV\w*|LEAQ?)\s+\S+\.(_[\w\d_]+)\(SB\).*\1\s(CALL|(?:LEAQ?)|(?:\w?MOV\w*))\s(runtime\.[\w\d]+)\(SB\)",
    re.S,  # singleline
)


x86_asm_op_links = {
    "CALL": "https://www.felixcloutier.com/x86/call",
    "LEA": "https://www.felixcloutier.com/x86/lea",
    "MOV": "https://www.felixcloutier.com/x86/mov",
}

arm_asm_op_links = {
    "CALL": "https://developer.arm.com/documentation/ddi0602/2021-12/Base-Instructions/BLR--Branch-with-Link-to-Register-?lang=en",
}

go_func_links = {
    "convT16": "https://github.com/golang/go/blob/41f485b9a7d8fd647c415be1d11b612063dff21c/src/runtime/iface.go#L352-L363",
    "convT32": "https://github.com/golang/go/blob/41f485b9a7d8fd647c415be1d11b612063dff21c/src/runtime/iface.go#L365-L376",
    "convT64": "https://github.com/golang/go/blob/41f485b9a7d8fd647c415be1d11b612063dff21c/src/runtime/iface.go#L378-L386",
    "convTstring": "https://github.com/golang/go/blob/41f485b9a7d8fd647c415be1d11b612063dff21c/src/runtime/iface.go#L388-L396",
    "convTnoptr": "https://github.com/golang/go/blob/41f485b9a7d8fd647c415be1d11b612063dff21c/src/runtime/iface.go#L335-L350",
    "staticuint64s": "https://github.com/golang/go/blob/41f485b9a7d8fd647c415be1d11b612063dff21c/src/runtime/iface.go#L492-L526",
}


def get_arm_asm_op_link(s):
    key = None
    if "CALL" in s:
        key = "CALL"
    if not key or key not in arm_asm_op_links:
        return s
    return "[{}]({})".format(s, arm_asm_op_links[key])


def get_x86_asm_op_link(s):
    key = None
    if "CALL" in s:
        key = "CALL"
    elif "LEA" in s:
        key = "LEA"
    if not key or key not in x86_asm_op_links:
        return s
    return "[{}]({})".format(s, x86_asm_op_links[key])


def get_asm_op_link(s):
    if "arm" in platform.processor():
        return get_arm_asm_op_link(s)
    return get_x86_asm_op_link(s)


def get_go_func_link(s):
    key = None
    if "convT16" in s:
        key = "convT16"
    elif "convT32" in s:
        key = "convT32"
    elif "convT64" in s:
        key = "convT64"
    elif "convTstring" in s:
        key = "convTstring"
    elif "convTnoptr" in s:
        key = "convTnoptr"
    elif "staticuint64s" in s:
        key = "staticuint64s"
    if not key:
        return s
    return "[{}]({})".format(s, go_func_links[key])


"""
data = [
    {
        "type": "_int64",

        "0": {
            "lino": "_test.go:248",
            "load_asm": "MOVQ",
            "stor_asm": "CALL",
            "stor_go":  "runtime.convT64",
        },

        # same as "0"
        "n": {},
    },

    # repeat for other types
]
"""
data = []

# Read the entire file at once. I really wanted to avoid this and read it
# a line at a time, but parsing ASM where Go source is not contiguous just
# proved too difficult. This approach enables the use of a single-line regex
# with back references to find what we need.
lines = f.read()
for m in rx.finditer(lines):

    lino = m.group(2)

    # This incredibly inefficient, but I was having trouble getting the correct
    # regex to eliminate the blocks of ASM that print the type.
    if re.search(r":" + lino + r"\).*go.string.\"real\(T\)=%T\"\(SB\)", lines):
        continue

    load_asm = m.group(3)
    type = m.group(4)
    stor_asm = m.group(5)
    stor_go = m.group(6)

    if type.endswith("_n"):
        data[len(data) - 1]["n"] = {
            "lino": lino,
            "load_asm": get_asm_op_link(load_asm),
            "stor_asm": get_asm_op_link(stor_asm),
            "stor_go": get_go_func_link(stor_go),
        }
    else:
        data.append(
            {
                "type": type.removeprefix("_"),
                "0": {
                    "lino": lino,
                    "load_asm": get_asm_op_link(load_asm),
                    "stor_asm": get_asm_op_link(stor_asm),
                    "stor_go": get_go_func_link(stor_go),
                },
                "n": {
                    "stor_asm": "NA",
                    "stor_go": "NA",
                },
            }
        )

if echo:
    print(lines)

# Go ahead and close the file.
f.close()

print(
    "| Type | Line no | Op to store zero value in asm | ...in go | Op to store non-zero, random value in asm | ...in go |"
)
print(
    "|:----:|:-------:|:-----------------------------:|:--------:|:-----------------------------------------:|:--------:|"
)

s = "| `{}` | {} | {} | {} | {} | {} |"
for o in data:
    print(
        s.format(
            o["type"],
            o["0"]["lino"],
            o["0"]["stor_asm"],
            o["0"]["stor_go"],
            o["n"]["stor_asm"],
            o["n"]["stor_go"],
        )
    )
