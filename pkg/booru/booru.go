package booru

type BooruResponse struct {
	Error error
	Posts []BooruPost
}

type BooruPost struct {
	FileName string
	FileUrl  string
	Tags     []string
}

type BooruSearchOptions struct {
	Tags  []string
	Count int
}

type Booru interface {
	Search(opts BooruSearchOptions) BooruResponse
}
