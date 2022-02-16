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

package lem_test

import (
	"testing"

	"go-interface-values/tests/lem/lem"
)

func TestTestCaseTreeInsert(t *testing.T) {
	testCases := []struct {
		name  string
		tests []lem.TestCase
	}{
		{
			name: "a1",
			tests: []lem.TestCase{
				{
					ID:   "a1",
					Name: "",
				},
			},
		},
		{
			name: "a",
			tests: []lem.TestCase{
				{
					ID:   "a",
					Name: "",
				},
			},
		},
		{
			name: "a.1",
			tests: []lem.TestCase{
				{
					ID:   "a",
					Name: "1",
				},
			},
		},
		{
			name: "a1./a/1/hello, a2./a/1/world",
			tests: []lem.TestCase{
				{
					ID:   "a1",
					Name: "/a/1/hello",
				},
				{
					ID:   "a2",
					Name: "/a/1/world",
				},
			},
		},
		{
			name: "a1./a/1/hello, a2./a/1/world, a3./a/2/hi",
			tests: []lem.TestCase{
				{
					ID:   "a1",
					Name: "/a/1/hello",
				},
				{
					ID:   "a2",
					Name: "/a/1/world",
				},
				{
					ID:   "a3",
					Name: "/a/2/hi",
				},
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			var tree lem.TestCaseTree
			for _, ltc := range tc.tests {
				path := lem.GetTestCasePath(ltc.ID, ltc.Name)
				tree.Insert(ltc, path...)
			}
			t.Logf("%+v", tree)
		})
	}
}

func TestGetTestCasePath(t *testing.T) {
	testCases := []struct {
		name    string
		argID   string
		argName string
		expPath []string
	}{
		{
			name:    "id",
			argID:   "leak",
			argName: "to sink",
			expPath: []string{"leak", "to sink"},
		},
		{
			name:    "id & name w multiple parts",
			argID:   "leak",
			argName: "to result/malloc",
			expPath: []string{"leak", "to result", "malloc"},
		},
		{
			name:    "id & no ID in path",
			argID:   "leak",
			argName: "/to result",
			expPath: []string{"to result"},
		},
		{
			name:    "id & no ID in path & name w multiple parts",
			argID:   "escape",
			argName: "/no malloc/storing single byte-wide value",
			expPath: []string{"no malloc", "storing single byte-wide value"},
		},
		{
			name:    "id w int",
			argID:   "leak2",
			argName: "to result",
			expPath: []string{"leak", "2", "to result"},
		},
		{
			name:    "id w int & name w multiple parts",
			argID:   "leak2",
			argName: "to result/malloc",
			expPath: []string{"leak", "2", "to result", "malloc"},
		},
		{
			name:    "id w int & no ID in path",
			argID:   "leak2",
			argName: "/to result",
			expPath: []string{"to result"},
		},
		{
			name:    "id w int & no ID in path & name w multiple parts",
			argID:   "escape3",
			argName: "/no malloc/storing single byte-wide value",
			expPath: []string{"no malloc", "storing single byte-wide value"},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			actPath := lem.GetTestCasePath(tc.argID, tc.argName)
			if el, al := len(tc.expPath), len(actPath); el != al {
				t.Errorf("len(expPath)=%d != len(actPath)=%d", el, al)
			} else {
				for j := range tc.expPath {
					if ep, ap := tc.expPath[j], actPath[j]; ep != ap {
						t.Errorf(
							"expPath[%[1]d]=%[2]s != actPath[%[1]d]=%[3]s",
							i, ep, ap,
						)
					}
				}
			}
			if t.Failed() {
				t.Errorf("expPath=%v, actPath=%v", tc.expPath, actPath)
			}
		})
	}
}
