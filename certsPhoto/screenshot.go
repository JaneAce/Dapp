package certsPhoto

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"os"
)
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func Screenshot(certName string){
	//自定义截图
	fmt.Println("证书名称",certName)
	img, err := screenshot.Capture(380, 300, 700, 1000)
	if err != nil {
		panic(err)
	}
	save(img, "F:\\DAPP\\certsPhoto\\"+certName+".png")

	//获取所有活动屏幕
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("没有发现活动的显示器")
	}

	//全屏截取所有活动屏幕
	if n > 0 {
		for i := 0; i < n; i++ {
			img, err = screenshot.CaptureDisplay(i)
			if err != nil {
				panic(err)
			}
			fileName := fmt.Sprintf("第%d屏幕截图.png", i)
			save(img, fileName)
		}
	}

	//使用Union获取指定屏幕 矩形最小点和最大点
	var all image.Rectangle = image.Rect(0, 0, 0, 0)
	bounds1 := screenshot.GetDisplayBounds(0)
	all = bounds1.Union(all)
	fmt.Println(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
}
