package domain

import "errors"

var (
	ErrWrongLength        = errors.New("word length is not string value")
	ErrGeneratingNewWord  = errors.New("cannot generate new word")
	ErrRapidAPIRandomWord = errors.New("cannot generate new word")
	ErrWordDoesntExist    = errors.New("provided word doesnt exist")
)
