package astiav

//#cgo pkg-config: libswresample
//#include <libswresample/swresample.h>
import "C"
import "unsafe"

type SwrContext struct {
	c *C.struct_SwrContext
}

func (c *SwrContext) Free() {
	C.swr_free(&c.c)
}

func (c *SwrContext) ConvertFrame(outFrame *Frame, inFrame *Frame) (n int, err error) {
	n = int(C.swr_convert(c.c, (**C.uint8_t)(&outFrame.c.data[0]), outFrame.c.nb_samples, (**C.uint8_t)(&inFrame.c.data[0]), inFrame.c.nb_samples))
	if n < 0 {
		err = ErrInvaliddata
	}
	return
}

func (c *SwrContext) Convert(outData [][]byte, outSize int, inData [][]byte, inSize int) (n int, err error) {
	n = int(C.swr_convert(c.c, (**C.uint8_t)(unsafe.Pointer(&outData[0][0])), C.int(outSize), (**C.uint8_t)(unsafe.Pointer(&inData[0][0])), C.int(inSize)))
	if n < 0 {
		err = ErrInvaliddata
	}
	return
}
func AllocSwrContext(outChannelLayout *ChannelLayout, outSampleFormat SampleFormat, outSampleRate int,
	inChannelLayout *ChannelLayout, inSampleFormat SampleFormat, intSampleRate int, logLevel LogLevel) (ctx *SwrContext, err error) {
	ctx = &SwrContext{
		c: nil,
	}
	err = newError(C.swr_alloc_set_opts2(&ctx.c, outChannelLayout.c, C.enum_AVSampleFormat(outSampleFormat),
		C.int(outSampleRate), inChannelLayout.c, C.enum_AVSampleFormat(inSampleFormat), C.int(intSampleRate),
		C.int(logLevel), nil))
	if err != nil {
		return
	}
	err = newError(C.swr_init(ctx.c))
	return
}
