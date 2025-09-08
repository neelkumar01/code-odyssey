package slice

func Example(a [3]int, b []int) []int {
	a[0] = 24		// x assigned to a, a change but x don't
	b[0] = 24		
	c := make([]int, 5)			// slice for array = [0, 0, 0, 0, 0]
	c[3] = 50

	copy(c , b)
	return c
}
