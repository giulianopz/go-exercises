package model

type Base struct {
	Id       int    `json:"id"`
	Deleted  bool   `json:"deleted"`
	BaseType string `json:"type"`
	By       string `json:"by"`
	Time     int    `json:"time"`
	Dead     bool   `json:"dead"`
	Kids     []int  `json:"kids"`
}

type Story struct {
	Id          int    `json:"id"`
	Deleted     bool   `json:"deleted"`
	BaseType    string `json:"type"`
	By          string `json:"by"`
	Time        int    `json:"time"`
	Dead        bool   `json:"dead"`
	Kids        []int  `json:"kids"`
	Descendants int    `json:"descendants"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}
