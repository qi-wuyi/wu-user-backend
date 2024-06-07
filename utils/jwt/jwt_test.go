package jwt

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestCreatToken(t *testing.T) {
	token, err := CreatToken(1)
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(token)
}

func TestValidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJpc3MiOiJ3dS1wYW8iLCJleHAiOjE3MTc3MzczMDIsImlhdCI6MTcxNzQ3ODEwMn0.YIuGhRSdi7rNXGM3VOh6NUVvegL0h57vJiZ5F1oBhvM"
	c, err := ValidateToken(token)
	if err != nil {
		spew.Dump(err)
	}
	spew.Dump(c)
}
