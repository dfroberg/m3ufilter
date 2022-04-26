package writer

import (
	"fmt"
	"io"

	"github.com/dfroberg/m3ufilter/logger"
	"github.com/dfroberg/m3ufilter/m3u"
)

var log = logger.Get()

func WriteOutput(Output string, w io.Writer, streams m3u.Streams) {
	switch Output {
	case "m3u":
		writeM3U(w, streams)
	case "csv":
		writeCsv(w, streams)
	default:
		panic(fmt.Errorf("output type unknown expected m3u|csv, got %s", Output))
	}
}
