package internal

import (
	"example/internal/api"
	"example/internal/model"
	"net/http"
	"sync"
	"text/template"
)

var templates = template.Must(template.ParseFiles("../../web/index.html"))

var storiesNum int = 30

func nextPage(page int) ([]model.Story, error) {

	ids, _ := api.TopStoriesIds()

	first := page
	if page != 0 {
		first = page*storiesNum - storiesNum
	}

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

	stories := order(subset, results)

	//fmt.Println("DEBUG", "done consuming channel")
	return stories, nil
}

func order(ids []int, stories []model.Story) []model.Story {

	ordered := make([]model.Story, 0)
	for id := range ids {

		for _, v := range stories {
			if v.Id == ids[id] {
				ordered = append(ordered, v)
			}
		}
	}
	return ordered
}

func Home(w http.ResponseWriter, r *http.Request) {

	stories, err := nextPage(0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	p := &model.Page{Stories: stories}
	templErr := templates.ExecuteTemplate(w, "index.html", p)
	if templErr != nil {
		http.Error(w, templErr.Error(), http.StatusInternalServerError)
	}
}
