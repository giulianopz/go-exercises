package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println("already present api:", stringutil.Reverse("Hello"))
	fmt.Println("new api:", stringutil.ToUpper("Hello"))
}
