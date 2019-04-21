package register

import (
	"database/sql"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	return string(bytes), err
}

func RegisterNewPloyer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ployer/register" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	db, err := sql.Open("mysql", "quickploy:quickploy2019@tcp(awsdbs.cqfpsvdee72a.sa-east-1.rds.amazonaws.com)/Quickploy")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	switch r.Method {
	case "POST":
		email := r.FormValue("email")
		senha := r.FormValue("senha")
		senha, _ = HashPassword(senha)
		tx, _ := db.Begin()
		stmt, _ := tx.Prepare("INSERT INTO ployers(email_ployer, senha_ployer) VALUES(?,?)")
		_, erro := stmt.Exec(email, senha)
		if erro != nil {
			tx.Rollback()
			log.Fatal(erro)
		}
		tx.Commit()
	}
	http.Redirect(w, r, "/ployer/complete", http.StatusSeeOther)
}
