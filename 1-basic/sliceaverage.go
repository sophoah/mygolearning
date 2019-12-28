package main
import "fmt"

func averagefuncanswer() {
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
}

func main() {
	var array [100]float64

	for i := 0; i < 100; i++ {
		array[i] = 5.57 + float64(i) + float64(i) * 5.55
	}
	slice := array [0:99]
	total := 0.0
	for i := 0; i < 99; i++ {
		total += slice[i]
	}
	fmt.Printf("The average is : %f\n", total / 99)
}