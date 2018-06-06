package main

import (
	"strings"
	"testing"

	"github.com/azer/snakecase"
	"github.com/codemodus/kace"
	"github.com/iancoleman/strcase"
	gostrcase "github.com/stoewer/go-strcase"
)

func BenchmarkStrcaseSnake6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strcase.ToSnake("MyLovelyAPIMethodIsDumb")
	}
}

func BenchmarkGostrcaseSnake6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = gostrcase.SnakeCase("MyLovelyAPIMethodIsDumb")
	}
}

func BenchmarkSnakecaseSnake6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = snakecase.SnakeCase("MyLovelyAPIMethodIsDumb")
	}
}

func BenchmarkKaceSnake6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = kace.Snake("MyLovelyAPIMethodIsDumb")
	}
}

func BenchmarkLower6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strings.ToLower("MyLovelyAPIMethodIsDumb")
	}
}

func BenchmarkStrcaseUnicode6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strcase.ToSnake("ТвоюЖеМатьКакЖеТак")
	}
}

func BenchmarkGostrcaseUnicode6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = gostrcase.SnakeCase("ТвоюЖеМатьКакЖеТак")
	}
}

func BenchmarkSnakecaseUnicode6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = snakecase.SnakeCase("ТвоюЖеМатьКакЖеТак")
	}
}

func BenchmarkKaceUnicode6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = kace.Snake("ТвоюЖеМатьКакЖеТак")
	}
}

func BenchmarkLowerUnicode6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = strings.ToLower("ТвоюЖеМатьКакЖеТак")
	}
}
