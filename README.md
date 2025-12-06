📝 GopherBlog

Uma API simples de blog construída com Go, Gin, Gorm e PostgreSQL.
O objetivo é oferecer uma estrutura limpa, didática e próxima do que empresas usam, seguindo as camadas routes → handler → service → db, com documentação via Swagger.

🚀 Objetivo do Projeto

O GopherBlog é um CRUD de posts com foco em boas práticas:

Organização em camadas

Validações no handler

Regras de negócio no service

Persistência com Gorm

Filtros reais (por autor, tag, busca e estado publicado)

Ponto de entrada simples e escalável

Ideal para quem está aprendendo Go e quer um projeto profissional, mas sem relações complexas entre tabelas.

🧱 Entidade Principal
Post
Campo	Tipo	Descrição
id	int/uuid	Identificador único
title	string	Título do post
body	string	Conteúdo do post
author	string	Nome simples do autor
tags	string	Lista de tags em formato simples
published	bool	Indica se o post está publicado
createdAt	datetime	Data de criação
updatedAt	datetime	Data de atualização
🎯 Casos de Uso
✔ Criar post

Cria um post em rascunho (published = false por padrão).

✔ Editar post

Atualiza título, conteúdo, tags e autor.

✔ Publicar post

Altera o estado para publicado caso a validação permita.

✔ Despublicar post

Retorna o post para rascunho.

✔ Listar posts

Permite filtros opcionais:

published=true/false

author=nome

tag=go

search=palavra

page e limit para paginação

✔ Buscar post por ID

Retorna um único post.

✔ Excluir post

Remove definitivamente.

🔍 Validações Importantes
title

obrigatório

máximo 120 caracteres

não pode ser vazio ou espaços

body

obrigatório

mínimo 10 caracteres

author

obrigatório

até 60 caracteres

tags

opcional

máximo 5 tags

cada tag com até 20 caracteres

publicação

só publica se title e body forem válidos

não publica se já estiver publicado

🧩 Arquitetura
/routes      → registra endpoints  
/handler     → recebe requisição, valida dados, chama o service  
/service     → regras de negócio (publicar, filtrar, etc.)  
/db          → camada de persistência com Gorm


Separação simples e muito usada em empresas. Fácil de manter e evoluir.

📚 Tecnologias

Go

Gin Framework

Gorm ORM

PostgreSQL

Swagger (OpenAPI)

Docker (opcional para subir o banco)

📌 Próximos Passos (opcional)

Sistema de autores real (segunda tabela)

Comentários

Likes

Autenticação JWT

Middleware de logs

Cache (Redis)

Paginação avançada

Pré-visualização de markdown

🐹 GopherBlog — simples hoje, escalável amanhã.