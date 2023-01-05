package idgener

import (
	"github.com/oklog/ulid/v2"
)

type ULIDGen struct{}

func NewULIDGen() (*ULIDGen, error) {
	g := new(ULIDGen)
	return g, nil
}

func (g *ULIDGen) Next() (string, error) {
	return ulid.Make().String(), nil
}
func (g *ULIDGen) String() string {
	return "ulid"
}
