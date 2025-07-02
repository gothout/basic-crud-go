# BASIC-CRUD-GO

Este é um projeto de portfólio desenvolvido em **Golang**, utilizando os princípios de **Domain-Driven Design (DDD)**.  
A aplicação implementa um CRUD completo para o domínio `admin`, com foco em organização de camadas, boas práticas de arquitetura e extensibilidade para sistemas reais.

---

## 🧱 Estrutura do Projeto

```bash
.
├── cmd/                    # Ponto de entrada da aplicação
│   ├── configuration/      # Carregamento de configurações
│   └── server/             # Inicialização do servidor HTTP
├── internal/
│   ├── app/
│   │   ├── admin/          # Lógica da aplicação (CRUD Admin)
│   │   ├── configuration/  # Configuração aplicada no contexto da app
│   │   └── middleware/     # Middlewares da aplicação
│   ├── infrastructure/     # Banco de dados, serviços externos
│   └── configuration/      # Configs globais da aplicação
└── main.go                 # Entrada alternativa ou atalho
