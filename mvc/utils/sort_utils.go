package utils

import "sort"

// []int{9,8,7,6,5}
// []int{5,6,7,8,9}
func BubbleSort(elements []int) []int {
	keepRunning := true

	for keepRunning {
		keepRunning = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
	return elements
}

func Sort(els []int) {
	if len(els) < 1000 {
		BubbleSort(els)
	}

	sort.Ints(els)
}
