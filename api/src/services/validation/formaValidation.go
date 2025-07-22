package validation

import (
	"errors"
	"fmt"
	"reflect"
)

type User struct {
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Age   uint   `validate:"required"`
}

type FormValidation struct {
	concrete  any
	error_bag map[string]string
}

func NewFormValidation[T interface{}](concrete T) *FormValidation {
	return &FormValidation{concrete: concrete, error_bag: map[string]string{}}
}
func (f FormValidation) HasFailed() bool {
	return len(f.error_bag) >= 1
}
func (f FormValidation) GetErrors() map[string]string {
	return f.error_bag
}

func Validate[T interface{}](concrete T, f *FormValidation) error {
	checkStruct := reflect.TypeOf(concrete)
	if checkStruct.Kind().String() != "struct" {
		return errors.New("struct type only is validated")
	}
	var type_assertion interface{} = f.concrete
	concreate_struct := type_assertion.(T)
	v := reflect.ValueOf(concreate_struct)
	t := reflect.TypeOf(concreate_struct)
	for i := range v.NumField() {
		vField := v.Field(i)
		tField := t.Field(i)
		hasRequired := tField.Tag.Get("validate")
		property_type := tField.Type.Name()
		if property_type == "string" && hasRequired == "required" {
			if vField.String() == "" {
				f.error_bag[tField.Name] = fmt.Sprintf("%s is required", tField.Name)
			}
		}
		if property_type == "int" && hasRequired == "required" {
			if vField.Int() == 0 {
				f.error_bag[tField.Name] = fmt.Sprintf("%s is required", tField.Name)
			}
		}
		if property_type == "uint" && hasRequired == "required" {
			if vField.Uint() == 0 {
				f.error_bag[tField.Name] = fmt.Sprintf("%s is required", tField.Name)
			}
		}
	}
	return nil
}
