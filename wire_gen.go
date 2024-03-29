// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

// Injectors from write.go:

func InitializeBlogService() BlogService {
	database := NewDataBase()
	blogService := NewBlogService(database)
	return blogService
}
