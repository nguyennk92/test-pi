package common

import (
	"math/rand"
)

func Pi(num int) float64 {
	rand.Seed(100)
	count1 := 0
	count2 := 0
	for i := 0; i < num; i++ {
		x := rand.Float64()
		y := rand.Float64()
		d := x*x + y*y
		if d <= 1 {
			count1++
		}
		count2++
	}

	return 4 * (float64(count1) / float64(count2))
}

func Pi2(num int) float64 {
	return Pi(num)
}
