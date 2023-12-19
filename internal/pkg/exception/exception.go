package exception

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/macabot/hypp/window"
)

// TODO update package dependency graph

func Handle() {
	if r := recover(); r != nil {
		window.Console().Error(fmt.Sprint(r))
		for _, part := range strings.Split(string(debug.Stack()), "\n") {
			window.Console().Error(part)
		}
		// window.Console().Error(fmt.Sprintf("%v\n%s", r, string(debug.Stack())))
		panic(r)
	}
}
