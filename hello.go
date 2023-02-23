package main

import "fmt"

func add(x int, y int) int {
	fmt.Printf("x=%d,y=%d\n", x, y)
	return x + y;
}

func main() {
	// var (
  //   a int = 15
  //   b bool = false
  //   c string = "6xxx5"
  //   d [5]int
	// )

	m := map[string]map[string]int{
		"a": map[string]int{
			"a1": 1,
			"a2": 2,
		},
		"b": map[string]int{
			"b1": 1,
			"b2": 2,
		},
		"c": map[string]int{
			"c1": 1,
			"c2": 2,
		},
		"d": map[string]int{
			"d1": 1,
			"d2": 2,
		},
	}

	// fmt.Println(m["a"])

	// fmt.Println(add(11, 22));

	// if name, ok := m["a"]; ok {
	// 	fmt.Println(name, ok)
	// }


	for k, v := range m {
		fmt.Printf("key=%s\n", k)
		sum := 0;
		for k1, item := range v {			
			sum = add(sum, item)
			fmt.Printf("sub\nkey=%s,value=%d\n", k1, item)
		}
		fmt.Printf("total=%d\n", sum)
	}
}
