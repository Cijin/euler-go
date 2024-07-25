package twentyfour

import "fmt"

func permutateInt8(inputs []int8) []int8 {
	output := make([]int8, len(inputs))
	copy(output, inputs)

	size := len(inputs)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}

	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		tmp := inputs[j]
		inputs[j] = inputs[i]
		inputs[i] = tmp
		output := make([]int8, len(inputs))
		copy(output, inputs)

		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}

	return output
}

func SolveTwentyFour() {
	fmt.Println(permutateInt8([]int8{1, 2}))
}
