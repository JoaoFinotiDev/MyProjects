# FireGopher Conector 🔥🔗

Conector para Firebird (ERP legado) e MongoDB (cloud), construído em Go seguindo princípios de Clean Architecture.

## 📋 Sobre o Projeto e Objetivos

FireGopher Conector é uma aplicação que automatiza a extração de dados de tabelas específicas do Firebird SQL e replica no MongoDB, buscando eliminar a necessidade de uso do Master para criar integrações de API para clientes da Alternativa.

## 🏗️ Arquitetura

Projeto segue **Clean Architecture** com separação clara de responsabilidades:
```
FireGopher/Conector/
├── cmd/api/                    # Entry point da aplicação
│   └── main.go
├── internal/
│   ├── config/                 # Configurações e validação de env vars
│   │   └── config.go
│   ├── application/
│   │   ├── port/              # Interfaces (contratos)
│   │   │   ├── FirebirdReader.go
│   │   │   └── MongoWriter.go
│   │   └── usecase/           # Casos de uso (lógica de negócio)
│   │       └── SyncFirebirdToMongo.go
│   └── infrastructure/
│       └── database/          # Implementações concretas
│           ├── firebird/
│           │   └── provider.go
│           └── mongo/
│               └── provider.go
└── scripts/                    # Utilitários de teste
    ├── Firebird_test_connect.go
    └── Mongo_test_connect.go
```

### Camadas

- **Domain**: Regras de negócio puras (ainda não implementado)
- **Application**: Casos de uso e interfaces (ports)
- **Infrastructure**: Implementações técnicas (database providers)
- **Presentation**: Entry points e APIs (cmd/api)

## 🚀 Tecnologias

- **Go** 1.25.4
- **Firebird SQL** 3.0+ (driver: `nakagami/firebirdsql` v0.9.15)
- **MongoDB** Atlas (driver: `mongo-driver` v1.17.6)
- **godotenv** v1.5.1 (gerenciamento de variáveis de ambiente)

## 🧪 Testando Conexões

### Testar Firebird
```bash
go run internal/scripts/Firebird_test_connect.go
```

Testa 4 variações de DSN (Srp, Legacy, case-sensitive, charset).

### Testar MongoDB
```bash
go run internal/scripts/Mongo_test_connect.go
```

Valida conexão com MongoDB usando credenciais do `.env`.

## ▶️ Executando
```bash
# Rodar aplicação
go run cmd/api/main.go

# Build para produção
go build -o conector cmd/api/main.go
./conector
```

**Output esperado:**
```
2026/01/08 12:00:00 Conectando Firebird em localhost:3050...
2026/01/08 12:00:00 ✅ | Conexão: FirebirdDB
2026/01/08 12:00:00 ✅ | Conexão: MongoDB
```


## 📚 Próximos Passos

- [ ] Implementar interface `FirebirdReader` (ReadTable, ReadTableWhere)
- [ ] Implementar interface `MongoWriter` (WriteDocuments, UpsertDocument)
- [ ] Criar UseCase `SyncFirebirdToMongo` com lógica de sincronização
- [ ] Adicionar testes unitários para providers
- [ ] Definir estratégia de sincronização (full/incremental/upsert)
- [ ] Implementar API REST para disparar sincronizações
- [ ] Adicionar logs estruturados
- [ ] Implementar health checks
- [ ] Criar Docker container


## 📝 Licença

Está restrito ao uso dos colaboradores deste repositorio

## 👤 Autor

**João Vitor**
- Estudante de Ciência da Computação
- Estagiário em Alternativa Sistemas

---

**Status do Projeto**: 🟡 Em Desenvolvimento

Última atualização: 08/01/2026