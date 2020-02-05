package utils

import (
  "encoding/json"
  "net/http"
)

func JsonResponse(w http.ResponseWriter, v interface{}) {
  b, err := json.MarshalIndent(v, "", "    ")
  if err != nil {
    _, _ = w.Write([]byte("failed to encode json"))
  } else {
    _, _ = w.Write(b)
  }
}

func JsonError(message string, w http.ResponseWriter, code int) {
  w.Header().Set("content-type", "application/json")
  w.WriteHeader(code)
  JsonResponse(w, map[string]string{"error": message})
}

func JsonSuccess(v interface{}, w http.ResponseWriter) {
  w.Header().Set("content-type", "application/json")
  JsonResponse(w, map[string]interface{}{"data": v})
}
