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

package benchmarks_test

import "testing"

var (
	_leakDst *int32
	_moveDst *int32
)

func leak2Global(leaked2global *int32) *int32 {
	_leakDst = leaked2global
	return leaked2global
}

func leak2Result(leaked2result *int32) *int32 {
	return leaked2result
}

func move2global() {
	var moved2heap int32 = 4096
	_moveDst = &moved2heap
}

func noop(args ...interface{}) {
}

//go:noinline
func BenchmarkLem(b *testing.B) {

	b.Run("Escape", func(b *testing.B) {
		b.Run("ByChan", func(b *testing.B) {
			c := make(chan *int32, 1)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				x := new(int32) // Escapes
				c <- x
				<-c
			}
		})
		b.Run("ByIface", func(b *testing.B) {
			b.Run("WithNil", func(b *testing.B) {
				var dst interface{}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var p *int32 // NA
					dst = p
				}
				_ = dst
			})
			b.Run("WithPointer", func(b *testing.B) {
				var dst interface{}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var p = new(int32) // Escapes
					dst = p
				}
				_ = dst
			})
			b.Run("WithValue", func(b *testing.B) {
				var dst interface{}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var x int32 = 4096 // NA
					dst = x
				}
				_ = dst
			})
		})
		b.Run("BySlice", func(b *testing.B) {
			b.Run("WithNil", func(b *testing.B) {
				dst := make([]*int32, 1, 1)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var p *int32 // NA
					dst[0] = p
				}
				_ = dst
			})
			b.Run("WithPointer", func(b *testing.B) {
				dst := make([]*int32, 1, 1)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var p = new(int32) // Escapes
					dst[0] = p
				}
				_ = dst
			})
			b.Run("WithValue", func(b *testing.B) {
				dst := make([]int32, 1, 1)
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var x int32 = 4096 // NA
					dst[0] = x
				}
				_ = dst
			})
		})
	})
	b.Run("Leak", func(b *testing.B) {
		b.Run("ToGlobal", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				p := new(int32) // Escapes
				*p = 4096
				leak2Global(p)
			}
		})
		b.Run("ToResult", func(b *testing.B) {
			b.Run("DoNotStore", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					p := new(int32) // Does not escape
					*p = 4096
					leak2Result(p)
				}
			})
			b.Run("Store", func(b *testing.B) {
				b.Run("InLoop", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						var pp *int32
						p := new(int32) // Does not escape
						*p = 4096
						pp = leak2Result(p)
						_ = pp
					}
				})
				b.Run("OutLoop", func(b *testing.B) {
					var pp *int32
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						p := new(int32) // Escapes
						*p = 4096
						pp = leak2Result(p)
					}
					b.StopTimer()
					_ = pp
				})
			})
		})
	})
	b.Run("Move", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			move2global()
		}
	})
	b.Run("NoEscape", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var (
				a = new(int32)
				b = new(int64)
				c = make([]int32, 5, 5)
				d = make([]*int64, 10, 10)
				s = struct {
					a int32
					b int32
				}{a: 4096, b: 4096}
				ps = &struct {
					a int32
					b int32
				}{a: 4096, b: 4096}
			)
			noop(a, b, c, d, s, ps)
		}
	})
}
