// Package for encoding/decoding data sequences.
package seq

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
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

// Returns Base64 encoded string from text.
// Uses AES(CBC) encryption if key is present.
func Encode(text string, key string) (s string) {
	keyEnc := []byte(key)
	plaintext := []byte(text)

	if len(key) == 0 {
		s = base64.StdEncoding.EncodeToString(plaintext)
		return
	} else {
		if len(plaintext)%aes.BlockSize != 0 {
			return
		}

		block, err := aes.NewCipher(keyEnc)
		if err != nil {
			return
		}

		ciphertext := make([]byte, aes.BlockSize+len(plaintext))
		iv := ciphertext[:aes.BlockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return
		}

		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

		s = fmt.Sprintf("%x", ciphertext)
		return
	}
	return
}

// Decodes/decrypts Base64 string.
func Decode(text string, key string) (s string) {
	if len(key) == 0 {
		data, err := base64.StdEncoding.DecodeString(text)
		if err != nil {
			return "failed to decode"
		}
		s = string(data)
		return
	} else {
		keyEnc := []byte(key)
		ciphertext, _ := hex.DecodeString(text)

		block, err := aes.NewCipher(keyEnc)
		if err != nil {
			return
		}

		if len(ciphertext) < aes.BlockSize {
			return "cipher text is too short"
		}

		iv := ciphertext[:aes.BlockSize]
		ciphertext = ciphertext[aes.BlockSize:]

		if len(ciphertext)%aes.BlockSize != 0 {
			return "ciphertext is not a multiple of the block size"
		}

		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(ciphertext, ciphertext)

		s = fmt.Sprintf("%s", ciphertext)
		return
	}
}
