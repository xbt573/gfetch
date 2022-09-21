package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/xbt573/gfetch/pkg/booru"
)

func Download(post booru.BooruPost) {
	data, err := http.Get(post.FileUrl)
	if err != nil {
		if verbose {
			fmt.Printf("download %v: %v", post.FileName, err)
		}

		return
	}
	defer data.Body.Close()

	f, err := os.Create(post.FileName)
	if err != nil {
		if verbose {
			fmt.Printf("download %v: %v\n", post.FileName, err)
		}

		os.Remove(post.FileName)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, data.Body)
	if err != nil {
		if verbose {
			fmt.Printf("download %v: %v\n", post.FileName, err)
		}

		os.Remove(post.FileName)
		return
	}

	if verbose {
		fmt.Printf("download %v: finished\n", post.FileName)
	}
}
