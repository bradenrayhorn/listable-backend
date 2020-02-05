package models

import "math/rand"

type ApiToken struct {
  UserID uint
  Token  string
  User   User
  ModelTimestamps
}

func (u ApiToken) TableName() string {
  return "api_tokens"
}

func (token *ApiToken) Generate() {
  var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
  
  b := make([]rune, 64)
  for i := range b {
    b[i] = letter[rand.Intn(len(letter))]
  }
  token.Token = string(b)
}
