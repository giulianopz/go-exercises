package main

import "golang.org/x/tour/reader"

type MyReader struct{
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, err error) {
	
	for i,_ := range b {
		b[i] = 65 // decimal representation of 'A' in ASCII
	}
	return len(b),nil
}

func main() {
	reader.Validate(MyReader{})
}
