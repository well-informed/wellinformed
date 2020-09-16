package main

import (
	"context"
	"testing"

	"github.com/go-chi/chi"
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/graph"
	"github.com/well-informed/wellinformed/graph/model"
)

func NewTestHarness() (wellinformed.Config, *chi.Mux, *graph.Resolver) {
	return initDependencies()
}

func TestRegister(t *testing.T) {
	_, _, resolver := NewTestHarness()
	input := model.RegisterInput{
		Username:        "deviator",
		Email:           "danielveenstra@protonmail.com",
		Password:        "ScoobyDoo69",
		ConfirmPassword: "ScoobyDoo69",
		Firstname:       "Dan",
		Lastname:        "Veenstra",
	}
	_, err := resolver.Mutation().Register(context.Background(), input)
	if err != nil {
		t.Error("failed to register. err: ", err)
	}
}
