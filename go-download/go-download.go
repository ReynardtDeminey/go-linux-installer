package download

import (
	"log"

	"github.com/cavaliercoder/grab"
)

// GoDownload downloads the Go binary
func GoDownload() {
	_, err := grab.Get("", "https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz")
	if err != nil {
		log.Fatal(err)
	}
}
