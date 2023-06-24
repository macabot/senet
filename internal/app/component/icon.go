package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/svg"
)

// ProtectedIcon is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Ashield%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func ProtectedIcon() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"width":   "80%",
			"height":  "auto",
			"viewBox": "0 96 960 960",
			"class":   "icon protected",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480 975q-140-35-230-162.5T160 533V295l320-120 320 120v238q0 152-90 279.5T480 975Zm0-62q115-38 187.5-143.5T740 533V337l-260-98-260 98v196q0 131 72.5 236.5T480 913Zm0-337Z",
			},
		),
	)
}

// BlockingIcon is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Afront_hand%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func BlockingIcon() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"width":   "80%",
			"height":  "auto",
			"viewBox": "0 96 960 960",
			"class":   "icon blocking",
		},
		svg.Path(
			hypp.HProps{
				"d": "M500 1056q-137 0-233.5-96.5T170 726V366q0-38 26.5-64t63.5-26q7 0 15 1.5t15 3.5v-35q0-38 26.5-64t63.5-26q8 0 16.5 1.5T413 162q11-29 34.5-47.5T500 96q35 0 62.5 25.5T590 182q5-2 13-4t17-2q38 0 64 26t26 64v164q8-2 16.5-3t13.5-1q38 0 64 26t26 64v210q0 137-96.5 233.5T500 1056Zm0-60q56 0 105.5-21t86-57.5q36.5-36.5 57.5-86T770 726V516q0-13-8.5-21.5T740 486q-12 0-21 8.5t-9 21.5v170l-13 2q-45 7-74.5 40T590 806h-60q3-56 35.5-100t84.5-66V266q0-13-8.5-21.5T620 236q-12 0-21 8.5t-9 21.5v260h-60V186q0-13-8.5-21.5T500 156q-12 0-21 8.5t-9 21.5v340h-60V246q0-13-8.5-21.5T380 216q-12 0-21 8.5t-9 21.5v320h-60V366q0-13-8.5-21.5T260 336q-12 0-21 8.5t-9 21.5v360q0 56 21 105.5t57.5 86q36.5 36.5 86 57.5T500 996Z",
			},
		),
	)
}

// ReturnToStartIcon is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Amove%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func ReturnToStartIcon() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"width":   "80%",
			"height":  "auto",
			"viewBox": "0 -960 960 960",
			"class":   "icon return-to-start",
		},
		svg.Path(
			hypp.HProps{
				"d": "M440-280q-7 0-12-4t-7-10q-14-42-34-70t-40-54q-20-26-33.5-54T300-540q0-58 41-99t99-41q58 0 99 41t41 99q0 40-13.5 68T533-418q-20 26-40 54t-34 70q-2 6-7 10t-12 4Zm0-112q9-14 18-26t17-23q23-30 34-50t11-49q0-33-23.5-56.5T440-620q-33 0-56.5 23.5T360-540q0 29 11 49t34 50q8 11 17 23t18 26Zm0-98q21 0 35.5-14.5T490-540q0-21-14.5-35.5T440-590q-21 0-35.5 14.5T390-540q0 21 14.5 35.5T440-490Zm3 374q-152 0-258-106T79-480q0-76 28.5-142t78-115.5Q235-787 301-815.5T443-844q76 0 142 28.5t115.5 78Q750-688 778.5-622T807-480v18l70-70 42 42-142 142-142-142 42-42 70 70v-18q0-125-89.5-214.5T443-784q-125 0-214.5 89.5T139-480q0 125 89.5 214.5T443-176q57 0 110.5-21.5T647-256l43 43q-48 45-113 71t-134 26Zm-3-424Z",
			},
		),
	)
}

// StartIcon is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Alocation_on%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func StartIcon() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"width":   "80%",
			"height":  "auto",
			"viewBox": "0 -960 960 960",
			"class":   "icon start",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.089-490Q509-490 529.5-510.589q20.5-20.588 20.5-49.5Q550-589 529.411-609.5q-20.588-20.5-49.5-20.5Q451-630 430.5-609.411q-20.5 20.588-20.5 49.5Q410-531 430.589-510.5q20.588 20.5 49.5 20.5ZM480-159q133-121 196.5-219.5T740-552q0-117.79-75.292-192.895Q589.417-820 480-820t-184.708 75.105Q220-669.79 220-552q0 75 65 173.5T480-159Zm0 79Q319-217 239.5-334.5T160-552q0-150 96.5-239T480-880q127 0 223.5 89T800-552q0 100-79.5 217.5T480-80Zm0-472Z",
			},
		),
	)
}

