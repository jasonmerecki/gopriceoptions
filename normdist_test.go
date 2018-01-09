package gopriceoptions

import (
	"testing"
	"fmt"
)

func TestErrFunc(t *testing.T) {
	ef := Errf(0.56)
	fmt.Printf("ef result = %f \n", ef)
	efx := 0.5716157766617889
	if ef != efx {
		t.Errorf("got %f, expected %f", ef, efx)
	}
}

func TestCreateNormdist(t *testing.T) {
	n := NewNormdist(34.3, 3.22)
	fmt.Printf("Got normdist object {%v} \n", n)
	// at 37.5 the pdf is 0.0756129282351069 
}

func TestNormPdf(t *testing.T) {
	m := 34.3
	s := 3.22
	n := NewNormdist(m, s)
	p := n.Pdf(37.5)
	epdf := 0.0756129282351069 
	if p != epdf {
		t.Errorf("with mean=%f and stddev=%f, for pdf, got %f, expected %f", m, s, p, epdf)
	}
}

func TestNormCdf(t *testing.T) {
	m := 34.3
	s := 3.22
	n := NewNormdist(m, s)
	c := n.Cdf(37.8)
	ecdf := 0.8614719786451529
	if c != ecdf {
		t.Errorf("with mean=%f and stddev=%f, for cdf, got %f, expected %f", m, s, c, ecdf)
	}
}



