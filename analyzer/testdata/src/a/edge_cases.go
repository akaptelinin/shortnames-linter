package a

import "fmt"

type Server struct{}
type Client struct{}

func noParams() {}

func noReceiver(name string) {}

func anonymousParams(string, int) {}

func mixedParams(a int, name string, b int) { // want `parameter name "a" is too short` `parameter name "b" is too short`
	fmt.Println(a, name, b)
}

func exactlyThreeChars(abc string, def int) {
	fmt.Println(abc, def)
}

func twoCharsNotWhitelisted(ab string) { // want `parameter name "ab" is too short`
	fmt.Println(ab)
}

func whitelistedNames(i, j, k int, ctx, err, ok, id, db, tx, mu, wg, ch, fn, cb, rw, ip, v, n interface{}) {
	fmt.Println(i, j, k, ctx, err, ok, id, db, tx, mu, wg, ch, fn, cb, rw, ip, v, n)
}

func (s *Server) pointerReceiver() { // want `receiver name "s" is too short`
	fmt.Println(s)
}

func (c Client) valueReceiver() { // want `receiver name "c" is too short`
	fmt.Println(c)
}

func (server *Server) goodPointerReceiver() {
	fmt.Println(server)
}

func (client Client) goodValueReceiver() {
	fmt.Println(client)
}

func variadicShort(a ...int) { // want `parameter name "a" is too short`
	fmt.Println(a)
}

func variadicOk(args ...int) {
	fmt.Println(args)
}

func multiReturn() (a, b int) { // want `named return "a" is too short` `named return "b" is too short`
	return 1, 2
}

func multiReturnMixed() (result int, e error) { // want `named return "e" is too short`
	return 0, nil
}

func multiReturnOk() (result int, err error) {
	return 0, nil
}

func unnamedReturn() (int, error) {
	return 0, nil
}

func initFunc() {
	fmt.Println("init")
}

var closure = func(x int) {
	fmt.Println(x)
}

var goodClosure = func(value int) {
	fmt.Println(value)
}

type Processor interface {
	Handle(r Request)
}

type Request struct{}

type GoodProcessor interface {
	Handle(request Request)
}

func genericFunc[T any](t T) { // want `parameter name "t" is too short`
	fmt.Println(t)
}

func goodGenericFunc[T any](value T) {
	fmt.Println(value)
}

type GenericType[T any] struct {
	value T
}

func (g *GenericType[T]) Method(t T) { // want `receiver name "g" is too short` `parameter name "t" is too short`
	fmt.Println(g.value, t)
}

func (gtype *GenericType[T]) GoodMethod(item T) {
	fmt.Println(gtype.value, item)
}
