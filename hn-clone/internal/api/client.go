package api

import (
	"encoding/json"
	"example/internal/model"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTopStoriesIds() ([]int, error) {

	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		return nil, fmt.Errorf("HTTP call failed due to: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	var ids []int
	if err := json.Unmarshal(body, &ids); err != nil {
		return nil, fmt.Errorf("Can not deserialize JSON")
	}
	fmt.Println(ids)

	return ids, nil
}

func NextPage(id int, pages int) ([]model.Story, error) {

	stories := make([]model.Story, 0)

	ids, _ := GetTopStoriesIds()

	idx := 0

	for i, v := range ids {
		if id == v {
			idx = i
		}
	}

	requested := ids[idx : idx+pages]

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
	fmt.Println(stories)
	return stories, nil
}
