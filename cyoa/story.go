package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

var defaultHandlerTempl = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Document</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
      {{range .Options}}
      <li>
        <a href="/{{.Chapter}}">{{.Text}}</a>
      </li>
      {{end}}
    </ul>
  </body>
</html>

`
var templ *template.Template

func init() {
	templ = template.Must(template.New("").Parse(defaultHandlerTempl))
}

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

type Demo struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func JSONStory(fileReader io.Reader) (Story, error) {

	d := json.NewDecoder(fileReader)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

func NewHandler(s Story, tmpl *template.Template) http.Handler {
	t := templ
	if tmpl != nil {
		t = tmpl
	}

	return handler{s, t}
}

type handler struct {
	s Story
	t *template.Template
}

func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := templ.Execute(rw, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(rw, "Something went wrong...", http.StatusInternalServerError)
			// 	panic(err)
		}
		return
	} else {
		http.NotFound(rw, r)
	}

}
