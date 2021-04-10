package main

var urls []Url = []Url{
	{Path: "/", Method: "GET", View: createIndexView()},
	{Path: "/index", Method: "GET", View: createIndexView()},
	{Path: "/calc", Method: "GET", View: createCalcView()},
	{Path: "/calc", Method: "POST", View: createPostIndexView()},
	{Path: "/map", Method: "GET", View: createMapView()},
}
