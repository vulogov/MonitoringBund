package banner

import (
	"github.com/common-nighthawk/go-figure"

	"github.com/vulogov/monitoringbund/internal/conf"
)

func Banner(txt string) {
	if *conf.VBanner {
		PrintBanner(txt)
	}
}

func PrintBanner(txt string) {
	if *conf.Color {
		b := figure.NewColorFigure(txt, "", "yellow", false)
		b.Print()
	} else {
		b := figure.NewFigure(txt, "", true)
		b.Print()
	}
}
