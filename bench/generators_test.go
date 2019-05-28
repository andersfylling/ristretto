/*
 * Copyright 2019 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bench

import (
	"testing"
)

func GetSame(benchmark *Benchmark) func(b *testing.B) {
	return func(b *testing.B) {
		b.Run("single", func(b *testing.B) {
			cache := benchmark.Create()
			cache.Set("*", []byte("*"))
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				cache.Get("*")
			}
		})
		b.Run("multiple", func(b *testing.B) {
			cache := benchmark.Create()
			cache.Set("*", []byte("*"))
			b.SetParallelism(benchmark.Para)
			b.SetBytes(1)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					cache.Get("*")
				}
			})
		})
	}
}

func GetZipf(benchmark *Benchmark) func(b *testing.B) {
	return func(b *testing.B) {
		b.Run("single", func(b *testing.B) {
			// setup
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				// do something
			}
		})
		b.Run("multiple", func(b *testing.B) {
			// setup
			b.SetParallelism(benchmark.Para)
			b.SetBytes(1)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					// do something
				}
			})
		})
	}
}

func SetSame(name string, para int, create func() Cache) func(b *testing.B) {
	return func(b *testing.B) {
		b.Run("single", func(b *testing.B) {
			// setup
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				// do something
			}
		})
		b.Run("multiple", func(b *testing.B) {
			// setup
			b.SetParallelism(para)
			b.SetBytes(1)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					// do something
				}
			})
		})
	}
}

func SetZipf(name string, para int, create func() Cache) func(b *testing.B) {
	return func(b *testing.B) {
		b.Run("single", func(b *testing.B) {
			// setup
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				// do something
			}
		})
		b.Run("multiple", func(b *testing.B) {
			// setup
			b.SetParallelism(para)
			b.SetBytes(1)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					// do something
				}
			})
		})
	}
}

func GetSet(name string, para int, create func() Cache) func(b *testing.B) {
	return func(b *testing.B) {
		b.Run("single", func(b *testing.B) {
			// setup
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				// do something
			}
		})
		b.Run("multiple", func(b *testing.B) {
			// setup
			b.SetParallelism(para)
			b.SetBytes(1)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					// do something
				}
			})
		})
	}
}