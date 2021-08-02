package main

import "testing"

func TestSoma(t *testing.T) {

	total := soma(11, 11)

	if total != 22 {
		t.Errorf("Resultado invalido")
	}
}
