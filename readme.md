# Turbo (TurboJPEG Go wrapper)

This is a very thin wrapper around turbojpeg.

Why not use https://github.com/pixiv/go-libjpeg ?

Because go-libjpeg is built to use the libjpeg-compatible API of either libjpeg or
libjpeg-turbo. That API does not allow one to specify the chroma sub-sampling, so
you're unable to compress to 4:2:0. That was the sole reason for building this
package -- to enable compression to 420 chroma sub-sampling.

This wrapper links explicitly to TurboJPEG. There is no fallback to libjpeg.

### How to use

```go
import "github.com/bmharper/turbo"

func compressImage(width, height int, rgba []byte) {
	raw := turbo.Image{
		Width: width,
		Height: height,
		Stride: width * 4,
		RGBA: rgba,
	}
	params := turbo.MakeCompressParams(turbo.PixelFormatRGBA, turbo.Sampling420, 35, 0)
	jpg, err := turbo.Compress(&raw, params)
}

func decompressImage(jpg []byte) (*Image, error) {
	return turbo.Decompress(jpg)
}
```