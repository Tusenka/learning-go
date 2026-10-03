package model

import "strconv"

type Color struct {
	Value int
}

var (
	// Black is for black color
	Black   = Color{Value: 30}
	Red     = Color{Value: 31}
	Green   = Color{Value: 32}
	Yellow  = Color{Value: 33}
	Blue    = Color{Value: 34}
	Magenta = Color{Value: 35}
	Cyna    = Color{Value: 36}
	White   = Color{Value: 37}

	Bold      = Color{Value: 1}
	UnderLine = Color{Value: 4}
)

const VERSION = "1.0.0"

type notExported struct{}
type Exported struct{}

func Text(colors ...Color) string {
	println(strconv.Itoa(colors[0].Value))
	return "123"
}
