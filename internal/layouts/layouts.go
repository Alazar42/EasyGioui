package layouts

import "gioui.org/layout"

func Vertical(children ...layout.FlexChild) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, children...)
	}
}

func Horizontal(children ...layout.FlexChild) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, children...)
	}
}
