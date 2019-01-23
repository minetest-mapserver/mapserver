package tilerenderer


import (
  "image"
	"image/draw"
	"image/png"
  "image/color"
  "bytes"
)

func CreateBlankTile(c color.RGBA) ([]byte){
  rect := image.Rectangle{
    image.Point{0, 0},
    image.Point{IMG_SIZE, IMG_SIZE},
  }

  img := image.NewNRGBA(rect)
  draw.Draw(img, rect, &image.Uniform{c}, image.ZP, draw.Src)

  buf := new(bytes.Buffer)
  png.Encode(buf, img)

  return buf.Bytes()
}
