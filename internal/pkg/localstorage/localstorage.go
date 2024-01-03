package localstorage

import (
	"github.com/macabot/hypp/js"
)

// SetItem is based on https://developer.mozilla.org/en-US/docs/Web/API/Storage/setItem
func SetItem(keyName, keyValue string) {
	js.Global().Get("localStorage").Call("setItem", keyName, keyValue)
}

// GetItem is based on https://developer.mozilla.org/en-US/docs/Web/API/Storage/getItem
func GetItem(keyName string) (string, bool) {
	v := js.Global().Get("localStorage").Call("getItem", keyName)
	if v.IsNull() {
		return "", false
	}
	return v.String(), true
}

// RemoveItem is based on https://developer.mozilla.org/en-US/docs/Web/API/Storage/removeItem
func RemoveItem(keyName string) {
	js.Global().Get("localStorage").Call("removeItem", keyName)
}
