package main

import (
	"fmt"
	"math/rand"
)

func main() {
	q := rand.Perm(34)

	fmt.Println(q)

	quickSort(q, 0, len(q)-1)

	fmt.Println(q)

}

func quickSort(q []int, low int, heigh int) {

	if heigh < low {
		return
	}

	j := partition(q, low, heigh)
	quickSort(q, low, j-1)
	quickSort(q, j+1, heigh)
}

func partition(q []int, low int, height int) int {
	i, j := low+1, height
	var res [][]int
	fmt.Printf("%vstirng", res)
	for {
		for q[i] < q[low] {
			i++
			if i >= height {
				break
			}
		}

		for q[j] > q[low] {
			j--
			if j <= low {
				break
			}
		}

		if i >= j {
			break
		}

		q[i], q[j] = q[j], q[i]
	}

	q[low], q[j] = q[j], q[low]

	return j
}
