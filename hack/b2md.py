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

"""b2md

Transforms Go benchmark output into a markdown table.

    b2md.py FILE
    cat FILE | b2md.py
"""

import argparse
import re
import sys


def intOrFloat(s):
    if "." in s:
        return float(s)
    return int(s)


def twoDecimalFloat(f):
    return float("{:.2f}".format(f))


def noZedZed(f):
    c = str(f)
    if "." in c:
        if c.endswith(".0"):
            return int(c.replace(".0", ""))
        return twoDecimalFloat(float(c))
    return int(c)


parser = argparse.ArgumentParser(
    description="b2md", formatter_class=argparse.ArgumentDefaultsHelpFormatter
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

rx_name = re.compile(r"^.+real\(T\)=(.+)$")
rx_data = re.compile(
    r"^BenchmarkMem/([^/]+)/([^/]+)/([^-]+)-\d+\s+([\d.]+)\s+([\d.]+)\s+ns/op\s+([\d.]+)\s+B/op\s+([\d.]+) allocs/op$"
)

"""
data = [
    {
        "type": "int64",
        "0": {
            "h": {
                "bytes": 0,
                "alloc": 0,
            },

            # same as h
            "s": {}
        },

        # same as 0
        "n": {}
    },

    # repeat for other types
"""
data = []

for line in f:
    if echo:
        print(line, end="")

    m = rx_name.match(line)
    if m:
        data.append(
            {
                "T": m.group(1),
                "0": {
                    "h": {
                        "bytes": "",
                        "alloc": "",
                    },
                    "s": {
                        "bytes": "",
                        "alloc": "",
                    },
                },
                "n": {
                    "h": {
                        "bytes": "",
                        "alloc": "",
                    },
                    "s": {
                        "bytes": "",
                        "alloc": "",
                    },
                },
            }
        )
    else:
        m = rx_data.match(line)
        if not m:
            continue

        type = m.group(1)
        zorn = m.group(2)
        hors = m.group(3)
        bytesOp = noZedZed(m.group(6))
        allocOp = noZedZed(m.group(7))

        data[len(data) - 1]["type"] = type
        data[len(data) - 1][zorn][hors]["bytes"] = bytesOp
        data[len(data) - 1][zorn][hors]["alloc"] = allocOp

# Go ahead and close the file.
f.close()

print(
    "| Type | `%T` | Bytes to store zero value | ....in an interface | Bytes to store non-zero, random value | ...in an interface |"
)
print(
    "|:----:|:----:|:-------------------------:|:-------------------:|:-------------------------------------:|:------------------:|"
)

s = "| {} | `{}` | {} | {} | {} | {} |"
for o in data:
    print(
        s.format(
            o["type"],
            o["T"],
            o["0"]["s"]["bytes"],
            o["0"]["h"]["bytes"],
            o["n"]["s"]["bytes"],
            o["n"]["h"]["bytes"],
        )
    )
