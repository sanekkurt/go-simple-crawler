package crawler

type ResultData struct {
	URL   string `json:"url"`
	Title string `json:"title,omitempty"`
	Error string `json:"error,omitempty"`
}
