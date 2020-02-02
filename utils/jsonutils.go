package utils

import (
  "encoding/json"
  "net/http"
)

func JsonResponse(w http.ResponseWriter, v interface{}) {
  w.Header().Set("Content-Type", "application/json")
  b, err := json.MarshalIndent(v, "", "    ")
  if err != nil {
    _, _ = w.Write([]byte("failed to encode json"))
  } else {
    _, _ = w.Write(b)
  }
}
