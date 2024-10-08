package mathy

import "testing"

type test struct {
	name     string
	arg1     int
	arg2     int
	expected int
}

func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAddTable(t *testing.T) {
	addTests := []test{
		{"one", 2, 3, 5},
		{"two", 4, 8, 12},
		{"three", 6, 9, 15},
		{"four", 3, 10, 13},
		{"five", -10, 10, 0},
	}

	for _, test := range addTests {
		t.Run(test.name, func(t *testing.T) {
			if output := Add(test.arg1, test.arg2); output != test.expected {
				t.Errorf("Output %q not equal to expected %q", output, test.expected)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}
