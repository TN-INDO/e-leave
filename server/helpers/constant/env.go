package constant

import (
	"os"
)

var (
	GOPATH     = os.Getenv("GOPATH")
	GOAPP      = os.Getenv("GOAPP")
	GOENV      = os.Getenv("GOENV")
	GOPWD      = os.Getenv("GOPWD")
	GOPWDRESET = os.Getenv("GOPWDRESET")
)
