package forms

import (
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		if strings.TrimSpace(f.Get(field)) == "" {
			f.Errors.Add(field, "This field cannot be empty")
		}
	}
}

func (f *Form) MaxLenght(field string, lenght int) {
	if utf8.RuneCountInString(f.Get(field)) > lenght {
		f.Errors.Add(field, "Field cannot surpass max lenght")
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
