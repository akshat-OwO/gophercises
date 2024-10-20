package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	chooseyourownadventure "github.com/akshat-OwO/gophercises/choose_your_own_adventure"
)

func main() {
	port := flag.String("port", "3000", "port to run the server")
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

	h := chooseyourownadventure.NewHandler(story)
	fmt.Printf("starting the server on port: %s\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(": %s", *port), h))
}
