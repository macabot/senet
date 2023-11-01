package main

import (
	"errors"

	_ "github.com/macabot/hypp/jsd"
	"github.com/macabot/hypp/window"
	"github.com/macabot/senet/internal/app"
)

func main() {
	el := window.Document().QuerySelector("html")
	if el.IsNull() {
		panic(errors.New("Could not find <html> element."))
	}

	app.Run(
		el,
	)
	select {} // run Go forever
}
