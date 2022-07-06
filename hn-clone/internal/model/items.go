package model

import (
	"fmt"
	"net/url"
)

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

func (s *Story) Rank(idx int) int {
	return idx + 1
}

func (s *Story) Site() string {

	u, err := url.Parse(s.Url)
	if err != nil {
		fmt.Println("ERR", "Cannot parse URL:", s.Url)
	}

	return u.Hostname()
}
