package main

var components []string = []string{
	"html/_google.html",
	"html/_meta_link.html",
	"html/_title.html",
	"html/_gundetail.html",
	"html/_title_en.html",
}

var urls []Url = []Url{
	{Path: "/", Method: "GET", View: createIndexView()},
	//日本語
	{Path: "/index", Method: "GET", View: createIndexView()},
	{Path: "/calc", Method: "GET", View: createCalcView()},
	{Path: "/calc", Method: "POST", View: createPostIndexView()},
	{Path: "/map", Method: "GET", View: createMapView()},
	{Path: "/guns", Method: "GET", View: createGunsView()},
	//英語
	{Path: "/index_en", Method: "GET", View: createIndexViewEn()},
	{Path: "/calc_en", Method: "GET", View: createCalcViewEn()},
	{Path: "/calc_en", Method: "POST", View: createPostIndexViewEn()},
	{Path: "/map_en", Method: "GET", View: createMapViewEn()},
	{Path: "/guns_en", Method: "GET", View: createGunsView()},
}
