package medium

import "testing"

func TestKthSmallest(t *testing.T) {
	type Item struct {
		matrix [][]int
		k      int
		except int
	}
	cases := []Item{
		{
			matrix: [][]int{
				[]int{1, 5, 9},
				[]int{10, 11, 13},
				[]int{12, 13, 15},
			},
			k:      8,
			except: 13,
		},
	}
	for _, item := range cases {
		out := KthSmallest(item.matrix, item.k)
		if out != item.except {
			t.Errorf("matrix: %v, k: %d, except: %d, got: %d", item.matrix, item.k, item.except, out)
		}
	}
}
