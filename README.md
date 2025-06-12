# API de Produtos em Go

Esta é uma simples API REST para gerenciamento de produtos, desenvolvida em Go. A API permite operações de CRUD (Criar, Ler, Deletar) para produtos.

## Tecnologias Utilizadas

* **Go:** Linguagem de programação principal.
* **Gin Gonic:** Framework web para Go, utilizado para o roteamento e manipulação de requisições HTTP.
* **PostgreSQL:** Banco de dados relacional para armazenamento dos dados dos produtos.
* **Docker & Docker Compose:** Utilizado para criar um ambiente de desenvolvimento conteinerizado e facilitar a execução da aplicação e do banco de dados.

## Estrutura do Projeto

O projeto segue uma arquitetura em camadas para uma clara separação de responsabilidades:

* `main.go`: Ponto de entrada da aplicação. Inicializa a conexão com o banco de dados, as dependências e as rotas da API.
* `controller/`: Camada responsável por receber as requisições HTTP, validar os dados de entrada e interagir com os casos de uso (use cases).
* `usecase/`: Contém a lógica de negócio da aplicação. Atua como um intermediário entre os controllers e os repositórios.
* `repository/`: Responsável pela comunicação com o banco de dados, abstraindo as queries SQL.
* `model/`: Define as estruturas de dados (structs) da aplicação, como `Product` e `Response`.
* `db/`: Contém a lógica para estabelecer a conexão com o banco de dados PostgreSQL.

## Como Executar

1.  **Inicie os containers:**
    A partir do diretório raiz do projeto, execute o seguinte comando para construir a imagem da aplicação Go e iniciar os containers do Go e do PostgreSQL.

    ```bash
    docker-compose up -d --build
    ```

    * A API estará disponível em `http://localhost:8000`.
    * O banco de dados PostgreSQL estará acessível na porta `5434` da sua máquina local.

## Endpoints da API

A seguir estão os endpoints disponíveis na API:

### Health Check

* **GET** `/ping`
    * Verifica se a API está em execução.
    * **Resposta de Sucesso (200 OK):**
        ```json
        {
            "message": "pong"
        }
        ```

### Produtos

* **GET** `/products`
    * Retorna uma lista de todos os produtos cadastrados.
    * **Resposta de Sucesso (200 OK):**
        ```json
        [
            {
                "id_product": 1,
                "name_product": "Notebook",
                "price_product": 3500.50
            }
        ]
        ```

* **GET** `/product/:product_id`
    * Retorna um produto específico com base no seu `id`.
    * **Resposta de Sucesso (200 OK):**
        ```json
        {
            "id_product": 1,
            "name_product": "Notebook",
            "price_product": 3500.50
        }
        ```
    * **Resposta de Erro (404 Not Found):**
        ```json
        {
            "Message": "Product not found"
        }
        ```

* **POST** `/product`
    * Cria um novo produto.
    * **Corpo da Requisição:**
        ```json
        {
            "name_product": "Teclado Mecânico",
            "price_product": 250.75
        }
        ```
    * **Resposta de Sucesso (201 Created):**
        ```json
        {
            "id_product": 3,
            "name_product": "Teclado Mecânico",
            "price_product": 250.75
        }
        ```

* **DELETE** `/product/:product_id`
    * Deleta um produto com base no seu `id`.
    * **Resposta de Sucesso (200 OK):**
        ```json
        {
            "Message": "Product deleted successfully"
        }
        ```