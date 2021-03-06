package main

import (
	"net/http"
)

type CalcView struct {
	defaultPoint int
	templatePath string
}

func createCalcView() IndexView {
	return IndexView{
		defaultPoint: 100,
		templatePath: "html/calc.html",
	}
}

func createCalcViewEn() IndexView {
	return IndexView{
		defaultPoint: 100,
		templatePath: "html/calc_en.html",
	}
}

func (v CalcView) Validate(r *http.Request) ApplicationError {
	return nil
}

func (v CalcView) WriteResponse(w http.ResponseWriter, r *http.Request) {
	indexTemplate, err := parseHtmlTemplate(v.templatePath, components)
	if err != nil {
		handleError(w, err)
	}
	indexTemplate.Execute(w, nil)
}
