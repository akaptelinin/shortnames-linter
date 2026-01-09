package whitelist_edge

import "fmt"

// Test whitelist parsing edge cases:
// - spaces around values: " ab , xy "
// - empty values from double comma: "ab,,xy"
// - trailing comma: "ab,xy,"
// All should work: ab and xy whitelisted, zz should warn

func testWhitelistEdge(ab, xy, zz interface{}) { // want `parameter name "zz" is too short`
	fmt.Println(ab, xy, zz)
}
