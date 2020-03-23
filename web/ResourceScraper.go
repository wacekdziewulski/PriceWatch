package web

import (
	"bytes"
	"net/http"
	"net/http/httputil"

	"github.com/sirupsen/logrus"
)

// DownloadImage downloads an image from a certain url
func DownloadImage(imageURL string) <-chan string {
	image := make(chan string, 1)

	logrus.Debug("Downloading image from: ", imageURL)

	resp, err := http.Get(imageURL)
	if err != nil {
		bytes, _ := httputil.DumpResponse(resp, true)
		logrus.Warnf("Failed to download image from: %s, because of: %+v. HttpResponse: %s", imageURL, err, bytes)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if buf.Len() > 0 {
		logrus.Infof("Downloaded image from: %s of size: %d", imageURL, buf.Len())
	}

	image <- buf.String()

	return image
}
