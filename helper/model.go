package helper

import (
	"irc/golang-restful-api-sample/model/api"
	"irc/golang-restful-api-sample/model/domain"
)

func ToCategoryResponse(category domain.Category) api.CategoryResponse {
	return api.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []api.CategoryResponse {
	var categoryResponses []api.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
