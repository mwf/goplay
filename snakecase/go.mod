module github.com/mwf/goplay/snakecase

require (
	github.com/azer/snakecase v0.0.0-20161028114325-c818dddafb5c
	github.com/codemodus/kace v0.5.0
	github.com/iancoleman/strcase v0.0.0-20180605031248-90d371a664d6
	github.com/stoewer/go-strcase v1.0.1

	// Transient dependency via `vgo test all`
	// Check out https://github.com/golang/go/issues/25672#issuecomment-395254850 for details
	github.com/stretchr/testify v1.2.2
)
