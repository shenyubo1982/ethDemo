package util

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
)

//var bar *progressbar.ProgressBar

func ProgressBarConfig(size int, Description string, part int, partIndex int) *progressbar.ProgressBar {

	barDescription := fmt.Sprintf("[cyan][%d/%d][reset] %s ", part, partIndex, Description)

	bar := progressbar.NewOptions(size,
		//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		//progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(50),
		//progressbar.OptionSetDescription("[cyan][%d/%d][reset] scanning addresses ..."),
		progressbar.OptionSetDescription(barDescription),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: "_",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	return bar
}

func ShowProgressBar(bar *progressbar.ProgressBar) {
	bar.Add(1)
}
