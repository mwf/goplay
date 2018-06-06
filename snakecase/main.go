package main

import (
	"fmt"

	"github.com/azer/snakecase"
	"github.com/codemodus/kace"
	"github.com/iancoleman/strcase"
	gostrcase "github.com/stoewer/go-strcase"
)

func snakeStrings(s string) {
	fmt.Printf("strcase: %s\n", strcase.ToSnake(s))
	fmt.Printf("gostrcase: %s\n", gostrcase.SnakeCase(s))
	fmt.Printf("snakecase: %s\n", snakecase.SnakeCase(s))
	fmt.Printf("kace: %s\n", kace.Snake(s))
}

func main() {
	testCases := []string{
		"MyLovelyAPIMethod",
		"Address2Point",
		"ТвоюЖеМать",
	}

	for _, tc := range testCases {
		snakeStrings(tc)
		fmt.Println()
	}
}
