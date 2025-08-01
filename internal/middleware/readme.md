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
    authMW := middleware.NewAuthMiddleware(postgres.GetDB())
    enterprise.RegisterEnterpriseRoutes(r, authMW)
}
```

Use `authMW.Handler("codigo-permissao")` nas rotas que necessitam autenticação.
