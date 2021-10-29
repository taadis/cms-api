package service

import "encoding/json"

type PostsData struct {
	Id      int64  `json:"id,omitempty" redis:"id"`
	Title   string `json:"title,omitempty" redis:"title"`
	Content string `json:"content,omitempty" redis:"content"`
	Author  string `json:"author,omitempty" redis:"author"`
	Time    int64  `json:"time,omitempty" redis:"time"`
}

func (d *PostsData) String() string {
	bs, _ := json.Marshal(d)
	return string(bs)
}

func (d *PostsData) ToMap() map[string]interface{} {
	bs, _ := json.Marshal(d)
	m := make(map[string]interface{}, 0)
	json.Unmarshal(bs, &m)
	return m
}

type SaveParams struct {
	Id      int64
	Title   string
	Content string
	Author  string
}

type ListParams struct {
	PageIndex int64
	PageSize int64
}
