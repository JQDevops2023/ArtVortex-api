package graph

//go:generate go get github.com/99designs/gqlgen
//go:generate go run github.com/99designs/gqlgen generate
import "artvortex-api/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	user model.User
}
