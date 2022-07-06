package api

import (
	"encoding/json"
	"example/internal/model"
	"fmt"
	"io/ioutil"
	"net/http"
)

var storiesNum int = 30

func GetTopStoriesIds() ([]int, error) {

	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		return nil, fmt.Errorf("HTTP call failed due to: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var ids []int
	if err := json.Unmarshal(body, &ids); err != nil {
		return nil, fmt.Errorf("Can not deserialize JSON")
	}
	//fmt.Println(ids)
	return ids, nil
}

func NextPage(page int) ([]model.Story, error) {

	stories := make([]model.Story, 0)

	ids, _ := GetTopStoriesIds()

	first := page
	if page != 0 {
		first = page*storiesNum - storiesNum
	}

	requested := ids[first : first+storiesNum]

	// can be parallelized
	for _, id := range requested {

		resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" + fmt.Sprint(id) + ".json")
		if err != nil {
			return nil, fmt.Errorf("HTTP call failed due to: %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		var result model.Story
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		stories = append(stories, result)
	}
	//fmt.Println(stories)
	return stories, nil
}
