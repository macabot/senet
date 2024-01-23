// sessionstorage corresponds to the sessionStorage property.
// See https://developer.mozilla.org/en-US/docs/Web/API/Window/sessionStorage
package sessionstorage

import "github.com/macabot/hypp/js"

// SetItem sets an item in the sessionStorage.
// See https://developer.mozilla.org/en-US/docs/Web/API/Storage/setItem
func SetItem(keyName, keyValue string) {
	js.Global().Get("sessionStorage").Call("setItem", keyName, keyValue)
}

// GetItem gets an item from the sessionStorage.
// See https://developer.mozilla.org/en-US/docs/Web/API/Storage/getItem
func GetItem(keyName string) *string {
	v := js.Global().Get("sessionStorage").Call("getItem", keyName)
	if v.IsNull() {
		return nil
	}
	s := v.String()
	return &s
}
