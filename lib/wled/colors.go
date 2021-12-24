package wled


import "github.com/lucasb-eyer/go-colorful"


/* Convert a colorful color into a wled 3-byte color for
usage with the rest of the wled library.
*/
func ToWled(c colorful.Color) WledColor {

	r,g,b := c.RGB255()

	return WledColor{r,g,b}

}

// creates a colorful color from a wled color.
// might be useful if reading from current WLED state.
func FromWled(c WledColor) colorful.Color {
	return colorful.FastLinearRgb(float64(c[0]) / 255.0, float64(c[1]) / 255.0, float64(c[2])/255.0)
}

