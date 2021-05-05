package main

import "net/http"

type GunsView struct {
	info         GunsViewInfo
	templatePath string
}

type GunsViewInfo struct {
	LightGunDetails   []GunDetail
	HeavyGunDetails   []GunDetail
	EnergyGunDetails  []GunDetail
	SniperGunDetails  []GunDetail
	ShotGunDetails    []GunDetail
	SpecialGunDetails []GunDetail
}

type GunDetail struct {
	Id   string
	Name string
	Head string
	Body string
	Foot string
}

func createGunsView() GunsView {
	light := []GunDetail{
		{Id: "g7", Name: "G7", Head: "60", Body: "34", Foot: "26"},
		{Id: "alternator", Name: "オルタネーター", Head: "24", Body: "16", Foot: "13"},
		{Id: "r99", Name: "R99", Head: "17", Body: "11", Foot: "9"},
		{Id: "r301", Name: "R301カービン", Head: "28", Body: "14", Foot: "11"},
		{Id: "p2020", Name: "P2020", Head: "23", Body: "15", Foot: "14"},
		{Id: "re45", Name: "RE45", Head: "18", Body: "12", Foot: "11"},
	}
	heavy := []GunDetail{
		{Id: "flatline", Name: "フラットライン", Head: "38", Body: "19", Foot: "14"},
		{Id: "hemlok", Name: "ヘムロック", Head: "35", Body: "20", Foot: "15"},
		{Id: "spitfire", Name: "スピットファイア", Head: "38", Body: "19", Foot: "14"},
		{Id: "wingman", Name: "ウイングマン", Head: "90", Body: "45", Foot: "41"},
		{Id: "repeater", Name: "30-30リピーター", Head: "100(74)", Body: "57(42)", Foot: "43(32)"},
	}
	energy := []GunDetail{
		{Id: "lstar", Name: "Lスター", Head: "36", Body: "18", Foot: "14"},
		{Id: "havoc", Name: "ハボック", Head: "36", Body: "18", Foot: "14"},
		{Id: "devotion", Name: "ディボーション", Head: "32", Body: "16", Foot: "12"},
		{Id: "volt", Name: "ボルト", Head: "23", Body: "15", Foot: "12"},
	}
	sniper := []GunDetail{
		{Id: "triple-take", Name: "トリプルテイク", Head: "138(46)", Body: "69(23)", Foot: "63(21)"},
		{Id: "longbow", Name: "ロングボウ", Head: "110", Body: "55", Foot: "44"},
		{Id: "charge-rifle", Name: "チャージライフル", Head: "116", Body: "90", Foot: "90"},
		{Id: "sentinel", Name: "センチネル", Head: "140", Body: "70", Foot: "63"},
	}
	shotgun := []GunDetail{
		{Id: "mastiff", Name: "マスティフ", Head: "112(14)", Body: "88(11)", Foot: "88(11)"},
		{Id: "eva8", Name: "EVA8", Head: "99(11)", Body: "63(7)", Foot: "63(7)"},
		{Id: "mozambique", Name: "モザンビーク", Head: "69(33)", Body: "45(15)", Foot: "45(15)"},
	}
	special := []GunDetail{
		{Id: "kraber", Name: "クレーバー", Head: "435", Body: "145", Foot: "116"},
		{Id: "prowler", Name: "プラウラー", Head: "23", Body: "15", Foot: "12"},
	}
	return GunsView{
		info: GunsViewInfo{
			LightGunDetails:   light,
			HeavyGunDetails:   heavy,
			EnergyGunDetails:  energy,
			SniperGunDetails:  sniper,
			ShotGunDetails:    shotgun,
			SpecialGunDetails: special,
		},
		templatePath: "html/guns.html",
	}
}

func (v GunsView) Validate(r *http.Request) ApplicationError {
	return nil
}

func (v GunsView) WriteResponse(w http.ResponseWriter, r *http.Request) {
	gunsTemplate, err := parseHtmlTemplate(v.templatePath, components)
	if err != nil {
		handleError(w, err)
	}
	gunsTemplate.Execute(w, v.info)
}
