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

package main

// ex1 illustrates a stack frame for a function with two local variables,
// an int64 and interface. The stack frame also includes an additional
// eight bytes of memory for the value stored in the interface.
func ex1() {
	var x int64
	var i interface{}
	x = 2
	i = x
	_ = i
}
