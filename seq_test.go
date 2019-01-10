package seq_test

import (
	"fmt"
	"testing"

	"seq"
)

func TestRand(t *testing.T) {
	number := seq.Rand(1, 100)

	if number < 1 && number > 100 {
		t.Errorf("Rand failed, got: %d, want range 1-100", number)
	}
}

func ExampleRand(t *testing.T) {
	number := seq.Rand(1, 100)
	fmt.Println(number)
}

func ExampleHash() {
	hash := seq.Hash("some text")

	fmt.Println(hash)
	// Output:
	// 552e21cd4cd9918678e3c1a0df491bc3
}

func ExampleHashCut() {
	shortHash := seq.HashCut("some text", 3)

	fmt.Println(shortHash)
	// Output:
	// 552
}

func ExampleMarshal() {
	// prepare our object to perform marshal
	obj := map[string]interface{}{
		"username": "Max",
		"age":      27,
		"hometown": "Kiev",
	}

	// marshal object to string
	str := seq.Marshal(obj)

	fmt.Println(str)
	// Output:
	// {"age":27,"hometown":"Kiev","username":"Max"}
}

func ExampleUnmarshal() {
	// prepare author structure
	type Author struct {
		Username string `json:"username"`
		Hometown string `json:"hometown"`
		Age      int    `json:"age"`
	}

	// prepare json text string
	json := `{"age":27,"hometown":"Kiev","username":"Max"}`

	// create author variable
	var obj Author

	// unmarshal json text string to object
	seq.Unmarshal(json, &obj)

	fmt.Println(obj)
	// Output:
	// {Max Kiev 27}
}

func ExampleEncode() {
	key := ""
	text := "A message for encoding."

	s, err := seq.Encode(text, key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)
	// Output:
	// QSBtZXNzYWdlIGZvciBlbmNvZGluZy4=
}

func ExampleDecode() {
	key := ""
	text := "QSBtZXNzYWdlIGZvciBlbmNvZGluZy4="

	s, err := seq.Decode(text, key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)
	// Output:
	// A message for encoding.
}

func ExampleGetFileSize() {
	err := seq.WriteFile("testing.txt", "MAX")
	if err != nil {
		fmt.Println(err)
		return
	}

	fsize, err := seq.GetFileSize("testing.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fsize)
	// Output:
	// 3
}

func ExampleCreateFile() {
	err := seq.CreateFile("testing.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
	// Output:
	// OK
}

func ExampleWriteFile() {
	err := seq.WriteFile("testing.txt", "A testing text file.")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
	// Output:
	// OK
}

func ExampleReadFile() {
	text, err := seq.ReadFile("testing.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(text)
	// Output:
	// A testing text file.
}

func ExampleCopyFile() {
	err := seq.CopyFile("testing.txt", "testing.txt.copy")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
	// Output:
	// OK
}

func ExampleMoveFile() {
	err := seq.MoveFile("testing.txt.copy", "testing.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
	// Output:
	// OK
}

func ExampleEncodeFile() {
	err := seq.EncodeFile("testing.txt", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	text, err := seq.ReadFile("testing.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(text)
	// Output:
	// QSB0ZXN0aW5nIHRleHQgZmlsZS4=
}

func ExampleDecodeFile() {
	err := seq.DecodeFile("testing.txt", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	text, err := seq.ReadFile("testing.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(text)
	// Output:
	// A testing text file.
}

func ExampleRmFile() {
	err := seq.RmFile("testing.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
	// Output:
	// OK
}

func ExampleCreateDir() {
	err := seq.CreateDir("newdir")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = seq.WriteFile("newdir/testing.txt", "Another test file.")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
	// Output:
	// OK
}

func ExampleLs() {
	files, err := seq.Ls("newdir")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		fmt.Println(f)
	}

	// Output:
	// testing.txt
}

func ExampleRmDir() {
	err := seq.RmDir("newdir")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
	// Output:
	// OK
}
