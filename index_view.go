package main

import (
	"net/http"
)

type IndexView struct {
	defaultPoint int
	templatePath string
}

func createIndexView() IndexView {
	return IndexView{
		defaultPoint: 100,
		templatePath: "html/index.html",
	}
}

func createIndexViewEn() IndexView {
	return IndexView{
		defaultPoint: 100,
		templatePath: "html/index_en.html",
	}
}

func (v IndexView) Validate(r *http.Request) ApplicationError {
	return nil
}

func (v IndexView) WriteResponse(w http.ResponseWriter, r *http.Request) {
	indexTemplate, err := parseHtmlTemplate(v.templatePath, components)
	if err != nil {
		handleError(w, err)
	}
	indexTemplate.Execute(w, nil)
}
