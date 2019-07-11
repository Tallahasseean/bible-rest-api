package main

import (
	"bible/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")

}
