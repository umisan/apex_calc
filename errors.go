package main

import "net/http"

type ApplicationError interface {
	Error() string
	GetCode() int
}

// NetworkError
type NetworkError struct{}

func (e NetworkError) Error() string {
	return "NetworkError"
}

func (e NetworkError) GetCode() int {
	return http.StatusInternalServerError
}

// InternalError
type InternalError struct{}

func (e InternalError) Error() string {
	return "InternalError"
}

func (e InternalError) GetCode() int {
	return http.StatusInternalServerError
}

// HtmlParseError
type HtmlParseError struct {
}

func (e HtmlParseError) Error() string {
	return "HtmlParseError"
}

func (e HtmlParseError) GetCode() int {
	return http.StatusInternalServerError
}
