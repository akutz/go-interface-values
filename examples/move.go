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

var _mp *int

func move2global() {
	var x int
	_mp = &x
}

func move2large() {
	var data [100 * 1024 * 1024]byte // 100 MiB
	_ = data
}

func move2chan() {
	var x int64
	c := make(chan *int64, 1)
	c <- &x
	close(c)
}

func move2slice() {
	var x int64
	var slice []*int64
	_ = append(slice, &x)
}
