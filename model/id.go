package model

type ID struct {
	Type  string      `json:"@type"`
	Value interface{} `json:"@value"`
}
