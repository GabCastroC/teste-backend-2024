# Teste iCasei: Backend Júnior/Pleno
Esta aplicação é um sistema de cadastro de produtos que mantém e sincroniza dados em duas bases de dados distintas usando mensageria (Kafka).

## Tecnologias e ferramentas utilizadas

- Golang
- Ruby-rails
- Kafka
- MongoDB
- Insomnia
- SQLite 
- Docker
- GIT

## Instruções para inicializar o projeto

- Com o Docker instalado no seu ambiente de desenvolvimento, rode o comando para subir os serviços do docker:

~~~docker composer up -d --build~~~

- Para acessar o banco de dados MongoDB utilize uma IDE de sua preferência (utilizei DataGrip) e conecte com as chaves encontradas no 'docker-compose.yml'.

- Para acessar o banco de dados SQLite utilize uma IDE de sua preferência (utilizei DataGrip) e aponte o arquivo 'ms-rails\storage\development.sqlite3'.

- Com Insomnia isntalado em seu ambiente de desenvolvimento, importe o arquivo 'Insomnia_teste_backend.json' para acessar os end-points da aplicação

## Funcionalidades Rails

1. Cadastro e listagem de produtos.
2. Gestão e armazenamento de dados pelo SQLite.
3. Envio de mensagem com o tópico 'rails-to-go'.

## Melhorias Rails

1. Consumir a mensagem do tópico 'rails-to-go' e atualizar a base de dados
2. Implementação de testes unitários

### Funcionalidades Golang

1. Cadastro e listagem de produtos.
2. Gestão e armazenamento de dados pelo MongoDB.
3. Envio de mensagem com o tópico 'go-to-rails'.

## Melhorias Rails

1. Consumir a mensagem do tópico 'go-to-rails' e atualizar a base de dados.
2. Implementação de testes unitários

