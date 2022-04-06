package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
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

func (f *Form) IsValid() bool {
	return len(f.Errors) == 0
}

func (f *Form) IsNotValid() bool {
	return !f.IsValid()
}

func (f *Form) CheckEmail(field string){
	if !govalidator.IsEmail(f.Get(field)){
		f.Errors.Add(field, "Invalid email address")
	}
}

func (f *Form) CheckMinLengh(field string, length int){
	value := f.Get(field)
	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
	}
}

func (f *Form) CheckRequiredFields(fields ...string){
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}



