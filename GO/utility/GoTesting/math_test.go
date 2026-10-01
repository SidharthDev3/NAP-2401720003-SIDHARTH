package gotesting

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, c := range cases {
		if c.a+c.b != c.want {
			t.Errorf("Add %v and %v = %v , wanted %v", c.a, c.b, c.a+c.b, c.want)
		}
	}
}
