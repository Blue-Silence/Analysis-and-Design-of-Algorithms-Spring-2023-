package main

import "fmt"

func graphGen(current2 []string, links map[string]([]string)) [][]string {
	var current []string = make([]string, len(current2))
	//fmt.Println(current2)
	copy(current, current2)
	//fmt.Println(current2)
	re := [][]string{current}

	tail := current[len(current)-1]
	for i := 0; i < len(current)-1; i++ {
		if current[i] == tail {
			//fmt.Println("current get:", current)
			return re
		}
	}
	//fmt.Println("111")
	var re2 [][]string = make([][]string, 0)
	for _, v := range links[tail] {
		//fmt.Println("v:", v, "  before111:", re2)
		r := graphGen(append(current, v), links)

		re2 = append(re2, r...)
		//fmt.Println("after:", re2)
	}
	//fmt.Println("current:", current, "  re2:", re2)
	return append(re, re2...)
}

func asmName(n []string) string {
	if len(n) == 1 {
		return fmt.Sprintf("%v", n[0])
	} else {
		s := fmt.Sprintf("%v,", n[0])
		return fmt.Sprintf("%s%s", s, asmName(n[1:]))
	}
}

func print(ns [][]string) string {
	var re string
	for _, n := range ns {
		if len(n) > 1 {
			name := asmName(n)
			re = fmt.Sprintf("%s%s", re, fmt.Sprintf("\"%v\"->\"%v\";\n", asmName(n[:len(n)-1]), name))

			tail := n[len(n)-1]
			c := 0
			for _, v := range n {
				if v == tail {
					c++
				}
			}
			switch {
			case tail == "a" && len(n) == 8:
				re = fmt.Sprintf("%s%s", re, fmt.Sprintf("\"%v\" [style=filled,color=green];\n", name))
			case c > 1:
				re = fmt.Sprintf("%s%s", re, fmt.Sprintf("\"%v\" [style=filled,color=red];\n", name))
			default:

			}
		}
	}
	return re
}

func main() {
	//fmt.Println("233")
	links := make(map[string]([]string))
	links["a"] = []string{"c", "f", "d", "b"}
	links["c"] = []string{"a", "f"}
	links["f"] = []string{"c", "a", "d", "g"}
	links["d"] = []string{"a", "b", "f", "g"}
	links["b"] = []string{"a", "d", "g", "e"}
	links["g"] = []string{"f", "d", "b", "e"}
	links["e"] = []string{"b", "g"}
	//fmt.Println(graphGen([]string{"a"}, links))
	//graphGen([]string{"a"}, links)
	fmt.Println(print(graphGen([]string{"a"}, links)))
}
