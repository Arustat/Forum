package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Connexion à la base de données
	var err error
	db, err = sql.Open("mysql", "utilisateur:motdepasse@tcp(localhost:3306)/basededonnees")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/creer-compte", CreerCompteHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CreerCompteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Récupération des données du formulaire
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusBadRequest)
		return
	}

	nom := r.Form.Get("nom")
	email := r.Form.Get("email")

	// Insertion des données dans la base de données
	_, err = db.Exec("INSERT INTO utilisateurs (nom, email) VALUES (?, ?)", nom, email)
	if err != nil {
		http.Error(w, "Erreur lors de l'insertion des données dans la base de données", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Compte créé avec succès pour %s", nom)
}

//aled
