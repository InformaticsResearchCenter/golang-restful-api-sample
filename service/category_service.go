package service

import (
	"context"
	"irc/golang-restful-api-sample/model/api"
)

type CategoryService interface {
	Create(ctx context.Context, request api.CategoryCreateRequest) api.CategoryResponse
	Update(ctx context.Context, request api.CategoryUpdateRequest) api.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) api.CategoryResponse
	FindAll(ctx context.Context) []api.CategoryResponse
}
