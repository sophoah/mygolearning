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


	fmt.Printf("map :\n")

	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}

	for month, days := range monthdays {
		fmt.Printf("%s has %d days\n", month, days)
	}
	
	year := 0
	for _, days := range monthdays {
		year += days
	}
	fmt.Printf("Numbers of days in a year: %d\n", year)

	fmt.Printf("February has %d days\n", monthdays["Feb"])


	fmt.Printf("Comma ok form test:\n")
	value, present := monthdays["Jan"]
	value2, present2 := monthdays["Und"]

	fmt.Printf("Value = %d, present = %t\n", value, present)
	fmt.Printf("Value2 = %d, present2 = %t\n", value2, present2)
}