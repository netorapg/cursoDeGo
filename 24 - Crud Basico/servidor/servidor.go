package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		http.Error(w, "Falha ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		http.Error(w, "Erro ao converter o usuário", http.StatusBadRequest)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao conectar no banco de dados: "+erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Usando RETURNING id para obter o ID inserido
	query := "INSERT INTO usuarios (nome, email) VALUES ($1, $2) RETURNING id"

	// Executa a query e recupera o ID diretamente
	var idInserido uint32
	erro = db.QueryRow(query, usuario.Nome, usuario.Email).Scan(&idInserido)
	if erro != nil {
		http.Error(w, "Erro ao inserir no banco de dados: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}
