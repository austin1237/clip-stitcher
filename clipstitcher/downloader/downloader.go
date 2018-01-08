package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func downloadClip(url string, fileName string, response chan error) {
	out, err := os.Create("./tmpClips/" + fileName)
	defer out.Close()
	if err != nil {
		response <- err
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		resp.Body.Close()
		response <- err
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		response <- err
		return
	}
	response <- nil
}

func DownloadClips(urls []string) []error {
	downloadResponses := make(chan error, len(urls))
	errs := []error{}
	for i, url := range urls {
		fileName := "clip" + strconv.Itoa(i) + ".mp4"
		go downloadClip(url, fileName, downloadResponses)
	}
	fmt.Println("file downloads started")
	for err := range downloadResponses {
		errs = append(errs, err)
		if err != nil {
			fmt.Println("Download error " + err.Error())
		}
		if len(errs) == len(urls) {
			close(downloadResponses)
		}
	}
	fmt.Println("file downloads finished")

	return errs
}
