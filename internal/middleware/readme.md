# Middleware

Este diretório armazena middlewares reutilizáveis para o framework Gin.

## Auth

O middleware `Auth` valida o token enviado no cabeçalho `Authorization` e executa a verificação de permissões para rotas protegidas.

A inicialização deve ocorrer apenas uma vez na aplicação, conforme exemplo abaixo:

```go
// cmd/server/routes.go
import (
    "basic-crud-go/internal/infrastructure/db/postgres"
    middleware "basic-crud-go/internal/middleware"
)

func RegisterRoutes(r *gin.Engine) {
    middleware.SetupDefault(postgres.GetDB())
    enterprise.RegisterEnterpriseRoutes(r)
}
```

Use `mw.AuthMiddleware("codigo-permissao")` nas rotas que necessitam autenticação.
