package internal

import (
	"example/internal/api"
	"example/internal/model"
)

func NextPage() ([]model.Story, error) {

	return api.NextPage(31993429, 20)
}

func Latest() ([]int, error) {

	return api.GetTopStoriesIds()
}
