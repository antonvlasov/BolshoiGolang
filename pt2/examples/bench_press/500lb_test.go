package benchpress

import (
	"math/rand"
	"testing"
)

func BenchmarkSliceFunc(b *testing.B) {
	sliceLen := 10
	list := make([]int, 0)
	for i := 0; i < sliceLen; i++ {
		list = append(list, rand.Intn(sliceLen))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SliceFunc(list)
	}
}

type bench struct {
	name     string
	sliceLen int
}

var cases = []bench{
	// {"5", 5},
	// {"10", 10},
	// {"100", 100},
	// {"500", 500},
	{"1000", 10000},
}

func BenchmarkFindSumCases(b *testing.B) {
	for _, tCase := range cases {
		b.Run(tCase.name, func(bb *testing.B) {
			list := make([]int, 0)
			for i := 0; i < tCase.sliceLen; i++ {
				list = append(list, rand.Intn(tCase.sliceLen))
			}

			bb.ResetTimer()

			for i := 0; i < bb.N; i++ {
				FindSum(list)
			}
		})
	}
}

func BenchmarkFindSumSlowCases(b *testing.B) {
	for _, tCase := range cases {
		b.Run(tCase.name, func(bb *testing.B) {
			list := make([]*int, 0)
			for i := 0; i < tCase.sliceLen; i++ {
				x := rand.Intn(tCase.sliceLen)
				list = append(list, &x)
			}

			bb.ResetTimer()

			for i := 0; i < bb.N; i++ {
				FindSumSlow(list)
			}
		})
	}
}
