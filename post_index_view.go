package main

import (
	"net/http"
)

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
	classList    []int
	classMap     map[int]string
}

func createPostIndexView() PostIndexView {
	return PostIndexView{
		defaultPoint: 100,
		templatePath: "html/result.html",
		classList:    []int{0, 1, 2, 3, 4},
		classMap: map[int]string{
			0: "ブロンズ",
			1: "シルバー",
			2: "ゴールド",
			3: "プラチナ",
			4: "ダイヤ",
			5: "マスター",
		},
	}
}

func createPostIndexViewEn() PostIndexView {
	return PostIndexView{
		defaultPoint: 100,
		templatePath: "html/result_en.html",
		classList:    []int{0, 1, 2, 3, 4},
		classMap: map[int]string{
			0: "bronze",
			1: "Silver",
			2: "Gold",
			3: "Platinum",
			4: "Diamond",
			5: "Master",
		},
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
	for _, class := range v.classList {
		result := reverseCalculation(needPoint, class)
		results = append(results, Result{Class: v.classMap[class], Cands: result})
	}
	return results
}
