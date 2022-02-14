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

import "os"

//go:noinline
func validateID(id *int32) *int32 { // leaking param: id to result ~r1 level=0
	return id
}

func main() {
	var id1 int32 = 4096
	if validateID(&id1) == nil {
		os.Exit(1)
	}

	var id2 *int32 = new(int32) // new(int32) does not escape
	*id2 = 4096
	validID := validateID(id2)
	if validID == nil {
		os.Exit(1)
	}
}
