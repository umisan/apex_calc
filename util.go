package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func handleError(w http.ResponseWriter, e ApplicationError) {
	log.Println(e)
	w.WriteHeader(e.GetCode())
	w.Write([]byte(e.Error()))
}

func writeBlob(w http.ResponseWriter, blob []byte) ApplicationError {
	_, err := w.Write(blob)
	if err != nil {
		return NetworkError{}
	}
	return nil
}

func strToFloat32(s string) (float32, error) {
	tmp, err := strconv.ParseFloat(s, 32)
	return float32(tmp), err
}

func getParameterOrDefaultFloat(parameters url.Values, key string, defaultValue float32) float32 {
	candidates, ok := parameters[key]
	if !ok {
		return defaultValue
	}
	parameter, err := strToFloat32(candidates[0])
	if err != nil {
		return defaultValue
	}
	return parameter
}

func getParameterOrDefaultInt(parameters url.Values, key string, defaultValue int) int {
	candidates, ok := parameters[key]
	if !ok {
		return defaultValue
	}
	parameter, err := strconv.Atoi(candidates[0])
	if err != nil {
		return defaultValue
	}
	return parameter
}

func getParameterOrDefaultString(parameters url.Values, key, defaultValue string) string {
	candidates, ok := parameters[key]
	if !ok {
		return defaultValue
	}
	return candidates[0]
}

func parseHtmlTemplate(filename ...string) (*template.Template, ApplicationError) {
	indexTemplate, err := template.ParseFiles(filename...)
	if err != nil {
		return nil, HtmlParseError{}
	}
	return indexTemplate, nil
}
