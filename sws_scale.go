package astiav

//#cgo pkg-config: libswscale
//#include <libswscale/swscale.h>
import "C"

type SwsFilter struct {
	c *C.struct_SwsFilter
}

type SwsContext struct {
	c *C.struct_SwsContext
}

func (c *SwsContext) Free() {
	C.sws_freeContext(c.c)
}
func (c *SwsContext) Scale(srcFrame *Frame, dstFrame *Frame, srcSliceY int, srcSliceH int) error {
	return newError(C.sws_scale(
		c.c,
		srcFrame.c.data,
		srcFrame.c.linesize,
		C.int(srcSliceY),
		C.int(srcSliceH),
		dstFrame.c.data,
		dstFrame.c.linesize,
	))
}

//	func allocSwsContextSetOptionsFromDictionary(d *astiav.Dictionary) *SwsContext {
//		var c *C.struct_SwsContext
//		c = C.sws_alloc_context()
//		if c == nil {
//			return nil
//		}
//		astiav.NewDictionary()
//		C.av_opt_set_dict(c.c, d)
//		return &SwsContext{c: c}
//	}
func GetSwsContext(srcW int, srcH int, srcFormat PixelFormat, dstW int, dstH int, dstFormat PixelFormat,
	flags int, srcFilter *SwsFilter, dstFilter *SwsFilter, param []float64) *SwsContext {
	var cParam *C.double
	if len(param) > 0 {
		cParam = (*C.double)(&param[0])
	}
	var cSrcFilter *C.struct_SwsFilter
	if srcFilter != nil {
		cSrcFilter = srcFilter.c
	}
	var cDstFilter *C.struct_SwsFilter
	if cDstFilter != nil {
		cDstFilter = dstFilter.c
	}
	c := C.sws_getContext(
		C.int(srcW),
		C.int(srcH),
		C.enum_AVPixelFormat(srcFormat),
		C.int(dstW),
		C.int(dstH),
		C.enum_AVPixelFormat(dstFormat),
		C.int(flags),
		cSrcFilter,
		cDstFilter,
		cParam,
	)
	if c == nil {
		return nil
	}
	return &SwsContext{c: c}
}
