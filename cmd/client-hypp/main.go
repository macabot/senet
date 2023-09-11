package main

import (
	"errors"
	"syscall/js"

	jsd "github.com/macabot/hypp/driver/js"
	"github.com/macabot/senet/internal/app"
)

func main() {
	el := js.Global().Get("document").Call("querySelector", "html")
	if el.IsNull() {
		panic(errors.New("Could not find <html> element."))
	}

	app.Run(
		jsd.Driver{},
		jsd.Node(el),
	)
	select {} // run Go forever
}
