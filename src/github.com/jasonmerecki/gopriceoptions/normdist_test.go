package gopriceoptions

import (
	"testing"
	"fmt"
)

func TestErrFunc(t *testing.T) {
	ef := ErrFunc(0.56)
	fmt.Printf("ef result = %f \n", ef)
	efx := 0.5716157766617889
	if ef != efx {
		t.Errorf("got %f, expected %f", ef, efx)
	}
}
