package a

import (
	x "context" // want `import alias "x" is too short`
	"fmt"
	y "net/http" // want `import alias "y" is too short`
)

type Handler struct{}

func (h *Handler) Bad(a x.Context, b *y.Request) error { // want `receiver name "h" is too short` `parameter name "a" is too short` `parameter name "b" is too short`
	fmt.Println(h, a, b)
	return nil
}

func (handler *Handler) Good(ctx x.Context, request *y.Request) error {
	fmt.Println(handler, ctx, request)
	return nil
}

func shortReturn() (s string, err error) { // want `named return "s" is too short`
	return "", nil
}

func goodReturn() (result string, err error) {
	return "", nil
}

func whitelisted(i, j, k int, id string, ctx x.Context, err error) (ok bool) {
	fmt.Println(i, j, k, id, ctx, err)
	return true
}