// ZeroPoints is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Acounter_0%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func ZeroPoints() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon points points-0",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.266-80q-82.734 0-155.5-31.5t-127.266-86q-54.5-54.5-86-127.341Q80-397.681 80-480.5q0-82.819 31.5-155.659Q143-709 197.5-763t127.341-85.5Q397.681-880 480.5-880q82.819 0 155.659 31.5Q709-817 763-763t85.5 127Q880-563 880-480.266q0 82.734-31.5 155.5T763-197.684q-54 54.316-127 86Q563-80 480.266-80Zm.234-60Q622-140 721-239.5t99-241Q820-622 721.188-721 622.375-820 480-820q-141 0-240.5 98.812Q140-622.375 140-480q0 141 99.5 240.5t241 99.5Zm-.5-340Zm-55 213h110q24.75 0 42.375-17.625T595-327v-307q0-24.75-17.625-42.375T535-694H425q-24.75 0-42.375 17.625T365-634v307q0 24.75 17.625 42.375T425-267Zm0-367h110v307H425v-307Z",
			},
		),
	)
}

// OnePoint is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Acounter_1%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func OnePoint() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon points points-1",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.266-80q-82.734 0-155.5-31.5t-127.266-86q-54.5-54.5-86-127.341Q80-397.681 80-480.5q0-82.819 31.5-155.659Q143-709 197.5-763t127.341-85.5Q397.681-880 480.5-880q82.819 0 155.659 31.5Q709-817 763-763t85.5 127Q880-563 880-480.266q0 82.734-31.5 155.5T763-197.684q-54 54.316-127 86Q563-80 480.266-80Zm.234-60Q622-140 721-239.5t99-241Q820-622 721.188-721 622.375-820 480-820q-141 0-240.5 98.812Q140-622.375 140-480q0 141 99.5 240.5t241 99.5Zm-.5-340Zm-8 203h60v-406H389v60h83v346Z",
			},
		),
	)
}

// TwoPoints is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Acounter_2%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func TwoPoints() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon points points-2",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.266-80q-82.734 0-155.5-31.5t-127.266-86q-54.5-54.5-86-127.341Q80-397.681 80-480.5q0-82.819 31.5-155.659Q143-709 197.5-763t127.341-85.5Q397.681-880 480.5-880q82.819 0 155.659 31.5Q709-817 763-763t85.5 127Q880-563 880-480.266q0 82.734-31.5 155.5T763-197.684q-54 54.316-127 86Q563-80 480.266-80Zm.234-60Q622-140 721-239.5t99-241Q820-622 721.188-721 622.375-820 480-820q-141 0-240.5 98.812Q140-622.375 140-480q0 141 99.5 240.5t241 99.5Zm-.5-340ZM365-277h230v-60H425v-115h110q24 0 42-18t18-42v-111q0-24-18-42t-42-18H365v60h170v111H425q-24 0-42 18t-18 42v175Z",
			},
		),
	)
}

// ThreePoints is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Acounter_3%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func ThreePoints() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon points points-3",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.266-80q-82.734 0-155.5-31.5t-127.266-86q-54.5-54.5-86-127.341Q80-397.681 80-480.5q0-82.819 31.5-155.659Q143-709 197.5-763t127.341-85.5Q397.681-880 480.5-880q82.819 0 155.659 31.5Q709-817 763-763t85.5 127Q880-563 880-480.266q0 82.734-31.5 155.5T763-197.684q-54 54.316-127 86Q563-80 480.266-80Zm.234-60Q622-140 721-239.5t99-241Q820-622 721.188-721 622.375-820 480-820q-141 0-240.5 98.812Q140-622.375 140-480q0 141 99.5 240.5t241 99.5Zm-.5-340ZM365-277h170q24 0 42-18t18-42v-87q0-27-14.5-42.5T546-482q20 0 34.5-13.5T595-537v-86q0-24-18-42t-42-18H365v60h170v111h-87v60h87v115H365v60Z",
			},
		),
	)
}

// FourPoints is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Acounter_4%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func FourPoints() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon points points-4",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.266-80q-82.734 0-155.5-31.5t-127.266-86q-54.5-54.5-86-127.341Q80-397.681 80-480.5q0-82.819 31.5-155.659Q143-709 197.5-763t127.341-85.5Q397.681-880 480.5-880q82.819 0 155.659 31.5Q709-817 763-763t85.5 127Q880-563 880-480.266q0 82.734-31.5 155.5T763-197.684q-54 54.316-127 86Q563-80 480.266-80Zm.234-60Q622-140 721-239.5t99-241Q820-622 721.188-721 622.375-820 480-820q-141 0-240.5 98.812Q140-622.375 140-480q0 141 99.5 240.5t241 99.5Zm-.5-340Zm55 203h60v-406h-60v171H425v-171h-60v231h170v175Z",
			},
		),
	)
}

