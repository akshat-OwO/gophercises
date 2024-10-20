package main

import (
	"flag"
	"fmt"
	"os"

	chooseyourownadventure "github.com/akshat-OwO/gophercises/choose_your_own_adventure"
)

func main() {
	fileName := flag.String("file", "gopher.json", ".json file containing story")
	flag.Parse()
	fmt.Printf("using %s file for story.\n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := chooseyourownadventure.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(story)
}
