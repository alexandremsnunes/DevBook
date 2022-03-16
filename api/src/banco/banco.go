package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

//Abre a conexao com o banco
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.BdConectionStr)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
