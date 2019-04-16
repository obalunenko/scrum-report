package reporter

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/oleg-balunenko/scrum-report/web"
)

var (
	homePageHTML = string(web.MustAsset("templates/index.html"))
	homePageTmpl = template.Must(template.New("index").Parse(homePageHTML))

	reportPageHTML = string(web.MustAsset("templates/report.html"))
	reportPageTmpl = template.Must(template.New("report").Parse(reportPageHTML))
)

type report struct {
	Yesterday   []string
	Today       []string
	Impediments []string
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	var r report

	today := request.FormValue("today")
	r.Today = strings.Split(today, "\n")

	yesterday := request.FormValue("yesterday")
	r.Yesterday = strings.Split(yesterday, "\n")

	impediments := request.FormValue("impediments")
	r.Impediments = strings.Split(impediments, "\n")

	writer.Header().Set("Content-Type", "text/html")
	if err := reportPageTmpl.Execute(writer, r); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, "failed to execute template")
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	if err := homePageTmpl.Execute(writer, nil); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, "failed to execute template")
	}

}

// optionsHandler set up allowed verbs
func optionsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Allow", "GET,POST")
}

// loggerHandler Log all HTTP requests to output in a proper format.
func loggerHandler(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Debugf("%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
