package storage

import "testing"

type testCase struct {
	name  string
	key   string
	value string
}

func TestSetGet(t *testing.T) {
	cases := []testCase{
		{"hello world", "hello", "world"},
		{"hello world", "hello", "world"},
	}

	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)

			sValue := s.Get(c.key)

			if *sValue != c.value {
				t.Errorf("values not equal")
			}
		})
	}
}

type testCaseWithKind struct {
	name  string
	key   string
	value string
	kind  Kind
}

func TestSetGetWithType(t *testing.T) {
	cases := []testCaseWithKind{
		{"hello world", "hello", "world", KindString},
		{"int value", "key", "666667778", KindInt},
	}

	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)

			sValue := s.Get(c.key)

			if *sValue != c.value {
				t.Errorf("values not equal")
			}

			if getType(*sValue) != getType(c.value) {
				t.Errorf("value kinds not equal")
			}

			if getType(*sValue) != c.kind {
				t.Errorf("expected value kind: %v", c.kind)
			}
		})
	}
}
