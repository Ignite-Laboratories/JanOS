package main

import (
	"encoding/json"
	"os"
)

type Lexicon struct {
	Lexeme           string   `json:"lexeme"`
	Type             string   `json:"type"`
	Docs             []string `json:"docs"`
	Related          []string `json:"related"`
	Types            []Lexeme `json:"types"`
	AdditionalFields []string `json:"additionalFields,omitempty"`
}

type Lexeme struct {
	Name       string            `json:"name"`
	Docs       []string          `json:"docs"`
	Interface  bool              `json:"interface,omitempty"`
	Base       string            `json:"base,omitempty"`
	NameSet    string            `json:"nameSet,omitempty"`
	Set        map[string]string `json:"set,omitempty"`
	Exhaustive bool              `json:"exhaustive,omitempty"`
	Alias      string            `json:"alias,omitempty"`
}

func ParseJSON(path string) (lexicon *Lexicon, err error) {
	var data []byte
	data, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &lexicon); err != nil {
		return nil, err
	}
	return lexicon, nil
}
