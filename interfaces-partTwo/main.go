package main

import "fmt"

type IEncoder interface {
	Encode(value string)
	Decode(value string)
}

type XEncoder struct {
}

type YEncoder struct {
}

func (xEncoder *XEncoder) Encode(value string) {
	// implementation
	fmt.Println("XEncoder ile encod edildi", value)
}
func (xEncoder *XEncoder) Decode(value string) {
	// implementation
	fmt.Println("XEncoder ile Decode edildi", value)
}

func (yEncoder *YEncoder) Encode(value string) {
	// implementation
	fmt.Println("YEncoder ile encod edildi", value)
}
func (yEncoder *YEncoder) Decode(value string) {
	// implementation
	fmt.Println("YEncoder ile Decode edildi", value)
}

func main() {
	var xEncoder *XEncoder
	xEncoder = &XEncoder{}
	xEncoder.Encode("deneme")
	xEncoder.Decode("deneme")

	yEncoder := &YEncoder{}
	// yEncoder = &YEncoder{}
	yEncoder.Encode("deneme")
	yEncoder.Decode("deneme")

	var encoder IEncoder
	encoder = &XEncoder{}
	encoder.Encode("deneme Interface")
	encoder.Decode("deneme Interface")

	encoder = &YEncoder{}
	encoder.Encode("deneme Interface")
	encoder.Decode("deneme Interface")

}
