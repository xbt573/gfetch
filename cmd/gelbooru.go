package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/xbt573/gfetch/pkg/booru"
	"github.com/xbt573/gfetch/pkg/booru/gelbooru"
)

var gelbooruCmd = cobra.Command{
	Use:   "gelbooru [tags]",
	Short: "batch photo downloading from gelbooru",
	Long:  "batch photo downloading from gelbooru",
	Run: func(_ *cobra.Command, args []string) {
		gb := gelbooru.Gelbooru{}

		posts := gb.Search(booru.BooruSearchOptions{
			Tags:  args,
			Count: limit,
		})

		if posts.Error != nil {
			fmt.Println(posts.Error)
			os.Exit(1)
		}

		wg := new(sync.WaitGroup)

		for _, post := range posts.Posts {
			wg.Add(1)
			go func(post booru.BooruPost) {
				Download(post)
				wg.Done()
			}(post)
		}

		wg.Wait()
	},
}
