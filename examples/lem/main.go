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

var lastLoginAttemptUsername *string

func Login(
	username *string, // leaking param
	password *string, // does not escape
	token interface{}, // leaking param to result ~r3 level=0
) interface{} {

	lastLoginAttemptUsername = username
	switch {
	case username != nil && password != nil:
		return "newLoginToken" // escapes to heap
	case token != nil:
		return token
	default:
		return nil
	}
}

func main() {
	var (
		username1 = "fake"               // moved to heap
		username2 = new(string)          // escapes to heap
		password  = "fake"               // does not escape
		cookieJar [15 * 1024 * 1024]byte // moved to heap
	)

	token := Login(&username1, &password, nil)
	if token != nil {
		_ = cookieJar
	}

	*username2 = username1
	Login(username2, nil, token)
}
