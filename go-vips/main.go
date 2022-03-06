package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	vips "github.com/davidbyttow/govips/v2"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}

func Resize(inPath string, factor float64) {

	image, err := vips.NewImageFromFile(inPath)
	checkError(err)

	err = image.Resize(factor, -1)
	checkError(err)

	ep := vips.NewDefaultExportParams()
	imageBytes, _, err := image.Export(ep)

	checkError(err)

	var format string

	switch image.Format() {

	case vips.ImageTypeJPEG:
		format = "jpeg"
	case vips.ImageTypeWEBP:
		format = "webp"
	case vips.ImageTypePNG:
		format = "png"
	}

	outPath := fmt.Sprintf("../images/output.%g.%s", factor, format)
	err = ioutil.WriteFile(outPath, imageBytes, 0644)
	checkError(err)

}

func Repeat(N int, inPath string, factor float64) int64 {
	var sum int64

	for i := 0; i < N; i++ {
		start := time.Now()
		Resize(inPath, factor)
		sum += time.Since(start).Milliseconds()
	}

	return sum / int64(N)

}

type input struct {
	path   string
	factor float64
	format string
}

func main() {
	vips.Startup(nil)
	defer vips.Shutdown()

	// Chart Average Latencies for different sizes
	jpegItems := make([]opts.BarData, 0)
	pngItems := make([]opts.BarData, 0)
	webpItems := make([]opts.BarData, 0)

	inputs := []input{
		{
			path:   "../images/input.jpg",
			factor: 0.5,
			format: "jpeg",
		},
		{
			path:   "../images/input.jpg",
			factor: 0.1,
			format: "jpeg",
		},
		{
			path:   "../images/input.jpg",
			factor: 0.05,
			format: "jpeg",
		},
		{
			path:   "../images/input.jpg",
			factor: 0.01,
			format: "jpeg",
		},
		{
			path:   "../images/input.png",
			factor: 0.5,
			format: "png",
		},
		{
			path:   "../images/input.png",
			factor: 0.1,
			format: "png",
		},
		{
			path:   "../images/input.png",
			factor: 0.05,
			format: "png",
		},
		{
			path:   "../images/input.png",
			factor: 0.01,
			format: "png",
		},
		{
			path:   "../images/input.webp",
			factor: 0.5,
			format: "webp",
		},
		{
			path:   "../images/input.webp",
			factor: 0.1,
			format: "webp",
		},
		{
			path:   "../images/input.webp",
			factor: 0.05,
			format: "webp",
		},
		{
			path:   "../images/input.webp",
			factor: 0.01,
			format: "webp",
		},
	}

	for _, i := range inputs {
		value := Repeat(20, i.path, i.factor)
		if i.format == "jpeg" {
			jpegItems = append(jpegItems, opts.BarData{
				Value: value,
			})
		} else if i.format == "png" {
			pngItems = append(pngItems, opts.BarData{
				Value: value,
			})
		} else {
			webpItems = append(webpItems, opts.BarData{
				Value: value,
			})
		}
	}

	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithLegendOpts(opts.Legend{
		Show: true,
	}), charts.WithTitleOpts(opts.Title{
		Title:    "Go Vips Latency",
		Subtitle: "for resize operations",
	}), charts.WithYAxisOpts(
		opts.YAxis{
			AxisLabel: &opts.AxisLabel{
				Show:      true,
				Formatter: "{value} ms",
			}}))

	bar.SetXAxis([]string{
		"Resize Factor = 0.5x",
		"Resize Factor = 0.1x",
		"Resize Factor = 0.05x",
		"Resize Factor = 0.01x",
	}).
		AddSeries("JPEG", jpegItems, charts.WithLabelOpts(opts.Label{
			Show:      true,
			Formatter: "JPEG",
		})).
		AddSeries("PNG", pngItems, charts.WithLabelOpts(opts.Label{
			Show:      true,
			Formatter: "PNG",
		})).
		AddSeries("WEBP", webpItems, charts.WithLabelOpts(opts.Label{
			Show:      true,
			Formatter: "WEBP",
		}))

	// export html file for the Bar Chart
	f, _ := os.Create("go-vips-latencies.html")
	bar.Render(f)
}
