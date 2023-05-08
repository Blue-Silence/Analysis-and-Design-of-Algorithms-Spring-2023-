package main

import (
	"fmt"
)

type node struct {
	isEnd  bool
	vaild  bool
	path   []string
	len    int
	nodeN  int
	pnodeN int
}

type link struct {
	w   int
	dst string
}

func graphGen(ns2 []node, ls map[string][]link, step int, best node) (node, string, string) {
	//var ns2 []node
	//var nM node
	ns := make([]node, len(ns2))
	copy(ns, ns2)
	fmt.Println(ns)
	if len(ns) == 0 {
		return best, "", ""
	}
	n := ns[0]
	ns = ns[1:]
	log := fmt.Sprintf("Select node: %v, in queue:%v\n", n.nodeN, ns)

	if n.isEnd {
		if n.len < best.len {
			b, g, l := graphGen(ns, ls, step, n)
			return b, g, fmt.Sprintf("%s%s", log, l)
		} else {
			b, g, l := graphGen(ns, ls, step, best)
			return b, g, fmt.Sprintf("%s%s", log, l)
		}

		//return n, "", log
	}
	if n.len > best.len {
		n.vaild = false
		name := asmName(n.path)
		b, g, l := graphGen(ns, ls, step, best)
		return b, fmt.Sprintf("%s%s", fmt.Sprintf("\"%v:%v\" [style=filled,color=red];\n", n.nodeN, name), g), fmt.Sprintf("%s%s", log, l)
	}

	var nsN []node
	var nsNA []node

	for _, l := range ls[n.path[len(n.path)-1]] {
		n1 := n
		n1.path = make([]string, len(n.path))
		copy(n1.path, n.path)
		n1.pnodeN = n.nodeN
		n1.nodeN = step + 1
		step++
		n1.path = append(n1.path, l.dst)
		n1.len += l.w
		n1.vaild = testVaild(n1, best)
		n1.isEnd = testIsEnd(n1)
		/*if n1.isEnd {
			n1.vaild = true
		}*/
		if n1.vaild {
			nsN = append(nsN, n1)
		}

		nsNA = append(nsNA, n1)

	}
	//fmt.Println("nsN:", nsN, "  n:", n, "  nsN2:", nsN2)
	n, g, l := graphGen(append(ns, nsN...), ls, step, best)
	return n, fmt.Sprintf("%v%v", print(nsNA), g), fmt.Sprintf("%v%v", log, l)

}

func testIsEnd(n node) bool {
	return len(n.path) == 5 && n.path[len(n.path)-1] == "a"
}

func testVaild(n node, best node) bool {
	w := 0
	tail := n.path[len(n.path)-1]
	for _, s := range n.path {
		if s == tail {
			w++
		}
	}
	return (w < 2 || testIsEnd(n)) && n.len < best.len
}

func asmName(n []string) string {
	switch len(n) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%v", n[0])
	default:
		s := fmt.Sprintf("%v,", n[0])
		return fmt.Sprintf("%s%s", s, asmName(n[1:]))
	}
}

func print(ns []node) string {
	var re string
	for _, n := range ns {
		if len(n.path) > 0 {
			name := asmName(n.path)
			re = fmt.Sprintf("%s%s", re, fmt.Sprintf("\"%v:%v\"->\"%v:%v\";\n", n.pnodeN, asmName(n.path[:len(n.path)-1]), n.nodeN, name))

			switch {
			case !n.vaild:
				re = fmt.Sprintf("%s%s", re, fmt.Sprintf("\"%v:%v\" [style=filled,color=red];\n", n.nodeN, name))
			case n.isEnd:
				re = fmt.Sprintf("%s%s", re, fmt.Sprintf("\"%v:%v\" [style=filled,color=green];\n", n.nodeN, name))
			default:

			}
		}
	}
	return re
}

func main() {
	ls := make(map[string][]link)
	ls["a"] = []link{link{5, "c"}, link{7, "d"}, link{2, "b"}}
	ls["c"] = []link{link{5, "a"}, link{1, "d"}, link{8, "b"}}
	ls["d"] = []link{link{1, "c"}, link{7, "a"}, link{3, "b"}}
	ls["b"] = []link{link{2, "a"}, link{3, "d"}, link{8, "c"}}
	n, g, l := graphGen([]node{node{path: []string{"a"}}}, ls, 0, node{len: 999999})

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("\nSteps:\n\n", l)
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("\nGraph:\n\n", g)
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Result:", n)
}
