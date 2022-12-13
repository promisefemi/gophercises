package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

// var defaultHandlerTempl = `
// <!DOCTYPE html>
// <html lang="en">
//   <head>
//     <meta charset="UTF-8" />
//     <meta http-equiv="X-UA-Compatible" content="IE=edge" />
//     <meta name="viewport" content="width=device-width, initial-scale=1.0" />
//     <title>Document</title>
//   </head>
//   <body>
//     <h1>{{.Title}}</h1>
//     {{range .Paragraphs}}
//     <p>{{.}}</p>
//     {{end}}
//     <ul>
//       {{range .Options}}
//       <li>
//         <a href="/{{.Chapter}}">{{.Text}}</a>
//       </li>
//       {{end}}
//     </ul>
//   </body>
// </html>

// `
var templ *template.Template

func init() {
	templ = template.Must(template.New("").Parse(ParseTemplateHTML("cmd/cyoa/cyoahtml.html")))
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

func ParseTemplateHTML(filePath string) string {
	defaultTemplateFile, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}
	return string(defaultTemplateFile)
}

func JSONStory(fileReader io.Reader) (Story, error) {

	d := json.NewDecoder(fileReader)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

// Configurations for Handlers
type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func NewHandler(s Story, options ...HandlerOption) http.Handler {
	h := handler{s, templ, defaultPathFn}

	for _, opt := range options {
		opt(&h)
	}
	return h
}

func defaultPathFn(r *http.Request) string {
	path := r.URL.Path

	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}
func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)
	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(rw, chapter)
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
