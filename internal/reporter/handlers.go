package reporter

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	log "github.com/obalunenko/logger"

	"github.com/obalunenko/scrum-report/internal/reporter/assets"
)

type report struct {
	Yesterday   []string
	Today       []string
	Impediments []string
}

func createHandler(_ context.Context) http.HandlerFunc {
	reportPageHTML := string(assets.MustAsset("report.gohtml"))
	reportPageTmpl := template.Must(template.New("report").Parse(reportPageHTML))

	return func(writer http.ResponseWriter, request *http.Request) {
		if err := request.ParseForm(); err != nil {
			http.Error(writer, "failed to parse form", http.StatusBadRequest)

			return
		}

		today := processFormValue(request.Form.Get("today"))
		yesterday := processFormValue(request.Form.Get("yesterday"))
		impediments := processFormValue(request.FormValue("impediments"))

		writer.Header().Set("Content-Type", "text/html")

		err := reportPageTmpl.Execute(writer, report{
			Yesterday:   yesterday,
			Today:       today,
			Impediments: impediments,
		})
		if err != nil {
			http.Error(writer, "failed to execute template", http.StatusInternalServerError)

			return
		}
	}
}

func indexHandler(_ context.Context) http.HandlerFunc {
	homePageHTML := string(assets.MustAsset("index.gohtml"))
	homePageTmpl := template.Must(template.New("index").Parse(homePageHTML))

	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html")

		if err := homePageTmpl.Execute(writer, nil); err != nil {
			http.Error(writer, "failed to execute template", http.StatusInternalServerError)

			return
		}
	}
}

// optionsHandlerOld set up allowed verbs.
func optionsHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Allow", "GET,POST")
}

// loggerHandlerOld Log all HTTP requests to output in a proper format.
func loggerHandler(ctx context.Context, inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		msg := fmt.Sprintf("%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))

		log.Debug(ctx, msg)
	})
}
