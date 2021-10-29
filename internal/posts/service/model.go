package service

type PostsData struct {
	Id      int64
	Title   string
	Content string
	Author  string
	Time    int64
}

type SaveParams struct {
	Id      int64
	Title   string
	Content string
	Author  string
}
