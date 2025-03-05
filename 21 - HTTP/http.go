package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bem vindo a Home"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar Pagina de uisuarios"))
}

func main() {

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)
	log.Fatal(http.ListenAndServe(":5000", nil))

}

// HTTP [E UM PROTOCOLO DE COMUNICACAO - BASE DA COMUNICACAO WEB
//CLIENTE (FAZ A RQUISICAO) - SERVIDOR(PROCESSA A REQUISICAO E ENVIA RESPOSTA)
//Request - Response
//Request: GET, POST, PUT, DELETE
//Response: 200, 404, 500
//Rotas
//URI - Identificador do Recurso
//URL - Localizador do Recurso
//URN - Nome do Recurso
//Metodos
//GET - Obter dados
//POST - Enviar dados
//PUT - Atualizar dados
//DELETE - Deletar dados
