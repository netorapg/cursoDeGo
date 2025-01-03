// TESTES DE UNIDADE
package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista"

	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}
