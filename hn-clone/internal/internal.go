package internal

import (
	"example/internal/api"
	"example/internal/model"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

var templates = template.Must(template.ParseFiles("../../web/index.html"))

var storiesNum int = 30

func nextPage(page int) ([]model.Story, error) {

	ids, _ := api.TopStoriesIds()

	first := page*storiesNum - storiesNum

	subset := ids[first : first+storiesNum]

	ch := make(chan model.Story)
	var wg sync.WaitGroup
	for _, id := range subset {

		wg.Add(1)
		go api.Story(id, ch, &wg)
	}
	//fmt.Println("DEBUG", "done fetching stories")

	// close the channel in the background
	go func() {
		wg.Wait()
		close(ch)
	}()

	results := make([]model.Story, 0)
	for story := range ch {
		results = append(results, story)
	}

	stories := order(subset, results, page)

	//fmt.Println("DEBUG", "done consuming channel")
	return stories, nil
}

func order(ids []int, stories []model.Story, page int) []model.Story {

	ordered := make([]model.Story, 0)
	for id := range ids {

		for _, v := range stories {
			if v.Id == ids[id] {

				v.Rank = id + (30 * (page - 1)) + 1
				ordered = append(ordered, v)
			}
		}
	}
	return ordered
}

func Home(w http.ResponseWriter, r *http.Request) {

	stories, err := nextPage(1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	p := &model.Page{Stories: stories, NextPage: 2}
	templErr := templates.ExecuteTemplate(w, "index.html", p)
	if templErr != nil {
		http.Error(w, templErr.Error(), http.StatusInternalServerError)
	}
}

func News(w http.ResponseWriter, r *http.Request) {

	num := r.URL.Query().Get("p")
	page, err := strconv.Atoi(num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	stories, err := nextPage(page)
	p := &model.Page{Stories: stories, NextPage: page + 1}
	templErr := templates.ExecuteTemplate(w, "index.html", p)
	if templErr != nil {
		http.Error(w, templErr.Error(), http.StatusInternalServerError)
	}
}
