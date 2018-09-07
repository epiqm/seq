package seq_test

import (
	"fmt"

	"github.com/epiqm/seq"
)

func ExampleHash() {
	hash := seq.Hash("some text")

	fmt.Println(hash)
	// Output:
	// 552e21cd4cd9918678e3c1a0df491bc3
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

	s := seq.Encode(text, key)

	fmt.Println(s)
	// Output:
	// QSBtZXNzYWdlIGZvciBlbmNvZGluZy4=
}

func ExampleDecode() {
	key := ""
	text := "QSBtZXNzYWdlIGZvciBlbmNvZGluZy4="

	s := seq.Decode(text, key)

	fmt.Println(s)
	// Output:
	// A message for encoding.
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
