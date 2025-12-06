# GopherBlog 🐹

[![Go version](https://img.shields.io/badge/go-1.XX-blue)]()  
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)]()

> **GopherBlog** é uma API RESTful de blog — feita com Go, Gin, Gorm e PostgreSQL — projetada para ser simples, organizada e escalável desde o início.  

---

## 📌 Visão Geral

GopherBlog oferece um backend minimalista para gerenciamento de posts: criação, edição, publicação, listagem com filtros, paginação e remoção.  
A estrutura foi pensada para reproduzir boas práticas corporativas, mesmo em um projeto simples, seguindo a arquitetura:


Com isso, o código fica limpo, modular e fácil de evoluir — ideal para portfólios ou base de um sistema real.

---

## ✅ Funcionalidades Principais

- Criar post (rascunho)  
- Editar post  
- Publicar / despublicar post  
- Listar posts com filtros (por autor, tag, estado, busca full-text)  
- Paginação (page / limit)  
- Buscar post por ID  
- Deletar post  

---

## 🧱 Modelo de Dados

**Post**

| Campo       | Tipo        | Descrição                          |
|-------------|-------------|------------------------------------|
| `id`        | UUID / int  | Identificador único                |
| `title`     | string      | Título do post                     |
| `body`      | string      | Conteúdo do post                   |
| `author`    | string      | Nome simples do autor              |
| `tags`      | []string    | Tags associadas (opcional)         |
| `published` | bool        | Indica se está publicado ou rascunho |
| `createdAt` | datetime    | Data/hora de criação               |
| `updatedAt` | datetime    | Data/hora da última atualização    |

---

## 🔐 Validações & Regras de Negócio

- `title`: obrigatório, máximo 120 caracteres, não vazio  
- `body`: obrigatório, mínimo 10 caracteres  
- `author`: obrigatório, máximo 60 caracteres  
- `tags`: opcional — até 5 tags, cada uma até 20 caracteres  
- Publicação só é permitida se `title` e `body` forem válidos  
- Não permitir publicar um post já publicado  

---

## 📦 Tecnologias & Ferramentas

- **Go** – linguagem do backend  
- **Gin** – framework HTTP  
- **Gorm** – ORM para acesso a banco  
- **PostgreSQL** – banco relacional  
- **Swagger (OpenAPI)** – documentação da API  
- **Docker / docker-compose** – para facilitar setup local  

---

## 🚀 Como Rodar Localmente

> ⚠️ Presume que você já tem Go e Docker instalados.

```bash
# clone o repositório  
git clone https://github.com/felipematheus1337/Gopher_Blog.git  
cd Gopher_Blog  

# (opcional) configuração de ambiente  
cp .env.example .env  
# edite variáveis se necessário (DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME, etc.)

# subir o banco + serviço via Docker  
docker-compose up --build  

A API estará disponível em http://localhost:8080 por padrão.
Endpoints e documentação podem ser acessados via Swagger (ex: http://localhost:8080/swagger/index.html).

📄 Licença

Licenciado sob os termos da MIT License. Veja o arquivo LICENSE
 para mais detalhes.

# ou (alternativa) rodar diretamente com Go  
go run main.go  
