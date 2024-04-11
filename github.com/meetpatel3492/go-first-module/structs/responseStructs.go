package structs

type TopStoriesResponse []int

type ArticleResponseById struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Parts       []int  `json:"parts"`
	Score       int    `json:"score"`
	Text        string `json:"text"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}
