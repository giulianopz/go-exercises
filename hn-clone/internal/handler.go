package internal

import (
	"example/internal/api"
	"example/internal/mapper"
	"example/internal/model"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

var templates = template.Must(template.ParseFiles("../../web/index.html"))

var storiesNum int = 30

func nextTopStoriesPage(page int) ([]model.Story, error) {

	return nextPage(page, api.TopStoriesURL)
}

func nextNewestStoriesPage(page int) ([]model.Story, error) {
	return nextPage(page, api.NewStoriesURL)
}

func nextAskStoriesPage(page int) ([]model.Story, error) {
	return nextPage(page, api.AskStoriesURL)
}

func nextPage(page int, url string) ([]model.Story, error) {

	ids, _ := api.StoriesIds(url)

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

func Top(w http.ResponseWriter, r *http.Request) {

	stories, err := nextTopStoriesPage(1)
	if err != nil {
		mapper.ErrToISE(w, err)
	}

	p := &model.Page{Stories: stories, NextPage: "/news?p=2"}
	templErr := templates.ExecuteTemplate(w, "index.html", p)
	if templErr != nil {
		mapper.StrToISE(w, templErr.Error())
	}
}

func pageParam(w http.ResponseWriter, r *http.Request) int {
	num := r.URL.Query().Get("p")
	page, err := strconv.Atoi(num)
	if err != nil {
		mapper.ErrToISE(w, err)
	}
	return page
}

func News(w http.ResponseWriter, r *http.Request) {

	page := pageParam(w, r)

	stories, err := nextTopStoriesPage(page)
	if err != nil {
		mapper.ErrToISE(w, err)
	}
	next := strconv.Itoa(page + 1)
	p := &model.Page{Stories: stories, NextPage: "/news?p=" + next}
	templErr := templates.ExecuteTemplate(w, "index.html", p)
	if templErr != nil {
		mapper.StrToISE(w, templErr.Error())
	}
}

func Newest(w http.ResponseWriter, r *http.Request) {

	page := pageParam(w, r)

	stories, err := nextNewestStoriesPage(page)
	if err != nil {
		mapper.ErrToISE(w, err)
	}

	next := strconv.Itoa(page + 1)
	p := &model.Page{Stories: stories, NextPage: "/newest?p=" + next}
	templErr := templates.ExecuteTemplate(w, "index.html", p)
	if templErr != nil {
		mapper.ErrToISE(w, templErr)
	}
}

func Ask(w http.ResponseWriter, r *http.Request) {

	page := pageParam(w, r)

	stories, err := nextAskStoriesPage(page)
	if err != nil {
		mapper.ErrToISE(w, err)
	}

	next := strconv.Itoa(page + 1)
	p := &model.Page{Stories: stories, NextPage: "/ask?p=" + next}
	templErr := templates.ExecuteTemplate(w, "index.html", p)
	if templErr != nil {
		mapper.StrToISE(w, templErr.Error())
	}
}
