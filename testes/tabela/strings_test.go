package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte %s) - Índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Coder é top", "Coder", 0},
		{"", "", 0},
		{"Ola", "ola", -1},
		{"Lucca", "a", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
