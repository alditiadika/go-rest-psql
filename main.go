package main

import (
	"fmt"

	"github.com/alditiadika/go-rest-psql/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	fmt.Println("HTTP REST run at port 3000")
	app.Run(":3000")
}
