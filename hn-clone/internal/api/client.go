package api

import (
	"encoding/json"
	"example/internal/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var storiesNum int = 30

func TopStoriesIds() ([]int, error) {

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
	return ids, nil
}

func Story(id int, ch chan model.Story, wg *sync.WaitGroup) {

	//fmt.Println("DEBUG", "fectching story with id:", id)
	defer wg.Done()

	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" + fmt.Sprint(id) + ".json")
	if err != nil {
		//TODO bubble up err
		fmt.Println(fmt.Errorf("HTTP call failed due to: %v", err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result model.Story
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	ch <- result
}
