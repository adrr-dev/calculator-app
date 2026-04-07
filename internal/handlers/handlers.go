// Package handlers contains the handling logic
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/adrr-dev/calculator-app/internal/repository"
)

type CalcService interface {
	NewDisplay(dis string) (*repository.Display, error)
	EvalString(eval string) float64
	ShowDisplay(character string) (*repository.Display, error)
}

type Handling struct {
	Service  CalcService
	Tmpls    *template.Template
	Fragment *template.Template
}

func NewHandling(tmpls, fragments *template.Template, service CalcService) *Handling {
	handling := &Handling{Service: service, Tmpls: tmpls, Fragment: fragments}
	return handling
}

func (h Handling) RootHandler(w http.ResponseWriter, r *http.Request) {
	display, err := h.Service.NewDisplay("")
	if err != nil {
		message := fmt.Sprintf("display not found: %e", err)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	err = h.Tmpls.ExecuteTemplate(w, "calculator.html", display)
	if err != nil {
		message := fmt.Sprintf("template not found: %e", err)
		http.Error(w, message, http.StatusNotFound)
		return
	}
}

func (h Handling) KeyHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")

	newDisplay, err := h.Service.ShowDisplay(key)
	if err != nil {
		message := fmt.Sprintf("display not found: %e", err)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	err = h.Fragment.ExecuteTemplate(w, "display.html", newDisplay)
	if err != nil {
		log.Fatal(err)
	}
}

func (h Handling) EnterHandler(w http.ResponseWriter, r *http.Request) {
	display, err := h.Service.ShowDisplay("")
	if err != nil {
		message := fmt.Sprintf("display not found: %e", err)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	evaluable := display.Display

	floater := h.Service.EvalString(evaluable)
	data := fmt.Sprintf("%.2f", floater)

	newDisplay, err := h.Service.NewDisplay(data)
	if err != nil {
		message := fmt.Sprintf("display not found: %e", err)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	err = h.Fragment.ExecuteTemplate(w, "display.html", newDisplay)
	if err != nil {
		log.Fatal(err)
	}
}

func (h Handling) ClearHandler(w http.ResponseWriter, r *http.Request) {
	display, err := h.Service.NewDisplay("")
	if err != nil {
		message := fmt.Sprintf("display not found: %e", err)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	err = h.Fragment.ExecuteTemplate(w, "display.html", display)
	if err != nil {
		message := fmt.Sprintf("template not found: %e", err)
		http.Error(w, message, http.StatusNotFound)
		return
	}
}
