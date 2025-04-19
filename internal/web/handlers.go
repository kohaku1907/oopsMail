package web

import (
	"html/template"
	"net/http"
)

type TemplateData struct {
	Title string
	Data  interface{}
}

type Handler struct {
	templates map[string]*template.Template
}

func NewHandler() (*Handler, error) {
	// Define template files for each page
	templateFiles := map[string][]string{
		"home": {
			"internal/web/templates/base.html",
			"internal/web/templates/home.html",
		},
		"create": {
			"internal/web/templates/base.html",
			"internal/web/templates/create.html",
		},
		"view": {
			"internal/web/templates/base.html",
			"internal/web/templates/view.html",
		},
	}

	// Parse templates
	templates := make(map[string]*template.Template)
	for name, files := range templateFiles {
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}
		templates[name] = tmpl
	}

	return &Handler{
		templates: templates,
	}, nil
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Title: "Home",
		Data:  nil,
	}

	if err := h.templates["home"].ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateMailbox(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Title: "Create Mailbox",
		Data:  nil,
	}

	if err := h.templates["create"].ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ViewEmails(w http.ResponseWriter, r *http.Request) {
	data := TemplateData{
		Title: "View Emails",
		Data:  nil,
	}

	if err := h.templates["view"].ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
