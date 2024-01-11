package localminio

import (
	"github.com/schollz/progressbar/v3"
)

type ProgressListener struct {
	progress *progressbar.ProgressBar
}

func NewDefaultProgressListener() *ProgressListener {
	return &ProgressListener{}
}

func (p *ProgressListener) ProgressChanged() {

}
