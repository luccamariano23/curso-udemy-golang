package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona a arquitetura amd64")
	}
	t.Fail()
}

//para verificar quanto de um arquivo ja fora testado
//rodar o seguinte comando via terminal:
//go tool cover -html=nome_arquivo.out
