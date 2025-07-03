# BASIC-CRUD-GO

Este é um projeto de portfólio desenvolvido em **Golang**, utilizando os princípios de **Domain-Driven Design (DDD)**.
A aplicação implementa um CRUD básico para o domínio `admin`, servindo como exemplo de organização de camadas e boas práticas de arquitetura.

---

## 🧱 Estrutura do Projeto
```bash
.
├── cmd/                    # Ponto de entrada da aplicação
│   ├── configuration/      # Carregamento de configurações
│   └── server/             # Inicialização do servidor HTTP
├── internal/
│   ├── app/                # Camada de aplicação (handlers, serviços, etc.)
│   ├── infrastructure/     # Banco de dados e serviços externos
│   └── configuration/      # Configurações globais
└── main.go                 # Entrada principal
```

## 🚀 Instalação
1. Clone o repositório
   ```bash
   git clone https://github.com/usuario/basic-crud-go.git
   cd basic-crud-go
   ```
2. Copie o arquivo `.env.example` para `.env` e ajuste os valores conforme seu ambiente.
3. Instale as dependências (opcional, caso não sejam baixadas automaticamente):
   ```bash
   go mod download
   ```

## 🔧 Configuração de Variáveis de Ambiente
As variáveis são lidas do arquivo `.env`. Abaixo um exemplo dos principais parâmetros:
```env
ENV=DEV
LISTEN_SERVER=0.0.0.0
HTTP_PORT=8080
HTTPS=FALSE
HTTPS_PORT=8081
DNS="example.dns.org"
LOG_LEVEL=1
LOG_DATABASE=FALSE
RECOVERY_EMAIL="email@example.org"
RECOVERY_PWD="Ex@mpl3PwD"
DATABASE_URL=127.0.0.1
DATABASE_PORT=5432
DATABASE_USER=user_acess
DATABASE_PW=pwExampl3
DATABASE_NAME=users_db
DATABASE_SSL=disable
```

## ▶️ Como Executar
Execute o projeto com o comando abaixo:
```bash
go run main.go
```
A API será exposta em `http://<LISTEN_SERVER>:<HTTP_PORT>` ou HTTPS caso habilitado.

## 📝 Licença
Este projeto está licenciado sob os termos da [MIT License](LICENSE).
