package wikiscrape

type Data struct {
	Url   string   `json:"url"`
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
	Text  string   `json:"text"`
}
