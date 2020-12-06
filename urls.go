package main

var urls []Url = []Url{
	{Path: "/", Method: "GET", View: createIndexView()},
	{Path: "/index", Method: "GET", View: createIndexView()},
	{Path: "/index", Method: "POST", View: createPostIndexView()},
}
