package main

import (
	"fmt"
	"strings"
)

var exampleHtml1 = `
	<html>
	<body>
		<h1>Hello</h1>
		<a href="/other-page">A link to another page</a>
	</body>
	</html>
	`

func main() {
	r := strings.NewReader(exampleHtml1)
	links, err := Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
