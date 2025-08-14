package main

type asdf struct {
}

func main() {
	var a int
	switch _ := a.(type) {
	case []T:
		if true {
		}
		fallthrough
	case [][]T:
	case [][][]T:
	case [][][][]T:
	case [][][][][]T:
	case [][][][][][]T:
	case [][][][][][][]T:
	}
}
