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

package examples

var i interface{}

func escape() {
	var x int64
	x = 4096
	i = x
}

func escape2chan() {
	p := new(int64)
	c := make(chan *int64, 1)
	c <- p
	close(c)
}

func escape2slice() {
	p := new(int64)
	var slice []*int64
	_ = append(slice, p)
}
