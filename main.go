package pet_xk6

import (
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/pet-util", new(PETUtil))
}

type PETUtil struct{}

func (pet *PETUtil) Test() string {
	return "testestes"
}
