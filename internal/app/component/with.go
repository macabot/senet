package component

import (
	"github.com/macabot/hypp"
	"golang.org/x/exp/maps"
)

type WithWrapper func(node *hypp.VNode) *hypp.VNode

func patchHProps(node *hypp.VNode, patch hypp.HProps) *hypp.VNode {
	if node == nil {
		return node
	}
	if node.Kind() != hypp.ElementNode {
		return node
	}
	props := node.Props()
	if props == nil {
		props = hypp.HProps{}
	}
	maps.Copy(props, patch)
	return hypp.H(
		node.Tag(),
		props,
		node.Children()...,
	)
}

func With(node *hypp.VNode, wrappers ...WithWrapper) *hypp.VNode {
	for _, wrapp := range wrappers {
		node = wrapp(node)
	}
	return node
}

func Key(key any) WithWrapper {
	return func(node *hypp.VNode) *hypp.VNode {
		return patchHProps(node, hypp.HProps{"key": key})
	}
}
