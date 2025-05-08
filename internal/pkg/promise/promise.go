package promise

import (
	"github.com/macabot/hypp/js"
)

// Await awaits a Promise.
// Based on https://stackoverflow.com/a/68427221
func Await(awaitable js.Value) (js.Value, error) {
	then := make(chan js.Value)
	defer close(then)
	thenFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		then <- args[0]
		return nil
	})
	defer thenFunc.Release()

	catch := make(chan js.Value)
	defer close(catch)
	catchFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		catch <- args[0]
		return nil
	})
	defer catchFunc.Release()

	awaitable.Call("then", thenFunc).Call("catch", catchFunc)

	select {
	case result := <-then:
		return result, nil
	case err := <-catch:
		return js.Null(), js.Error{Value: err}
	}
}
