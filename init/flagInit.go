package init

import (
	"fileDownload/args"
	"flag"
)

func init() {
	flag.StringVar(&args.Url, "url", "", "set download url")
}
