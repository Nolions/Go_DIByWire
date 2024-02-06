// wire.go
//go:build wireinject
// +build wireinject

package di

import "github.com/google/wire"

func InitializeBlogService() BlogService {
	panic(wire.Build(NewBlogService, NewDataBase))
}
