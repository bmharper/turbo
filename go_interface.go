package turbo

import "image"

// Convert a Go image.Image into a turbo.Image
// If allowDeepClone is true, and the source image is type NRGBA or RGBA,
// then the resulting Image points directly to the pixel buffer of the source image.
func FromImage(src image.Image, allowDeepClone bool) *Image {
	dst := &Image{
		Width:  src.Bounds().Dx(),
		Height: src.Bounds().Dy(),
		Stride: src.Bounds().Dx() * 4,
	}
	switch v := src.(type) {
	case *image.RGBA:
		if allowDeepClone {
			dst.Pixels = v.Pix
		} else {
			dst.Pixels = make([]byte, dst.Stride*dst.Height)
			copy(dst.Pixels, v.Pix)
		}
		return dst
	case *image.NRGBA:
		if allowDeepClone {
			dst.Pixels = v.Pix
		} else {
			dst.Pixels = make([]byte, dst.Stride*dst.Height)
			copy(dst.Pixels, v.Pix)
		}
		return dst
	}

	// This must be super slow - I haven't tested
	dst.Pixels = make([]byte, dst.Stride*dst.Height)
	p := 0
	for y := 0; y < dst.Height; y++ {
		for x := 0; x < dst.Width; x++ {
			r, g, b, a := src.At(x, y).RGBA()
			dst.Pixels[p] = byte(r >> 8)
			dst.Pixels[p+1] = byte(g >> 8)
			dst.Pixels[p+2] = byte(b >> 8)
			dst.Pixels[p+3] = byte(a >> 8)
			p += 4
		}
	}

	return dst
}
