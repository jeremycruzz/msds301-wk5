package wikiscrape

type Data struct {
	Url   string   `json:"url"`
	Title string   `json:"title"`
	Text  string   `json:"text"`
	Tags  []string `json:"tags"`
}
