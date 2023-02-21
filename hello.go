package main

import "fmt"

func main() {
	// var (
  //   a int = 15
  //   b bool = false
  //   c string = "6xxx5"
  //   d [5]int
	// )

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	
	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}
}
