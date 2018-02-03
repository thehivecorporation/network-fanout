package main

import "net/http"

type ResponseOrError struct {
	err    error
	res    *http.Response
	target *shortURL
}

type SingleResponse struct {
	StatusCode int                    `json:",omitempty"`
	Status     string                 `json:",omitempty"`
	Response   map[string]interface{} `json:",omitempty"`
	Error      string                 `json:",omitempty"`
	Target     *shortURL              `json:",omitempty"`
}

type CompoundResponse struct {
	Responses []SingleResponse `json:",omitempty"`
	Status    string           `json:",omitempty"`
}

type shortURL struct {
	Scheme string	`json:",omitempty"`
	Host   string	`json:",omitempty"`
	Path   string	`json:",omitempty"`
}
