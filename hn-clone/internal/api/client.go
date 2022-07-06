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
	return ids, nil
}

func getStory(id int, ch chan model.Story, wg *sync.WaitGroup) {

	fmt.Println("DEBUG", "fectching story with id:", id)

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

func NextPage(page int) ([]model.Story, error) {

	ids, _ := GetTopStoriesIds()

	first := page
	if page != 0 {
		first = page*storiesNum - storiesNum
	}

	requested := ids[first : first+storiesNum]

	ch := make(chan model.Story)
	var wg sync.WaitGroup
	for _, id := range requested {

		wg.Add(1)
		go getStory(id, ch, &wg)
	}
	fmt.Println("DEBUG", "done fetching stories")

	// close the channel in the background
	go func() {
		wg.Wait()
		close(ch)
	}()

	stories := make([]model.Story, 0)
	for story := range ch {
		stories = append(stories, story)
	}
	fmt.Println("DEBUG", "done consuming channel")
	return stories, nil
}
