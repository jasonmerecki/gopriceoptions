package gopriceoptions

import (
	"testing"
	"fmt"
)

func TestErrFunc(t *testing.T) {
	ef := Errf(0.56)
	efx := 0.5716157766617889
	msg := fmt.Sprintf("for errfunc, got %f, expected %f\n", ef, efx)
	if ef != efx {
		t.Error(msg)
	}
	fmt.Print(msg)
}

func TestCreateNormdist(t *testing.T) {
	n := NewNormdist(34.3, 3.22)
	fmt.Printf("Got normdist object {%v} \n", n)
}

func TestNormPdf(t *testing.T) {
	m := 34.3
	s := 3.22
	n := NewNormdist(m, s)
	p := n.Pdf(37.5)
	epdf := 0.0756129282351069 
	msg := fmt.Sprintf("with mean=%f and stddev=%f, for pdf, got %f, expected %f\n", m, s, p, epdf)
	if p != epdf {
		t.Error(msg)
	}
	fmt.Print(msg)
}

func TestNormCdf(t *testing.T) {
	m := 34.3
	s := 3.22
	n := NewNormdist(m, s)
	c := n.Cdf(37.8)
	ecdf := 0.8614719786451529
	msg := fmt.Sprintf("with mean=%f and stddev=%f, for cdf, got %f, expected %f\n", m, s, c, ecdf)
	if c != ecdf {
		t.Error(msg)
	}
	fmt.Print(msg)
}

func TestStdPdf(t *testing.T) {
	x := 0.56
	p := Stdnorm.Pdf(x)
	epdf := 0.34104578863035256
	msg := fmt.Sprintf("with Stdnorm and x %f, for pdf, got %f, expected %f\n", x, p, epdf)
	if p != epdf {
		t.Error(msg)
	}
	fmt.Print(msg)
	x = -0.56
	p = Stdnorm.Pdf(x)
	msg = fmt.Sprintf("with Stdnorm and x %f, for pdf, got %f, expected %f\n", x, p, epdf)
	if p != epdf {
		t.Error(msg)
	}
	fmt.Print(msg)
}

func TestStdCdf(t *testing.T) {
	x := 0.56
	c := Stdnorm.Cdf(x)
	// NOTE: outcome in Java is 0.7122603051006894
	ecdf := 0.7122603051006893
	msg := fmt.Sprintf("with Stdnorm and x %f, for cdf, got %f, expected %f\n", x, c, ecdf)
	if c != ecdf {
		t.Error(msg)
	}
	fmt.Print(msg)
}



