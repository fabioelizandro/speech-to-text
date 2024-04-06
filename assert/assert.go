package assert

import "fmt"

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func True(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}

func False(condition bool, msg string) {
	True(!condition, msg)
}

func Unreachable(msg string) {
	panic(fmt.Sprintf("Unreachable code executed: %s", msg))
}

func Must[T any](value T, err error) T {
	NoErr(err)
	return value
}
