package processing

import (
	"github.com/imgproxy/imgproxy/v3/imagedata"
	"github.com/imgproxy/imgproxy/v3/options"
	"github.com/imgproxy/imgproxy/v3/vips"
)

func scale(pctx *pipelineContext, img *vips.Image, po *options.ProcessingOptions, imgdata *imagedata.ImageData) error {
	if pctx.wscale != 1 || pctx.hscale != 1 {
		wscale, hscale := pctx.wscale, pctx.hscale
		if (pctx.angle+po.Rotate)%180 == 90 {
			wscale, hscale = hscale, wscale
		}

		if err := img.Resize(wscale, hscale); err != nil {
			return err
		}
	}

	return img.CopyMemory()
}
