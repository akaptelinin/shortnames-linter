package nodefault

import "fmt"

func withDefaultWhitelisted(i, j, k, n, v int, ok, id, ip, db, tx, mu, wg, rw, fn, cb, ch, t, b interface{}) { // want `parameter name "i" is too short` `parameter name "j" is too short` `parameter name "k" is too short` `parameter name "n" is too short` `parameter name "v" is too short` `parameter name "ok" is too short` `parameter name "id" is too short` `parameter name "ip" is too short` `parameter name "db" is too short` `parameter name "tx" is too short` `parameter name "mu" is too short` `parameter name "wg" is too short` `parameter name "rw" is too short` `parameter name "fn" is too short` `parameter name "cb" is too short` `parameter name "ch" is too short` `parameter name "t" is too short` `parameter name "b" is too short`
	fmt.Println(i, j, k, n, v, ok, id, ip, db, tx, mu, wg, rw, fn, cb, ch, t, b)
}
