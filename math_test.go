package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é invalid: Resultado %d. Esperado %d", total, 30)
	}
}
