package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type CustomBuffer struct {
	Description string
	Buffer      bytes.Buffer
}

func main() {
	b := new(CustomBuffer)
	b2 := CustomBuffer{}

	fmt.Println("new(CustomBuffer)", reflect.TypeOf(b))
	fmt.Println("CustomBuffer{}", reflect.TypeOf(b2))

	_, err := fmt.Fprintf(&b.Buffer, "sdfds fdssfsd sdds sddsdfdsdsfds")
	if err != nil {
		panic(err)
	}

	fmt.Println(b.Buffer.String())

	sb := strings.NewReader("teste")
	bugg := make([]byte, 20)
	n, err := sb.Read(bugg)
	if err != nil {
		panic(err)
	}
	fmt.Println("bugg: bytes read:", n)
	fmt.Println("bugg:", bugg[:n], string(bugg[:n]))

	sb = strings.NewReader("teste")
	bggg := bytes.Buffer{}
	n, err = sb.Read(bggg.Bytes())

	fmt.Println("bugg: bytes read:", n)
	fmt.Println("bugg:", bugg[:n], string(bugg[:n]))
}
