package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkDraw(b *testing.B) {
	base := "/Users/crebert/cursus42/side_project/src/picrew/assets/"
	if err := loadAllAssets(base); err != nil {
		panic(err)
	}
	for n := 0; n < b.N; n++ {
		if err := generate(base, fmt.Sprintf("../out/%d.png", n), 0, 0, 0, 0, 0); err != nil {
			panic(err)
		}
	}
}

func BenchmarkDrawParallel(b *testing.B) {
	base := "/Users/crebert/cursus42/side_project/src/picrew/assets/"
	if err := loadAllAssets(base); err != nil {
		panic(err)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := generate(base, fmt.Sprintf("../out/%d.png", rand.Int()), 0, 0, 0, 0, 0); err != nil {
				panic(err)
			}
		}
	})
}
