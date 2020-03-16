package web

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httputil"
)

// DownloadImage downloads an image from a certain url
func DownloadImage(imageURL string) <-chan string {
	image := make(chan string, 1)

	log.Println("Downloading image from: " + imageURL)

	resp, err := http.Get(imageURL)
	if err != nil {
		log.Println("Failed to download image from: '" + imageURL + "', because: '" + err.Error() + "'")
		bytes, _ := httputil.DumpResponse(resp, true)
		log.Println("Response: " + string(bytes))
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if buf.Len() > 0 {
		log.Println("Downloaded image from: " + imageURL + " of size: " + string(buf.Len()))
	}

	image <- buf.String()

	return image
}
