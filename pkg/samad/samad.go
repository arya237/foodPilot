package pkg

import (
	"github.com/arya237/foodPilot/pkg"
)

type Samad struct {
	rf pkg.RequiredFunctions
}

func NewSamad(rf pkg.RequiredFunctions) *Samad {
	return &Samad{rf: rf}
}

func (s *Samad) GetAccessToken(studentNumber string, password string) (string, error) {
	return "", nil
}
