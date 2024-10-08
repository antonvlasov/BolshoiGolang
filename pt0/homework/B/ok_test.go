package b

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type okTest struct {
	name string
	s    []int
	r    bool
}

func TestYes(t *testing.T) {
	cases := []okTest{
		{name: "golang", s: []int{1, 1, 1, 1, 1}, r: false},
		{name: "empty", s: []int{}, r: false},
		{name: "yes", s: []int{0}, r: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, isOK(c.s), c.r)
		})
	}
}
