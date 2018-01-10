package gopriceoptions

import (
	"fmt"
	"testing"
)

func TestBlackCall1(t *testing.T) {
	s := 1177.62
	k := 1195.00
	time := 0.084931506849315 // date 12/19/2017, expiration 1/19/2018, 31 days
	v := 0.20
	r := 0.0135
	q := 0.0
	bsprice := PriceBlackScholes("C", s, k, time, v, r, q)
	eprice := 20.29616303951127
	msg := fmt.Sprintf("TestBlackCall1, got %f, expected %f\n", bsprice, eprice)
	if bsprice != eprice {
		t.Error(msg)
	}
	fmt.Print(msg)
}

func TestBlackPut1(t *testing.T) {
	s := 214.76
	k := 190.00
	time := 0.084931506849315 // date 12/19/2017, expiration 1/19/2018, 31 days
	v := 0.25
	r := 0.0135
	q := 0.0
	bsprice := PriceBlackScholes("P", s, k, time, v, r, q)
	eprice := 0.2707906395245452
	msg := fmt.Sprintf("TestBlackPut1, got %f, expected %f\n", bsprice, eprice)
	if bsprice != eprice {
		t.Error(msg)
	}
	fmt.Print(msg)

}
