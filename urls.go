package main

var urls []Url = []Url{
	{Path: "/index", Method: "GET", View: createIndexView()},
	{Path: "/index", Method: "POST", View: createPostIndexView()},
}
