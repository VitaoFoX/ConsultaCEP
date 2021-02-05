﻿<h1 align="center">
  Consulta CEP - Utilizando Golang
</h1>

 # Projeto

Criado para pegar o CEP do site Via CEP
Exemplo bem simples usando http.Get("")

### Exemplo
Onde tem 37900040, trocar para o CEP desejado.
Em main.go tem o código completo



	var dadosCEP CEP

	resposta, _ := http.Get("https://viacep.com.br/ws/37900040/json/")

	dados, _ := ioutil.ReadAll(resposta.Body)

	json.Unmarshal([]byte(dados), &dadosCEP)

	fmt.Println(dadosCEP)

### Running

```sh
go run main.go
```
