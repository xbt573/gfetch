package cmd

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"

	"github.com/xbt573/gfetch/pkg/booru"
)

func Download(post booru.BooruPost) {
	res, err := http.Get(post.FileUrl)
	if err != nil {
		if verbose {
			fmt.Printf("download %v: %v\n", post.FileName, err)
		}

		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		if verbose {
			fmt.Printf("download %v: %v\n", post.FileName, err)
		}

		return
	}

	err = os.WriteFile(post.FileName, data, fs.ModePerm)
	if err != nil {
		if verbose {
			fmt.Printf("download %v: %v\n", post.FileName, err)
		}

		return
	}

	if verbose {
		fmt.Printf("download %v: finished\n", post.FileName)
	}
}
