package a

import "fmt"

func blankParams(_, _ int, name string) {
	fmt.Println(name)
}

func blankMixed(a int, _ string, _ bool) { // want `parameter name "a" is too short`
	fmt.Println(a)
}
