package main

import (
    "fmt"
    "image"
    "image/jpeg"
    "os"
    "io"
    "github.com/nfnt/resize"
)
//const GREY_SCALE=" .'`^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
//const GREY_SCALE=" .'^\",:;Ili><~_-](\\/tjnuvXCLZmwpkhaMW&%$"
  const GREY_SCALE=" .:-=+*#%@"



func main() {
    // You can register another format here
    image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

    file, err := os.Open("Leonardo_da_Vinci_-_Mona_Lisa_(Louvre,_Paris).jpg")

    if err != nil {
        fmt.Fprintf(os.Stderr,"%v",err)
        os.Exit(1)
    }

    defer file.Close()

    pixels, err := getPixels(file)

    if err != nil {
        fmt.Fprintf(os.Stderr,"%v",err)
        os.Exit(1)
    }

    for i,_:=range(pixels){
        for _,y:=range(pixels[i]){
          fmt.Print(string(GREY_SCALE[y]))
        }
        fmt.Println()
    }
}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][]int, error) {
    Oldimg, _, err := image.Decode(file)

    if err != nil {
        return nil, err
    }
    img_width:=uint(300)
    img_height:=uint(3*img_width/8)
    img := resize.Resize(img_width, img_height, Oldimg, resize.Lanczos3)


    bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y

    var pixels [][]int
    for y := 0; y < height; y++ {
        var row []int
        for x := 0; x < width; x++ {
            row = append(row, getGreyScaleValue(img.At(x, y).RGBA()))
        }
        pixels = append(pixels, row)
    }

    return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel

//765:len(g)-1=gray:x
func getGreyScaleValue(r uint32, g uint32, b uint32, a uint32)int{
  gray:=int(r / 257)+ int(g / 257)+ int(b / 257)
  return ((len(GREY_SCALE)-1)*gray)/765
}

// Pixel struct example
