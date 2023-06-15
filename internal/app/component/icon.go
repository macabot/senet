package component

import (
	"github.com/macabot/hypp"
	"github.com/macabot/hypp/tag/svg"
)

// protectedIcon is based on https://fonts.google.com/icons?selected=Material+Symbols+Outlined:protectedIcon:FILL@0;wght@400;GRAD@0;opsz@48
func protectedIcon() *hypp.VNode {
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

// blockingIcon is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Afront_hand%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func blockingIcon() *hypp.VNode {
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

// returnToStartIcon is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Amove%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func returnToStartIcon() *hypp.VNode {
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

// startIcon is based on https://fonts.google.com/icons?selected=Material%20Symbols%20Outlined%3Alocation_on%3AFILL%400%3Bwght%40400%3BGRAD%400%3Bopsz%4048
func startIcon() *hypp.VNode {
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
