// Code generated by gintool.

package main

import (
	"code.byted.org/gin/ginex"
	// "code.byted.org/gin/ginext/shutdown"
)

func main() {
	ginex.Init()

	r := ginex.Default()

	// waitFunc, mw := shutdown.New(10e9, 0)
	// r.Use(mw)
	// defer waitFunc()

	register(r)

	r.Run()
}
