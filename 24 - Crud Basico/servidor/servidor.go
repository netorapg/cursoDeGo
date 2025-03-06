package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuário no banco
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

// BuscarUsuarios busca todos os usuários salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao conectar no banco de dados: "+erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		http.Error(w, "Erro ao buscar os usuários: "+erro.Error(), http.StatusInternalServerError)
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			http.Error(w, "Erro ao escanear o usuário: "+erro.Error(), http.StatusInternalServerError)
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

// BuscarUsuario busca um usuário especifico salvo no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		http.Error(w, "Erro ao converter o parâmetro para inteiro", http.StatusBadRequest)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao conectar no banco de dados: "+erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	linha, erro := db.Query("SELECT * FROM usuarios WHERE id = $1", ID)
	if erro != nil {
		http.Error(w, "Erro ao buscar o usuário: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			http.Error(w, "Erro ao escanear o usuário: "+erro.Error(), http.StatusInternalServerError)
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		http.Error(w, "Erro ao converter o usuário para JSON", http.StatusInternalServerError)
		return
	}

}

// AtualizarUsuario atualiza um usuário especifico salvo no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		http.Error(w, "Erro ao converter o parâmetro para inteiro", http.StatusBadRequest)
		return
	}

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

	statement, erro := db.Prepare("update usuarios set nome = $1, email = $2 where id = $3")
	if erro != nil {
		http.Error(w, "Erro ao criar o statement: "+erro.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		http.Error(w, "Erro ao atualizar o usuário: "+erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
