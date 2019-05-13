package main

import (
	"fmt"
	"github.com/go-modules-by-example/goinfo/designers"
	v2designers "github.com/go-modules-by-example/goinfo/v2/designers"
	"strings"
)

func main() {
	fmt.Printf("The designers of Go: %v\n", strings.Join(designers.Names(), ", "))
	fmt.Printf("The designers of Go: %v\n", strings.Join(v2designers.FullNames(), ", "))
}
