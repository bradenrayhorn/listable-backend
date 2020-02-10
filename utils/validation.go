package utils

import (
	"encoding/json"
	"fmt"
	"github.com/bradenrayhorn/listable-backend/db"
	"github.com/iancoleman/strcase"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func validateRequired(value string, field string, ruleValue string) ApiError {
	if len(value) > 0 {
		return ApiError{Code: 200}
	} else {
		return ApiError{
			Code:   422,
			Reason: field + " is required.",
		}
	}
}

func validateUnique(value string, field string, ruleValue string) ApiError {
	x := strings.Split(ruleValue, ",")
	count, err := db.GetDB().GetModelCount(x[0], x[1], value)
	if count > 0 {
		return ApiError{
			Code:   422,
			Reason: field + " must be unique.",
		}
	} else if err != nil {
		return ApiError{
			Code:  500,
			error: err,
		}
	} else {
		return ApiError{Code: 200}
	}
}

func validateInteger(value string, field string, ruleValue string) ApiError {
	if _, err := strconv.Atoi(value); err == nil {
		return ApiError{Code: 200}
	} else {
		return ApiError{
			Code:   422,
			Reason: field + " must be an integer.",
		}
	}
}

func validateBoolean(value string, field string, ruleValue string) ApiError {
	if _, err := strconv.ParseBool(value); err == nil {
		return ApiError{Code: 200}
	} else {
		return ApiError{
			Code:   422,
			Reason: field + " must be a boolean.",
		}
	}
}

var validationRules = map[string]interface{}{
	"required": validateRequired,
	"unique":   validateUnique,
	"integer":  validateInteger,
	"boolean":  validateBoolean,
}

func ValidateRequest(rules interface{}, r *http.Request, w http.ResponseWriter) (interface{}, *ApiError) {
	t := reflect.TypeOf(rules)
	err := r.ParseForm()
	if err != nil {
		// TODO better error handling
		panic("Failed to parse form")
	}

	decoder := json.NewDecoder(r.Body)
	var body map[string]interface{}
	if decoder.More() {
		err = decoder.Decode(&body)
		if err != nil {
			return nil, &ApiError{http.StatusBadRequest, "invalid json body", nil}
		}
	}

	resultStruct := reflect.New(t)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonName := strcase.ToSnake(field.Name)
		fieldRules := strings.Split(field.Tag.Get("validate"), ";")
		for _, rule := range fieldRules {
			rule = strings.TrimSpace(rule)
			ruleParts := strings.Split(rule, ":")
			ruleValue := ""
			// TODO need to validate rules better
			if len(ruleParts) > 1 {
				rule = ruleParts[0]
				ruleValue = ruleParts[1]
			} else {
				ruleValue = ruleParts[0]
			}
			validator, found := validationRules[rule]
			if found {
				var bodyValue string // TODO allow body value to be any type
				if body[jsonName] != nil {
					bodyValue = fmt.Sprintf("%v", body[jsonName])
				} else {
					bodyValue = r.Form.Get(jsonName)
				}
				apiError := validator.(func(string, string, string) ApiError)(bodyValue, jsonName, ruleValue)
				if apiError.error != nil {
					return nil, &apiError
				}
				if apiError.Code != 200 {
					return nil, &apiError
				}
				switch resultStruct.Elem().Field(i).Type().String() {
				case "int":
					value, _ := strconv.Atoi(bodyValue)
					resultStruct.Elem().Field(i).SetInt(int64(value))
				case "bool":
					value, _ := strconv.ParseBool(bodyValue)
					resultStruct.Elem().Field(i).SetBool(value)
				default:
					resultStruct.Elem().Field(i).SetString(bodyValue)
				}
			}
		}
	}

	return resultStruct.Interface(), nil
}
