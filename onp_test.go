package onp

import (
	"testing"
)

func TestOnp(t *testing.T) {
	cases := []struct {
		a        []rune
		b        []rune
		expected int
	}{
		{
			a:        []rune("abc"),
			b:        []rune("abcdef"),
			expected: 3,
		},
		{
			a:        []rune("abc"),
			b:        []rune("ab"),
			expected: 1,
		},
		{
			a:        []rune("abc"),
			b:        []rune("abc"),
			expected: 0,
		},
	}

	for _, c := range cases {
		if NewDiff(c.a, c.b).ONP() != c.expected {
			t.Errorf("a:%v b:%v expected:%v", c.a, c.b, c.expected)
		}
	}
}
