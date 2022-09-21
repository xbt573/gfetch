package gelbooru

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/xbt573/gfetch/pkg/booru"
)

type GelbooruResponseAttributes struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

type GelbooruPost struct {
	Id            int    `json:"id"`
	CreatedAt     string `json:"created_at"`
	Score         int    `json:"score"`
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	MD5           string `json:"md5"`
	Directory     string `json:"directory"`
	Image         string `json:"image"`
	Rating        string `json:"rating"`
	Source        string `json:"source"`
	Change        int64  `json:"change"`
	Owner         string `json:"owner"`
	CreatorID     int    `json:"creator_id"`
	ParentID      int    `json:"parent_id"`
	Sample        int    `json:"sample"`
	PreviewWidth  int    `json:"preview_width"`
	PreviewHeight int    `json:"preview_height"`
	Tags          string `json:"tags"`
	Title         string `json:"title"`
	HasNotes      string `json:"has_notes"`
	HasComments   string `json:"has_comments"`
	FileURL       string `json:"file_url"`
	PreviewURL    string `json:"preview_url"`
	SampleURL     string `json:"sample_url"`
	SampleHeight  int    `json:"sample_height"`
	SampleWidth   int    `json:"sample_width"`
	Status        string `json:"status"`
	PostLocker    int    `json:"post_locked"`
	HasChildren   string `json:"has_children"`
}

type GelbooruResponse struct {
	Attributes GelbooruResponseAttributes `json:"@attributes"`
	Post       []GelbooruPost             `json:"post"`
}

type Gelbooru struct{}

func (Gelbooru) Search(opts booru.BooruSearchOptions) booru.BooruResponse {
	values := url.Values{}

	values.Set("page", "dapi")
	values.Set("s", "post")
	values.Set("q", "index")
	values.Set("json", "1")

	if len(opts.Tags) > 0 {
		values.Set("tags", strings.Join(opts.Tags, " "))
	}

	if opts.Count != 0 {
		values.Set("limit", strconv.Itoa(opts.Count))
	}

	res := booru.BooruResponse{}

	req, err := http.Get("https://gelbooru.com/index.php?" + values.Encode())
	if err != nil {
		res.Error = err
		return res
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		res.Error = err
		return res
	}

	var gres GelbooruResponse
	err = json.Unmarshal(body, &gres)
	if err != nil {
		res.Error = err
		return res
	}

	for _, task := range gres.Post {
		res.Posts = append(res.Posts, booru.BooruPost{
			Tags:     strings.Split(task.Tags, " "),
			FileUrl:  task.FileURL,
			FileName: task.Image,
		})
	}

	return res
}
