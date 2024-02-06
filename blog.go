package di

type BlogPost string

type Database struct {
}

func NewDataBase() Database {
	return Database{}
}

func (db *Database) GetBlog() BlogPost {
	return BlogPost("this first blog post")
}
