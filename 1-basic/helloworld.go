package main
import "fmt"// Implements formatted I/O.

/* Print something */
func main() {
	fmt.Printf("Hello, world.\n")

	s := "hello"
	c := []rune(s)
	c[0] = 'c'	 
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	J:  for j := 0; j < 5; j++ {
        for i := 0; i < 10; i++ {
            if i > 5 {
                break J
            }
            fmt.Println(i)
        }
    }

	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		// do something with k and v
		fmt.Printf("%d: %s\n", k, v)
	}


	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0])
}