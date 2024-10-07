# Multithreading Full Cycle challenge

Este projeto implementa um buscador de CEP utilizando multithreading para buscar o resultado mais rápido entre duas APIs distintas. As requisições são feitas simultaneamente, e a resposta mais rápida é escolhida para ser exibida no terminal, juntamente com a informação de qual API forneceu os dados. Se nenhuma das APIs responder em até 1 segundo, é exibida uma mensagem de erro indicando o timeout.

## 🚧 Como instalar e rodar o projeto

### 1. Clonar o repositório

```bash
git clone https://github.com/edsonjuniordev/multithreading-fullcycle-challenge.git
cd multithreading-fullcycle-challenge
```

### 2. Rodar o desafio

```bash
go run main.go
```