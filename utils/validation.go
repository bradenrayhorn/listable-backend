package utils

import (
  "github.com/bradenrayhorn/listable-backend/db"
  "github.com/iancoleman/strcase"
  "net/http"
  "reflect"
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

var validationRules = map[string]interface{}{
  "required": validateRequired,
  "unique":   validateUnique,
}

func ValidateRequest(rules interface{}, r *http.Request, w http.ResponseWriter) (interface{}, *ApiError) {
  t := reflect.TypeOf(rules)
  err := r.ParseForm()
  if err != nil {
    // TODO better error handling and parse JSON bodies
    panic("Failed to parse form")
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
        bodyValue := r.Form.Get(jsonName)
        apiError := validator.(func(string, string, string) ApiError)(bodyValue, field.Name, ruleValue)
        if apiError.error != nil {
          return nil, &apiError
        }
        if apiError.Code != 200 {
          return nil, &apiError
        }
        resultStruct.Elem().Field(i).SetString(bodyValue)
      }
    }
  }
  
  return resultStruct.Interface(), nil
}
