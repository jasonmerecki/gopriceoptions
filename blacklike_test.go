package gopriceoptions

import (
	"fmt"
	"math"
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

func TestBSCallGreeks(t *testing.T) {
	otype := "C"
	s := 1177.62
	k := 1195.00
	time := 0.084931506849315 // date 12/19/2017, expiration 1/19/2018, 31 days
	v := 0.20
	r := 0.0135
	q := 0.0
	delta := BSDelta(otype, s, k, time, v, r, q)
	gamma := BSGamma(s, k, time, v, r, q)
	vega := BSVega(s, k, time, v, r, q)
	theta := BSTheta(otype, s, k, time, v, r, q)
	rho := BSRho(otype, s, k, time, v, r, q)
	msg := fmt.Sprintf("TestBSCallGreeks, delta %.24f, gamma %f, vega %f, theta %f, rho %f\n", delta, gamma, vega, theta, rho)
	edelta := 0.4197454548230388
	egamma := 0.005694
	evega := 1.341348
	etheta := -0.450224
	erho := 0.402579
	if delta != edelta || math.Abs(egamma - gamma) > 0.00001 || math.Abs(evega - vega) > 0.00001 || math.Abs(etheta - theta) > 0.00001 || math.Abs(erho - rho) > 0.00001 {
		t.Error(msg)
	}
	fmt.Print(msg)
}

func TestBSPutGreeks(t *testing.T) {
	otype := "P"
	s := 214.76
	k := 190.00
	time := 0.084931506849315 // date 12/19/2017, expiration 1/19/2018, 31 days
	v := 0.25
	r := 0.0135
	q := 0.0
	delta := BSDelta(otype, s, k, time, v, r, q)
	gamma := BSGamma(s, k, time, v, r, q)
	vega := BSVega(s, k, time, v, r, q)
	theta := BSTheta(otype, s, k, time, v, r, q)
	rho := BSRho(otype, s, k, time, v, r, q)
	msg := fmt.Sprintf("TestBSPutGreeks, delta %.24f, gamma %f, vega %f, theta %f, rho %f\n", delta, gamma, vega, theta, rho)
	edelta := -0.04150437210202529
	egamma := 0.005675
	evega := 0.055574
	etheta := -0.022069
	erho := -0.007800
	if delta != edelta || math.Abs(egamma - gamma) > 0.00001 || math.Abs(evega - vega) > 0.00001 || math.Abs(etheta - theta) > 0.00001 || math.Abs(erho - rho) > 0.00001 {
		t.Error(msg)
	}
	if delta != edelta {
		t.Error(msg)
	}
	fmt.Print(msg)
}

func TestBSImpVol(t *testing.T) {
	otype := "C"
	p := 20.29616
	s := 1177.62
	k := 1195.00
	time := 0.084931506849315 // date 12/19/2017, expiration 1/19/2018, 31 days
	r := 0.0135
	q := 0.0
	biv := BSImpliedVol(otype, p, s, k, time, 0.0, r, q)
	msg := fmt.Sprintf("TestBSImpVol, implied vol %f \n", biv)
	diff := math.Abs(biv - 0.20)
	if diff > 0.00001 {
		t.Error(msg)
	}
	fmt.Print(msg)
}
