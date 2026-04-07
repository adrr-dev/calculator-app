// Package service contains the logic service
package service

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/adrr-dev/calculator-app/internal/repository"
)

type CalcRepo interface {
	FetchDisplay() (*repository.Display, error)
	WriteData(data *repository.Display) error
}
type Service struct {
	Repo CalcRepo
}

func NewService(repo CalcRepo) *Service {
	newService := &Service{Repo: repo}
	return newService
}

func (s Service) NewDisplay(dis string) (*repository.Display, error) {
	display := &repository.Display{Display: dis}

	// saves data immediately
	err := s.Repo.WriteData(display)
	if err != nil {
		return nil, err
	}
	return display, nil
}

func (s Service) ShowDisplay(character string) (*repository.Display, error) {
	currentDisplay, err := s.Repo.FetchDisplay()
	if err != nil {
		return nil, err
	}
	newDisplay := fmt.Sprintf("%s%s", currentDisplay.Display, character)
	Display := &repository.Display{Display: newDisplay}

	// saves data immediately
	err = s.Repo.WriteData(Display)
	if err != nil {
		return nil, err
	}

	return Display, nil
}

func (s Service) EvalString(eval string) float64 {
	expression, _ := govaluate.NewEvaluableExpression(eval)
	result, _ := expression.Evaluate(nil)
	return result.(float64)
}
