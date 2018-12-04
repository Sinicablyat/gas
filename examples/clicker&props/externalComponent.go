package main

import (
	"fmt"
	"github.com/Sinicablyat/gas"
)

// GetNumberViewer return very cool number viewer.
// It can be in another directory too.
// For reference from not parent component you can use `values` (they will reload).
func GetNumberViewer(this *gas.Component, values ...interface{}) interface{} {
	return gas.NewComponent(
		this,
		gas.NilData,
		gas.NilMethods,
		gas.NilDirectives,
		gas.NilBinds,
		gas.NilHandlers,
		"i",
		map[string]string{
			"id": "needful_wrapper--number-viewer",
		},
		func(this3 *gas.Component) interface{} {
			return fmt.Sprintf("%d times", values[0])
		})
}