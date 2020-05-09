package base

import (
	"testing"
)

type pair struct {
	a   int
	b   int
	res int
}

func TestSum(t *testing.T) {
	testData := []pair{
		{a: 1, b: 2, res: 3},
		{a: -1, b: 0, res: -1},
		{a: -1, b: 13, res: 12},
	}

	for _, p := range testData {
		if sumInt(p.a, p.b) != p.res {
			t.Errorf("sum for pair %d,%d expect: %d, but got:%d", p.a, p.b, p.res, sumInt(p.a, p.b))
		}
	}
}

// go test -v -cover -run SumGroup/Group1
func TestSumGroup(t *testing.T) {
	testData := []pair{
		{a: 1, b: 2, res: 3},
		{a: -1, b: 0, res: -1},
		{a: -1, b: 13, res: 12},
	}

	t.Run("Group1", func(t *testing.T) {
		for _, p := range testData {
			if sumInt(p.a, p.b) != p.res {
				t.Errorf("sum for pair %d,%d expect: %d, but got:%d", p.a, p.b, p.res, sumInt(p.a, p.b))
			}
		}
	})

	t.Run("Group2", func(t *testing.T) {
		for _, p := range testData {
			if sumInt(p.a, p.b) != p.res {
				t.Errorf("sum for pair %d,%d expect: %d, but got:%d", p.a, p.b, p.res, sumInt(p.a, p.b))
			}
		}
	})

}
