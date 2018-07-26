package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//CEP Guarda informações do CEP
type CEP struct {
	DsCEP         string `json:"cep"`
	DsLogradouro  string `json:"logradouro"`
	DsComplemento string `json:""`
	DsBairro      string `json:"bairro"`
	DsCidade      string `json:"localidade"`
	DsUf          string `json:"uf"`
	CdCidade      string `json:"ibge"`
}

func main() {

	var dadosCEP CEP

	resposta, _ := http.Get("https://viacep.com.br/ws/37900040/json/")

	dados, _ := ioutil.ReadAll(resposta.Body)

	json.Unmarshal([]byte(dados), &dadosCEP)

	fmt.Println(dadosCEP)

}
