//go:generate python3 ../hack/gen.py

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

// Package benchmarks_test is used to benchmark the possible
// mallocs that can occur when storing interface values.
//
// Print the size of the exemplar, predefined types and structs using the
// following command:
//
//     go test
//
// Run the benchmark with the following command:
//
//     go test -bench . -run Box -benchmem -count 1
package benchmarks_test

import (
	"math/rand"
	"strconv"
)

// nonZeroRandInt returns an integer between zero and the max size for the
// integer at the specified size/width (ex. 8, 16, 32, 64).
//
// Please note this function is also used with the other number types and
// the result cast into the appropriate type.
//
// For unsigned integers this is because the math/rand package does not
// implement UintN. Using Uint32 or Uint64 to try and get random values for
// uint8 and uint16 can cause stack overflows before a value that fits into the
// width of those types is found.
//
// For floating point numbers this is just convienent as math/rand *does*
// provide Float32() and Float64() to generate pseudo-random numbers for those
// types.
//
// For complex numbers this is because the math/rand package does not offer a
// random generation function for complex numbers.
func nonZeroRandInt(n int) int {
	if i := rand.Intn(1<<(n-1) - 1); i > 0 {
		return i
	}
	return nonZeroRandInt(n)
}

// nonConstBoolTrue returns a boolean true value that is not a constant.
func nonConstBoolTrue() bool {
	b, _ := strconv.ParseBool(strconv.Itoa(nonZeroRandInt(8)))
	return b
}

// nonZeroString returns a string of random characters at the specified length.
func nonZeroString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = _letters[rand.Intn(len(_letters))]
	}
	return string(b)
}

const _int_size = 32 << (^uint(0) >> 63)

var (
	_i       interface{}
	_letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)
