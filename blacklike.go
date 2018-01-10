package gopriceoptions

import (
	"math"
	"strings"
)

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
