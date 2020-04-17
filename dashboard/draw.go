package dashboard

import (
	"fmt"
	"github.com/medenzon/gobadge/dashboard/svg"
	geo "github.com/medenzon/gobadge/dashboard/svg/geometry"
	"github.com/medenzon/gobadge/dashboard/svg/style"
	"log"
	"math"
	"strings"
)

func (view *View) draw() {

	index, count := 0, len(view.Badges)
	split := float64(count) / float64(view.Columns)
	rows, cols := int(math.Ceil(split)), view.Columns

	for col := 0; col < cols; col++ {

		for row := 0; row < rows; row++ {
			if index < count {
				log.Println(index, row)
				x := ((badgeW + badgeH) * col) + hoff
				y := ((badgeH + vpad) * row) + voff

				badge := view.Badges[index]
				origin := geo.Coordinate{X: float64(x), Y: float64(y)}
				badge.draw(view.Canvas, origin)

				index++
			}
		}
	}
}

func (badge Badge) draw(canvas *svg.Canvas, origin geo.Coordinate) error {

	gradient := stockBadgeGradient("#000000", 0.1, "#BBBBBB", 0.1)
	filter := stockTextFilter()

	labelText := stockBadgeLabel(badge.Label)
	tagText := stockBadgeTag(badge.Tag)
	path1 := stockBadgeBlock()
	path2 := stockBadgeColor(colormap(badge.Color))
	rect := stockBadgeGradientRect()
	labelStyle := style.Text{ID: badge.Label, Filter: filter, Text: labelText}
	tagStyle := style.Text{ID: badge.Tag, Filter: filter, Text: tagText}
	translation := fmt.Sprintf("translate(%f,%f)", origin.X, origin.Y)

	output := strings.Join(
		[]string{
			`<g id="` + badge.Label + `" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd" transform="` + translation + `">`,
			`<defs>`,
			gradient.Vectorize(),
			filter.Vectorize(),
			labelText.Vectorize(),
			tagText.Vectorize(),
			`</defs>`,
			`<g>`,
			path1.Vectorize(),
			path2.Vectorize(),
			rect.Vectorize(),
			labelStyle.Vectorize(),
			tagStyle.Vectorize(),
			`</g>`,
			`</g>`,
		}, "\n")

	_, err := fmt.Fprint(canvas.Writer, output)

	return err
}