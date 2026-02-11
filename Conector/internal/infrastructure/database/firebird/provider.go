package firebird

import (
	"Conector/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

type FirebirdProvider struct {
	db *sql.DB
}

func NewFirebirdProvider(cfg config.FirebirdConfig) (*FirebirdProvider, error) {
	ConnectString := fmt.Sprintf("%s:%s@%s:%d/%s?charset=%s&auth_plugin_name=Srp",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
		cfg.Charset,
	)

	db, err := sql.Open("firebirdsql", ConnectString)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro no ping: %w", err)
	}

	return &FirebirdProvider{db}, nil
}

func (p *FirebirdProvider) Close() error {
	return p.db.Close()
}
