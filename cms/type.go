package main

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type R struct {
	Data string `json:"data"`
	Err  string `json:"err"`
}

type Post struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Prev       string `json:"prev"`
	Next       string `json:"next"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Author     string `json:"author"`
	Tags       string `json:"tags"`
	DisplayId  string `json:"display_id"`
}

type Label struct {
	ID          int64  `json:"id"`
	NodeId      string `json:"node_id"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
	Description string `json:"description"`
}
