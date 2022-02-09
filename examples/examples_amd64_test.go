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

package examples_test

func ExampleEscape() {
	compile("escape.go")
	// Output:
	// escape.go:24:4: x escapes to heap
	// escape.go:28:10: new(int64) escapes to heap
	// escape.go:35:10: new(int64) escapes to heap
}

func ExampleLeak() {
	compile("leak.go")
	// Output:
	// leak.go:21:18: leaking param: p
	// leak.go:26:18: leaking param: p to result ~r1 level=0
}
