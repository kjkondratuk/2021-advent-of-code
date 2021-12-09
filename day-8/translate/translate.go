package translate

import (
	"errors"
	"fmt"
)

var (
	noTranslation  = "no translation for: "
	ErrInvalidType = func(t interface{}) error {
		return errors.New(fmt.Sprintf("invalid type: %T", t))
	}
	ErrNoTranslationString = func(t string) error {
		return errors.New(fmt.Sprintf(noTranslation+"%s", t))
	}
	ErrNoTranslationInt = func(t int) error {
		return errors.New(fmt.Sprintf(noTranslation+"%d", t))
	}
)

type Translator interface {
	Translate(input interface{}) (output interface{}, err error)
}

type ToInt map[string]int

// Translate : translates a string to an int based on the translation mapping.
func (t ToInt) Translate(input interface{}) (output interface{}, err error) {
	if iv, ok := input.(string); ok {
		if mv, ok := t[iv]; ok {
			return mv, nil
		} else {
			return nil, ErrNoTranslationString(iv)
		}
	} else {
		return nil, ErrInvalidType(input)
	}
}

type ToString map[int]string

// Translate : translates an int to a string based on the translation mapping.
func (t ToString) Translate(input interface{}) (output interface{}, err error) {
	if iv, ok := input.(int); ok {
		if mv, ok := t[iv]; ok {
			return mv, nil
		} else {
			return nil, ErrNoTranslationInt(iv)
		}
	} else {
		return nil, ErrInvalidType(input)
	}
}

type StringToString map[string]string

// Translate : translates a string to another string based on the translation mapping.
func (t StringToString) Translate(input interface{}) (output interface{}, err error) {
	if iv, ok := input.(string); ok {
		if mv, ok := t[iv]; ok {
			return mv, nil
		} else {
			return nil, ErrNoTranslationString(iv)
		}
	} else {
		return nil, ErrInvalidType(input)
	}
}
