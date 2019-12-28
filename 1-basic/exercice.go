package main
import "fmt"

func main() {
	//Exercice 1
	fmt.Printf("Exercice 1:\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Printf("\n")

	//Exercice 2
	fmt.Printf("Exercice 2:\n")
	j:=0
	here: 
	fmt.Printf("%d", j)
		j++
		if j == 10 {
			goto here2
		}
		goto here
	here2:
	fmt.Printf("\n")
	//Exercice 3:
	fmt.Printf("Exercice 3:\n")
	var myarray [10]int

	for i:=0; i < 10; i++ {
		myarray[i] = i
	}
	fmt.Printf("%v", myarray)
	fmt.Printf("\n")
}