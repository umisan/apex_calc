package main

import (
	"net/http"
)

var classList []int = []int{0, 1, 2, 3, 4}
var classMap map[int]string = map[int]string{
	0: "ブロンズ",
	1: "シルバー",
	2: "ゴールド",
	3: "プラチナ",
	4: "ダイヤ",
	5: "マスター",
}

type Result struct {
	Class string
	Cands CandList
}

type ResultList = []Result

type TemplateInput struct {
	Point int
	Table ResultList
}

type PostIndexView struct {
	defaultPoint int
	templatePath string
}

func createPostIndexView() PostIndexView {
	return PostIndexView{
		defaultPoint: 100,
		templatePath: "html/result.html",
	}
}

func (v PostIndexView) Validate(r *http.Request) ApplicationError {
	return nil
}

func (v PostIndexView) WriteResponse(w http.ResponseWriter, r *http.Request) {
	needPoint, err := v.readPointFromRequest(r)
	if err != nil {
		handleError(w, err)
	}
	classMapResult := v.calcClassMapResult(needPoint)
	indexTemplate, err := parseHtmlTemplate(v.templatePath, components)
	if err != nil {
		handleError(w, err)
	}
	indexTemplate.Execute(w, TemplateInput{Point: needPoint, Table: classMapResult})
}

func (v PostIndexView) readPointFromRequest(r *http.Request) (int, ApplicationError) {
	err := r.ParseForm()
	if err != nil {
		return 0, InternalError{}
	}
	point := getParameterOrDefaultInt(r.PostForm, "point", v.defaultPoint)
	return point, nil
}

func (v PostIndexView) calcClassMapResult(needPoint int) ResultList {
	results := ResultList{}
	for _, class := range classList {
		result := reverseCalculation(needPoint, class)
		results = append(results, Result{Class: classMap[class], Cands: result})
	}
	return results
}
