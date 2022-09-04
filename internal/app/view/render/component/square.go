package component

import (
	"fmt"

	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/html"
	"github.com/macabot/senet/internal/app/model"
)

type SquareProps struct {
	Position    model.Position
	Selected    bool
	Highlighted bool
}

var iconToClass = map[model.Icon]string{
	model.Foo: "foo",
	model.Bar: "bar",
	model.Pit: "pit",
	model.Buz: "buz",
}

func Square(props SquareProps) *hypp.VNode {
	classes := map[string]bool{
		"square":                                 true,
		"selected":                               props.Selected,
		"highlighted":                            props.Highlighted,
		fmt.Sprintf("row-%d", props.Position[0]): true,
		fmt.Sprintf("column-%d", props.Position[1]): true,
	}
	if icon, ok := model.SpecialPositions[props.Position]; ok {
		classes[iconToClass[icon]] = true
	}
	return html.Button(
		hypp.HProps{
			"class":    classes,
			"disabled": !props.Highlighted,
		},
	)
}
