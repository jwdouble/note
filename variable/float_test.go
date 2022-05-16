package variable

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func Test_trunc(t *testing.T) {
	i := 1.79843434345843
	fmt.Println(i)
	fmt.Println(decimal.NewFromFloat(i).Round(0))
}
