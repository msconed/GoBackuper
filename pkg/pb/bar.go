package pb

import (
	"github.com/schollz/progressbar/v3"
)

var bar *progressbar.ProgressBar

func InitProgressBar(maxValue int64, description string) {
	bar = progressbar.Default(maxValue, description)
}

func GetProgressBar() *progressbar.ProgressBar {
	return bar
}
