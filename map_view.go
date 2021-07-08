package main

import "net/http"

type MapView struct {
	templatePath string
}

func createMapView() MapView {
	return MapView{
		templatePath: "html/map.html",
	}
}

func createMapViewEn() MapView {
	return MapView{
		templatePath: "html/map_en.html",
	}
}

func (v MapView) Validate(r *http.Request) ApplicationError {
	return nil
}

func (v MapView) WriteResponse(w http.ResponseWriter, r *http.Request) {
	indexTemplate, err := parseHtmlTemplate(v.templatePath, components)
	if err != nil {
		handleError(w, err)
	}
	indexTemplate.Execute(w, nil)
}
