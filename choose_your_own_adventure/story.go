package chooseyourownadventure

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func init() {
	tmpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

var tmpl *template.Template

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func JsonStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	err := decoder.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tmpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "chapter not found.", http.StatusNotFound)
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

var defaultHandlerTemplate = `
<!DOCTYPE HTML>
<html>
	<head>
		<meta charset="utf-8">
		<title>Choose Your Own Adventure</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
			<p>{{.}}</p>
		{{end}}
		<ul>
			{{range .Options}}
				<li><a href="{{.Chapter}}">{{.Text}}</a></li>
			{{end}}
		</ul>
	</body>
</html>
	`
