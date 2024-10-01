package b

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type sumTestCase struct {
	name string
	a    int
	b    int
	r    int
}

func TestSum(t *testing.T) {
	testCases := []sumTestCase{
		{name: "default", a: 1, b: 2, r: 3},
		{name: "zero", a: 0, b: 0, r: 0},
		{name: "minus plus zero", a: -1, b: 1, r: 0},
		{name: "1000-7", a: 1000, b: -7, r: 993},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, sum(tc.a, tc.b), tc.r)
		})
	}
}
