package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func brasilApi(brasilApiChannel chan interface{}) {
	brasilApiUrl := "https://brasilapi.com.br/api/cep/v1/77960000"
	request, err := http.NewRequest("GET", brasilApiUrl, nil)
	if err != nil {
		log.Panicf("Error creating request: %s", err.Error())
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Panicf("Error making request: %s", err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Panicf("Error reading response body: %s", err.Error())
	}

	var cep interface{}

	err = json.Unmarshal(body, &cep)
	if err != nil {
		log.Panicf("Error decoding response: %s", err.Error())
	}

	brasilApiChannel <- cep
}

func viaCepApi(viaCepApiCHannel chan interface{}) {
	viaCepApiUrl := "http://viacep.com.br/ws/77960000/json/"
	request, err := http.NewRequest("GET", viaCepApiUrl, nil)
	if err != nil {
		log.Panicf("Error creating request: %s", err.Error())
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Panicf("Error making request: %s", err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Panicf("Error reading response body: %s", err.Error())
	}

	var cep interface{}

	err = json.Unmarshal(body, &cep)
	if err != nil {
		log.Panicf("Error decoding response: %s", err.Error())
	}

	viaCepApiCHannel <- cep
}

func main() {
	brasilApiChannel := make(chan interface{})
	viaCepApiCHannel := make(chan interface{})

	go viaCepApi(viaCepApiCHannel)
	go brasilApi(brasilApiChannel)

	select {
	case cep := <-brasilApiChannel:
		log.Printf("A API Brasil Cep teve a reguinte resposta: %+v\n", cep)
	case cep := <-viaCepApiCHannel:
		log.Printf("A API Via Cep teve a reguinte resposta: %+v\n", cep)
	case <-time.After(time.Second * 1):
		log.Printf("Timeout na resposta das apis")
	}
}
