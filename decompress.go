package turbo

/*
#cgo LDFLAGS: -lturbojpeg
#include <turbojpeg.h>
*/
import "C"

func Decompress(encoded []byte) (*Image, error) {
	decoder := C.tjInitDecompress()
	defer C.tjDestroy(decoder)

	width := C.int(0)
	height := C.int(0)
	sampling := C.int(0)
	colorspace := C.int(0)

	err := makeError(decoder, C.tjDecompressHeader3(decoder, (*C.uchar)(&encoded[0]), C.ulong(len(encoded)), &width, &height, &sampling, &colorspace))
	if err != nil {
		return nil, err
	}

	outBuf := make([]byte, width*height*4)
	stride := C.int(width * 4)

	// int tjDecompress2(tjhandle handle, const unsigned char *jpegBuf, unsigned long jpegSize, unsigned char *dstBuf,
	// int width, int pitch, int height, int pixelFormat, int flags);
	err = makeError(decoder, C.tjDecompress2(decoder, (*C.uchar)(&encoded[0]), C.ulong(len(encoded)), (*C.uchar)(&outBuf[0]), width, stride, height, C.int(PixelFormatRGBA), 0))
	if err != nil {
		return nil, err
	}

	img := &Image{
		Width:  int(width),
		Height: int(height),
		Stride: int(stride),
		Pixels: outBuf,
	}
	return img, nil
}
