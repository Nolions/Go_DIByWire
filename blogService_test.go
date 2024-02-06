package di

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	db := NewDataBase()
	serv := NewBlogService(db)

	expected := BlogPost("this first blog post")
	actual := serv.Print()

	assert.Equal(t, expected, actual)
}
