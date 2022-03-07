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

	outPath := fmt.Sprintf("../images/output.go_vips.%g.%s", factor, format)
	err = ioutil.WriteFile(outPath, imageBytes, 0644)
	checkError(err)

}

func Repeat(N int, inPath string, factor float64) float64 {
	var sum float64

	for i := 0; i < N; i++ {
		start := time.Now()
		Resize(inPath, factor)
		sum += float64(time.Since(start).Milliseconds())
	}

	return sum / float64(N)

}

type input struct {
	path   string
	factor float64
	format string
}

func generateChart(titleOpts charts.GlobalOpts, jpegItems, pngItems, webpItems []opts.BarData) *charts.Bar {
	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithLegendOpts(opts.Legend{
		Show: true,
	}),
		titleOpts,
		charts.WithYAxisOpts(
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
	return bar
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

	titleOpts := charts.WithTitleOpts(opts.Title{
		Title:    "Go Vips Latency",
		Subtitle: "for resize operations",
	})
	bar := generateChart(titleOpts, jpegItems, pngItems, webpItems)
	// export html file for the Bar Chart
	f, _ := os.Create("go-vips-latencies.html")
	bar.Render(f)

	titleOpts = charts.WithTitleOpts(opts.Title{
		Title:    "WASM Vips Latency",
		Subtitle: "for resize operations",
	})
	// generate WASM VIPS chart from imported data
	bar = generateChart(titleOpts,
		[]opts.BarData{
			{Value: 116.21978175},
			{Value: 47.6798169},
			{Value: 25.57381305},
			{Value: 27.681828449999994},
		},
		[]opts.BarData{
			{Value: 246.77878360000005},
			{Value: 34.92498075},
			{Value: 24.464996149999997},
			{Value: 17.591525949999998},
		},
		[]opts.BarData{
			{Value: 318.08282045},
			{Value: 43.37513325},
			{Value: 30.443789099999996},
			{Value: 20.2916782},
		},
	)

	f, _ = os.Create("wasm-vips-latencies.html")
	bar.Render(f)
}
