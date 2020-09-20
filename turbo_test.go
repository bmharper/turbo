package turbo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func MakeRGBA(width, height int) *Image {
	img := &Image{
		Width:  width,
		Stride: width * 4,
		Height: height,
		Pixels: make([]byte, width*height*4),
	}
	buf := img.Pixels
	g := byte(0)
	b := byte(0)
	p := 0
	a := byte(255)
	for y := 0; y < height; y++ {
		r := byte(0)
		for x := 0; x < width; x++ {
			buf[p] = r
			buf[p+1] = g
			buf[p+2] = b
			buf[p+3] = a
			r += 3
			b += 5
			p += 4
		}
		g += 1
	}
	return img
}

func TestCompress(t *testing.T) {
	w := 300
	h := 200
	raw1 := MakeRGBA(w, h)
	params := MakeCompressParams(PixelFormatRGBA, Sampling444, 90, 0)
	jpg, err := Compress(raw1, params)
	t.Logf("Encode return: %v, %v", len(jpg), err)
	raw2, err := Decompress(jpg)
	t.Logf("Decode return: %v x %v, %v, %v, %v", raw2.Width, raw2.Height, raw2.Stride, len(raw2.Pixels), err)
	assert.Equal(t, &w, &raw2.Width, "Width same")
	assert.Equal(t, &h, &raw2.Height, "Height same")
	assert.Equal(t, &raw1.Stride, &raw2.Stride, "Stride same")
	//ioutil.WriteFile("test.jpg", jpg, 0660)
}
