package main

import "fmt"

func main() {
	// fmt.Println("regular map")

	// m := map[int]int{1: 2, 2: 3}

	// for v, k := range m {
	// 	fmt.Printf("%d:%d \n", v, k)
	// }

	//generic map
	fmt.Println("Generic map")
	nm := map[any]any{1: "lex", "lex": 2, "3": "lex"}
	for p, k := range nm {
		fmt.Println(p, k)
	}
}
