package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSum2(t *testing.T) {
	tables := []struct {
		a int
		b int
		c int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{3, 3, 6},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.c {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", total, item.c)
		}
	}
}
