package model

import (
	"fmt"
	"net/url"
	"strings"
	"time"
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
	Time        int64  `json:"time"`
	Dead        bool   `json:"dead"`
	Kids        []int  `json:"kids"`
	Descendants int    `json:"descendants"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Rank        int
}

func (s *Story) Site() string {

	if s.Url == "" {
		return ""
	}

	u, err := url.Parse(s.Url)
	if err != nil {
		fmt.Println("ERR", "Cannot parse URL:", s.Url)
	}
	return u.Hostname()
}

func (s *Story) Age() string {

	duration := time.Now().Sub(time.Unix(s.Time, 0)).String()
	if strings.Contains(duration, "h") {
		return format(duration, "h", " hours ago")
	} else if strings.Contains(duration, "m") {
		return format(duration, "m", " minutes ago")
	} else if strings.Contains(duration, "s") {
		return format(strings.Split(duration, ".")[0], "s", " seconds ago")
	}
	return ""
}

func format(duration string, unit string, replacement string) string {
	return strings.Split(duration, unit)[0] + replacement
}
