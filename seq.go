// Package for encoding/decoding data sequences.
package seq

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

// Performs MD5 hashing operation on given text.
func Hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

// Outputs JSON string from structure.
func Marshal(m map[string]interface{}) (s string) {
	b, err := json.Marshal(m)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

// Unmarshals JSON string to structure.
func Unmarshal(text string, s interface{}) {
	err := json.Unmarshal([]byte(text), s)
	if err != nil {
		return
	}
	return
}
