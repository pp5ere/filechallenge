# O que é filechallenge

É um desafio desenvolvido em Go que cria um serviço onde, ao receber um arquivo insere no banco de dados relacional PostGres.

# Como executar

Para executar o aplicativo basta instalar o docker e docker-compose. Após basta executar o arquivo docker-compose com o comando:

```
docker-compose up -d
```

# Como utilizar
Após o subir todos os containers, basta acessar o navegador de sua preferência e enviar o arquivo acessando a url:

```
http://localhost:8080
```

Para visualizar os dados inseridos no banco de dados Postgres basta acessar o PgAdmin na seguinte url:

```
http://localhost:8000
```

No primeiro acesso será solicitado o email e senha, caso não tenha alterado, o email e senha são respectivamentes:

```
1234@admin.com
1234
```
Crie a conexão com o servidor do Postgres para poder visualizar os dados, clicando sobre Servers com o botão direito, Create -> Server ...
Informe o nome do servidor e clique na aba Connection, preencha os dados conforme abaixo e salve:

```
Host name: postgres_container
Port: 5432
Maintenance database: postgres
Username: postgres
Password: 1234
```
