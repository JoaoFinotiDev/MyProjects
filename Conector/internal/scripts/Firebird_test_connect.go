package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/nakagami/firebirdsql"
)

func main() {
	// Teste com auth_plugin_name=Srp
	dsn1 := "sysdba:masterkey@localhost:3050/c:/asr/master/banco/MOSHE.fdb?auth_plugin_name=Srp"
	testConnection("Teste 1 (Srp)", dsn1)

	// Teste com auth_plugin_name=Legacy_Auth
	dsn2 := "sysdba:masterkey@localhost:3050/c:/asr/master/banco/MOSHE.fdb?auth_plugin_name=Legacy_Auth"
	testConnection("Teste 2 (Legacy)", dsn2)

	// Teste usuário maiúsculo + Srp
	dsn3 := "SYSDBA:masterkey@localhost:3050/c:/asr/master/banco/MOSHE.fdb?auth_plugin_name=Srp"
	testConnection("Teste 3 (SYSDBA + Srp)", dsn3)

	// Teste charset + Srp
	dsn4 := "sysdba:masterkey@localhost:3050/c:/asr/master/banco/MOSHE.fdb?charset=UTF8&auth_plugin_name=Srp"
	testConnection("Teste 4 (charset + Srp)", dsn4)
}

func testConnection(name, dsn string) {
	fmt.Printf("\n=== %s ===\n", name)
	fmt.Printf("DSN: %s\n", dsn)

	db, err := sql.Open("firebirdsql", dsn)
	if err != nil {
		log.Printf("❌ Open falhou: %v\n", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Printf("❌ Ping falhou: %v\n", err)
		return
	}

	fmt.Println("✅ SUCESSO!")
}
