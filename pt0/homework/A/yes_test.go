package a

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type yesCase struct {
	name string
	x    string
	y    string
}

func TestYes(t *testing.T) {
	cases := []yesCase{
		{name: "golang", x: "golang", y: "golang - YES"},
		{name: "yeei", x: "yeei", y: "yeei - YES"},
		{name: "empty", x: "", y: " - YES"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, yes(c.x), c.y)
		})
	}
}
