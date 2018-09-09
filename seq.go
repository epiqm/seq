// Package for encoding/decoding data sequences.
//
// A functional library for performing common operations
// on data. Rand, MD5 hash, Base64 encode/decode,
// AES(CBC) encryption/decryption.
package seq

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	mr "math/rand"
	"os"
	"time"
)

// Generates random number between min and max.
func Rand(min, max int) int {
	mr.Seed(time.Now().Unix())
	return mr.Intn(max-min) + min
}

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
func Encode(text string, key string) (s string, err error) {
	keyEnc := []byte(key)
	plaintext := []byte(text)
	if len(key) == 0 {
		s = base64.StdEncoding.EncodeToString(plaintext)
	} else {
		if len(plaintext)%aes.BlockSize != 0 {
			err := errors.New("decode: cipher text is too short")
			return "", err
		}

		block, err := aes.NewCipher(keyEnc)
		if err != nil {
			return "", err
		}

		ciphertext := make([]byte, aes.BlockSize+len(plaintext))
		iv := ciphertext[:aes.BlockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return "", err
		}

		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

		s = fmt.Sprintf("%x", ciphertext)
	}
	return s, nil
}

// Decodes/decrypts Base64 string.
func Decode(text string, key string) (s string, err error) {
	if len(key) == 0 {
		data, err := base64.StdEncoding.DecodeString(text)
		if err != nil {
			return "", err
		}
		s = string(data)
		return s, nil
	} else {
		keyEnc := []byte(key)
		ciphertext, _ := hex.DecodeString(text)

		block, err := aes.NewCipher(keyEnc)
		if err != nil {
			return "", err
		}

		if len(ciphertext) < aes.BlockSize {
			err := errors.New("decode: cipher text is too short")
			return "", err
		}

		iv := ciphertext[:aes.BlockSize]
		ciphertext = ciphertext[aes.BlockSize:]

		if len(ciphertext)%aes.BlockSize != 0 {
			err := errors.New("decode: ciphertext is not a multiple of the block size")
			return "", err
		}

		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(ciphertext, ciphertext)

		s = fmt.Sprintf("%s", ciphertext)
		return s, nil
	}
}

// Reads file contents to string.
func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Writes text string to file.
func WriteFile(path string, text string) error {
	err := ioutil.WriteFile(path, []byte(text), 0644)
	if err != nil {
		return err
	}
	return nil
}

// Creates a file.
func CreateFile(path string) error {
	f, err := os.Create(path)
	f.Close()
	if err != nil {
		return err
	}
	return nil
}

// Deletes file.
func RmFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// Moves file to the new path.
func MoveFile(oldpath string, newpath string) error {
	err := os.Rename(oldpath, newpath)
	if err != nil {
		return err
	}
	return nil
}

// Copies file to the new path.
func CopyFile(path string, newpath string) error {
	in, err := os.Open(path)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(newpath)
	if err != nil {
		return err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	err = out.Sync()
	return nil
}

// Encodes/encrypts contents and rewrites the file.
func EncodeFile(path string, key string) error {
	text, err := ReadFile(path)
	if err != nil {
		return err
	}
	enc, err := Encode(text, key)
	if err != nil {
		return err
	}
	err = WriteFile(path, enc)
	if err != nil {
		return err
	}
	return nil
}

// Decodes/decrypts contents and rewrites the file.
func DecodeFile(path string, key string) error {
	text, err := ReadFile(path)
	if err != nil {
		return err
	}
	dec, err := Decode(text, key)
	if err != nil {
		return err
	}
	err = WriteFile(path, dec)
	if err != nil {
		return err
	}
	return nil
}

// Creates a directory.
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// Lists a directory.
func Ls(path string) (fm []string, err error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		fm = append(fm, f.Name())
	}
	return
}

// Removes a directory.
func RmDir(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}
