package gopriceoptions

import (
	"math"
	"strings"
)

var sqtwopi float64 = math.Sqrt(2 * math.Pi)
var IVPrecision = 0.00001

func PriceBlackScholes(otype string, s float64, k float64, t float64, v float64, r float64, q float64) float64 {
	otype = strings.ToUpper(otype)
	var sign float64
	if "C" == otype {
		if t <= 0 {
			return math.Abs(s - k)
		}
		sign = 1
	}
	if "P" == otype {
		if t <= 0 {
			return math.Abs(k - s)
		}
		sign = -1
	}
	if sign == 0 {
		return 0.0
	}

	re := math.Exp(-r * t)
	qe := math.Exp(-q * t)
	vt := (v * (math.Sqrt(t)))
	d1 := d1f(s, k, t, v, r, q, vt)
	d2 := d2f(d1, vt)
	d1 = sign * d1
	d2 = sign * d2
	nd1 := Stdnorm.Cdf(d1)
	nd2 := Stdnorm.Cdf(d2)

	bsprice := sign * ((s * qe * nd1) - (k * re * nd2))
	return bsprice
}

func d1f(s float64, k float64, t float64, v float64, r float64, q float64, vt float64) float64 {
	d1 := math.Log(s/k) + (t * (r - q + ((v * v) * 0.5)))
	d1 = d1 / vt
	return d1
}

func d2f(d1 float64, vt float64) float64 {
	d2 := d1 - vt
	return d2
}

func d1pdff(s float64, k float64, v float64, t float64, r float64, q float64) float64 {
	vt := (v * (math.Sqrt(t)))
	d1 := d1f(s, k, t, v, r, q, vt)
	d1pdf := math.Exp(-(d1 * d1) * 0.5)
	d1pdf = d1pdf / sqtwopi
	return d1pdf
}

func BSDelta(otype string, s float64, k float64, t float64, v float64, r float64, q float64) float64 {
	var zo float64
	if "P" == otype {
		zo = -1
	}
	if "C" == otype {
		zo = 0
	}
	drq := math.Exp(-q * t)
	vt := (v * (math.Sqrt(t)))
	d1 := d1f(s, k, t, v, r, q, vt)
	cdfd1 := Stdnorm.Cdf(d1)
	delta := drq * (cdfd1 + zo)
	return delta
}

func BSVega(s float64, k float64, t float64, v float64, r float64, q float64) float64 {
	d1pdf := d1pdff(s, k, v, t, r, q)
	drq := math.Exp(-q * t)
	sqt := math.Sqrt(t)
	vega := (d1pdf) * drq * s * sqt * 0.01
	return vega
}

func BSGamma(s float64, k float64, t float64, v float64, r float64, q float64) float64 {
	drq := math.Exp(-q * t)
	drd := (s * v * math.Sqrt(t))
	d1pdf := d1pdff(s, k, v, t, r, q)
	gamma := (drq / drd) * d1pdf
	return gamma
}

func BSTheta(otype string, s float64, k float64, t float64, v float64, r float64, q float64) float64 {

	var sign float64
	if "P" == otype {
		sign = -1
	}
	if "C" == otype {
		sign = 1
	}
	sqt := math.Sqrt(t)
	drq := math.Exp(-q * t)
	dr := math.Exp(-r * t)
	d1pdf := d1pdff(s, k, v, t, r, q)
	twosqt := 2 * sqt
	p1 := -1 * ((s * v * drq) / twosqt) * d1pdf

	vt := (v * (sqt))
	d1 := d1f(s, k, t, v, r, q, vt)
	d2 := d2f(d1, vt)
	var nd1, nd2 float64

	d1 = sign * d1
	d2 = sign * d2
	nd1 = Stdnorm.Cdf(d1)
	nd2 = Stdnorm.Cdf(d2)

	p2 := -sign * r * k * dr * nd2
	p3 := sign * q * s * drq * nd1
	theta := (p1 + p2 + p3) / 365
	return theta
}

func BSRho(otype string, s float64, k float64, t float64, v float64, r float64, q float64) float64 {
	var sign float64
	if "P" == otype {
		sign = -1
	}
	if "C" == otype {
		sign = 1
	}

	dr := math.Exp(-r * t)
	p1 := sign * (k * t * dr) / 100

	vt := (v * (math.Sqrt(t)))
	d1 := d1f(s, k, t, v, r, q, vt)
	d2 := sign * d2f(d1, vt)
	nd2 := Stdnorm.Cdf(d2)
	rho := p1 * nd2
	return rho
}

func BSImpliedVol(otype string, p float64, s float64, k float64, t float64, v float64, r float64, q float64) float64 {
	if v > 0 == false {
		v = 0.5
	}
	errlim := IVPrecision
	maxl := 100
	dv := errlim + 1
	n := 0
	maxloops := 100

	for ; math.Abs(dv) > errlim && n < maxl; n++ {
		difval := PriceBlackScholes(otype, s, k, t, v, r, q) - p
		v1 := BSVega(s, k, v, t, r, q) / 0.01
		dv = difval / v1
		v = v - dv
	}
	var iv float64
	if n < maxloops {
		iv = v
	} else {
		iv = math.NaN()
	}

	return iv
}
