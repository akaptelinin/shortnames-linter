package combined

import "fmt"

// Test: -disable-default-whitelist=true -whitelist="i,ok"
// i and ok should be allowed (custom whitelist)
// j, k, id, db should trigger warnings (default whitelist disabled)

func testCombined(i, ok, j, k, id, db interface{}) { // want `parameter name "j" is too short` `parameter name "k" is too short` `parameter name "id" is too short` `parameter name "db" is too short`
	fmt.Println(i, ok, j, k, id, db)
}
