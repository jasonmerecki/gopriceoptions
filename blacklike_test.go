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


func TestBSCallGreeks(ts *testing.T) {
	otype := "C"
	s := 1177.62
	k := 1195.00
	t := 0.084931506849315 // date 12/19/2017, expiration 1/19/2018, 31 days
	v := 0.20
	r := 0.0135
	q := 0.0
	delta := BSDelta(otype, s, k, t, v, r, q)
	gamma := BSGamma(s, k, t, v, r, q)
	vega := 	BSVega(s, k, t, v, r, q)
	theta := BSTheta(otype, s, k, t, v, r, q)
	rho := BSRho(otype, s, k, t, v, r, q)
	msg := fmt.Sprintf("TestBSCallGreeks, delta %f, gamma %f, vega %f, theta %f, rho %f\n", delta, gamma, vega, theta, rho)
	// need to double-check Greeks results, they are slightly different from expected
	edelta := 0.4197454548230388
	if delta != edelta {
		ts.Error(msg)
	}
	fmt.Print(msg)
}

func TestBSPutGreeks(ts *testing.T) {
	otype := "P"
	s := 214.76
	k := 190.00
	t := 0.084931506849315 // date 12/19/2017, expiration 1/19/2018, 31 days
	v := 0.25
	r := 0.0135
	q := 0.0
	delta := BSDelta(otype, s, k, t, v, r, q)
	gamma := BSGamma(s, k, t, v, r, q)
	vega := 	BSVega(s, k, t, v, r, q)
	theta := BSTheta(otype, s, k, t, v, r, q)
	rho := BSRho(otype, s, k, t, v, r, q)
	msg := fmt.Sprintf("TestBSPutGreeks, delta %f, gamma %f, vega %f, theta %f, rho %f\n", delta, gamma, vega, theta, rho)
	// need to double-check Greeks results, they are slightly different from expected
	// gamma and theta need work
	edelta := 0.4197454548230388
	if delta != edelta {
		ts.Error(msg)
	}
	fmt.Print(msg)
}



