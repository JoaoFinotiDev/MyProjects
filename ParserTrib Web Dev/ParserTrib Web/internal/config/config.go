package config

// Config é uma struct que contem as configurações padrão do sistema
type Config struct {
	CaminhoPadrao string
	SheetPadrao   string
	DiretorioLogs string
}

// Nova cria uma instância de Config com valores padrão
func Nova() *Config {
	return &Config{
		CaminhoPadrao: "./xlsxModels",
		SheetPadrao:   "Produto",
		DiretorioLogs: "./logs",
	}
}