// FivePoints is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Acounter_5%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func FivePoints() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon points points-5",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.266-80q-82.734 0-155.5-31.5t-127.266-86q-54.5-54.5-86-127.341Q80-397.681 80-480.5q0-82.819 31.5-155.659Q143-709 197.5-763t127.341-85.5Q397.681-880 480.5-880q82.819 0 155.659 31.5Q709-817 763-763t85.5 127Q880-563 880-480.266q0 82.734-31.5 155.5T763-197.684q-54 54.316-127 86Q563-80 480.266-80Zm.234-60Q622-140 721-239.5t99-241Q820-622 721.188-721 622.375-820 480-820q-141 0-240.5 98.812Q140-622.375 140-480q0 141 99.5 240.5t241 99.5Zm-.5-340ZM365-277h170q24 0 42-18t18-42v-115q0-24-18-42t-42-18H425v-111h170v-60H365v231h170v115H365v60Z",
			},
		),
	)
}

// OneStep is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Acounter_1%3AFILL%401%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func OneStep() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon steps steps-1",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.266-80q-82.734 0-155.5-31.5t-127.266-86q-54.5-54.5-86-127.341Q80-397.681 80-480.5q0-82.819 31.5-155.659Q143-709 197.5-763t127.341-85.5Q397.681-880 480.5-880q82.819 0 155.659 31.5Q709-817 763-763t85.5 127Q880-563 880-480.266q0 82.734-31.5 155.5T763-197.684q-54 54.316-127 86Q563-80 480.266-80ZM472-277h60v-406H389v60h83v346Z",
			},
		),
	)
}

// TwoSteps is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Acounter_2%3AFILL%401%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func TwoSteps() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon steps steps-2",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.266-80q-82.734 0-155.5-31.5t-127.266-86q-54.5-54.5-86-127.341Q80-397.681 80-480.5q0-82.819 31.5-155.659Q143-709 197.5-763t127.341-85.5Q397.681-880 480.5-880q82.819 0 155.659 31.5Q709-817 763-763t85.5 127Q880-563 880-480.266q0 82.734-31.5 155.5T763-197.684q-54 54.316-127 86Q563-80 480.266-80ZM365-277h230v-60H425v-115h110q24 0 42-18t18-42v-111q0-24-18-42t-42-18H365v60h170v111H425q-24 0-42 18t-18 42v175Z",
			},
		),
	)
}

// ThreeSteps is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Acounter_3%3AFILL%401%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func ThreeSteps() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon steps steps-3",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.266-80q-82.734 0-155.5-31.5t-127.266-86q-54.5-54.5-86-127.341Q80-397.681 80-480.5q0-82.819 31.5-155.659Q143-709 197.5-763t127.341-85.5Q397.681-880 480.5-880q82.819 0 155.659 31.5Q709-817 763-763t85.5 127Q880-563 880-480.266q0 82.734-31.5 155.5T763-197.684q-54 54.316-127 86Q563-80 480.266-80ZM365-277h170q24 0 42-18t18-42v-87q0-27-14.5-42.5T546-482q20 0 34.5-13.5T595-537v-86q0-24-18-42t-42-18H365v60h170v111h-87v60h87v115H365v60Z",
			},
		),
	)
}

// FourSteps is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Acounter_4%3AFILL%401%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func FourSteps() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon steps steps-4",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.266-80q-82.734 0-155.5-31.5t-127.266-86q-54.5-54.5-86-127.341Q80-397.681 80-480.5q0-82.819 31.5-155.659Q143-709 197.5-763t127.341-85.5Q397.681-880 480.5-880q82.819 0 155.659 31.5Q709-817 763-763t85.5 127Q880-563 880-480.266q0 82.734-31.5 155.5T763-197.684q-54 54.316-127 86Q563-80 480.266-80ZM535-277h60v-406h-60v171H425v-171h-60v231h170v175Z",
			},
		),
	)
}

// SixSteps is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Acounter_6%3AFILL%401%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func SixSteps() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon steps steps-6",
		},
		svg.Path(
			hypp.HProps{
				"d": "M480.266-80q-82.734 0-155.5-31.5t-127.266-86q-54.5-54.5-86-127.341Q80-397.681 80-480.5q0-82.819 31.5-155.659Q143-709 197.5-763t127.341-85.5Q397.681-880 480.5-880q82.819 0 155.659 31.5Q709-817 763-763t85.5 127Q880-563 880-480.266q0 82.734-31.5 155.5T763-197.684q-54 54.316-127 86Q563-80 480.266-80ZM425-277h110q24 0 42-18t18-42v-115q0-24-18-42t-42-18H425v-111h127v-60H425q-24 0-42 18t-18 42v286q0 24 18 42t42 18Zm0-175h110v115H425v-115Z",
			},
		),
	)
}

// PlayerTurnArrow is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Aarrow_left_alt%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func PlayerTurnArrow() *hypp.VNode {
	return svg.Svg(
		hypp.HProps{
			"viewBox": "0 -960 960 960",
			"class":   "icon player-turn-arrow",
		},
		svg.Path(
			hypp.HProps{
				"d": "M400-240 160-480l241-241 43 42-169 169h526v60H275l168 168-43 42Z",
			},
		),
	)
}
