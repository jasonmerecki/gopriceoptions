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

func TestCreateNormdist(t *testing.T) {
	n := NewNormdist(34.3, 3.22)
	fmt.Printf("Got normdist object %v", n)
	// at 37.5 the pdf is 0.0756129282351069 
}

func TestNormPdf(t *testing.T) {
	n := NewNormdist(34.3, 3.22)
	p := n.Pdf(37.5)
	epdf := 0.0756129282351069 
	if p != epdf {
		t.Errorf("got %f, expected %f", p, epdf)
	}
}


