package main

import (
	"fmt"
	"sort"
)

type node struct {
	isEnd  bool
	vaild  bool
	elems  []bool
	lower  int
	upper  int
	nodeN  int
	pnodeN int
}

type good struct {
	w int
	m int
}

type ByKey []node

// for sorting by key.
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].upper < a[j].upper }

func graphGen(ns2 []node, gs map[int]good, step int) (node, string, string) {
	//var ns2 []node
	//var nM node
	ns := make([]node, len(ns2))
	copy(ns, ns2)
	//fmt.Println(ns)
	sort.Sort(ByKey(ns))
	n := ns[len(ns)-1]

	log := fmt.Sprintf("Select node: %v, in queue:%v\n", n.nodeN, ns)

	if n.isEnd {
		return n, "", log
	}

	n1 := n
	n1.elems = make([]bool, len(n.elems))
	copy(n1.elems, n.elems)
	n1.pnodeN = n.nodeN
	n1.nodeN = step + 1
	n1.elems = append(n1.elems, true)
	n1.lower += gs[len(n1.elems)].m
	n1.upper = n1.lower
	for i, v := range gs {
		if i > len(n1.elems) {
			n1.upper += v.m
		}
	}
	n1.vaild = testVaild(n1, gs)
	n1.isEnd = testIsEnd(n1)

	n2 := n
	n2.elems = make([]bool, len(n.elems))
	copy(n2.elems, n.elems)
	n2.pnodeN = n.nodeN
	n2.nodeN = step + 2
	n2.elems = append(n2.elems, false)
	//n2.lower += gs[len(n1.elem)]
	n2.upper = n2.lower
	for i, v := range gs {
		if i > len(n2.elems) {
			n2.upper += v.m
		}
	}
	n2.vaild = testVaild(n2, gs)
	n2.isEnd = testIsEnd(n2)

	ns = ns[:len(ns)-1]
	if n1.vaild {
		ns = append(ns, n1)
	}
	if n2.vaild {
		ns = append(ns, n2)
	}

	fmt.Println((n1))
	fmt.Println((n2))
	n, g, l := graphGen(ns, gs, step+2)
	return n, fmt.Sprintf("%v%v", print([]node{n1, n2}), g), fmt.Sprintf("%v%v", log, l)

}

func testIsEnd(n node) bool {
	return len(n.elems) == 4
}

func testVaild(n node, gs map[int]good) bool {
	w := 0
	for i, b := range n.elems {
		if b {
			w += gs[i+1].w
		}
	}
	return w <= 16
}

func asmName(n []bool) string {
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
		if len(n.elems) > 0 {
			name := asmName(n.elems)
			re = fmt.Sprintf("%s%s", re, fmt.Sprintf("\"%v:%v\"->\"%v:%v\";\n", n.pnodeN, asmName(n.elems[:len(n.elems)-1]), n.nodeN, name))

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
	gs := map[int]good{1: good{m: 100, w: 10},
		2: good{m: 63, w: 7},
		3: good{m: 56, w: 8},
		4: good{m: 12, w: 4}}
	n, g, l := graphGen([]node{node{}}, gs, 0)

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("\nSteps:\n\n", l)
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("\nGraph:\n\n", g)
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Result:", n)
}
