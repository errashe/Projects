package types

import (
	. "fmt"
)

type Cell float64

func (c *Cell) String() string {
	return Sprintf("%f", *c)
}
