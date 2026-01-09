package whitelist

import "fmt"

func customWhitelist(ab string, xy int, zz string) { // want `parameter name "zz" is too short`
	fmt.Println(ab, xy, zz)
}

func stillDefaultWhitelist(i, j, k int, ctx, err interface{}) {
	fmt.Println(i, j, k, ctx, err)
}
