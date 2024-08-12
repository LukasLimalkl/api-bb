package extrato

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var token string

var data ExtratoBB

var SaldoTotal float64

func getToken() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseUrl := os.Getenv("OAUTH_URL")
	basicToken := os.Getenv("BASIC_TOKEN")

	req, err := http.NewRequest("POST", baseUrl, nil)

	if err != nil {
		log.Fatalln("Deu error ao criar nova requisição:", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", basicToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return
	}

	var oauth Oauth

	err = json.Unmarshal(body, &oauth)

	if err != nil {
		log.Fatal("Error to unmarshal json", err)
	}

	token = oauth.AccessToken

}

func GetExtrato() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	getToken()

	bearer := "Bearer " + token

	extratoUrl := os.Getenv("URL_EXTRATO")

	req, err := http.NewRequest("GET", extratoUrl, nil)

	if err != nil {
		log.Fatalln("Deu error ao criar nova requisição:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)
	req.Header.Set("x-br-com-bb-ipa-mciteste", "178961031")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("Erro ao desfazer json:", err)
		return
	}

	fmt.Println(data)

}

/*func HandleData() float64 {

	GetExtrato()

	for _, lancamento := range data.ListaLancamento {

		//valor do extrato ta vindo errado

		valor := lancamento.ValorLancamento

		if lancamento.IndicadorSinalLancamento == "C" {
			SaldoTotal += valor
		} else if lancamento.IndicadorSinalLancamento == "D" {
			SaldoTotal -= valor
		} else if lancamento.TextoDescricaoHistorico == "Saldo do dia" {
			SaldoTotal -= valor
		}

	}

	return SaldoTotal

}
*/
