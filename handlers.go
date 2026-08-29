package main

import (
    "database/sql"
    "html/template"
    "net/http"

    _ "modernc.org/sqlite"
)

func home(w http.ResponseWriter, r *http.Request) {
    ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
    if err != nil {
        logger.Error("could not parse home template")
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    err = ts.Execute(w, nil)
    if err != nil {
        logger.Error(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

type Dream struct {
    Description string
}

func dreamsView(w http.ResponseWriter, r *http.Request) {
    dbName := r.PathValue("dbName")
    description := r.PathValue("description")

    ts, err := template.ParseFiles("./ui/html/pages/dreams.tmpl")
    if err != nil {
        logger.Error("could not parse dreams template")
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // TODO: db is opened and closed on every HTTP request
    db, err := sql.Open("sqlite", "data/" + dbName + ".sqlite3")
    if err != nil {
        logger.Error("could not open database file")
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS dream (
            description text NOT NULL
        )`)
    if err != nil {
        logger.Error(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    _, err = db.Exec(`
        INSERT INTO dream (description) VALUES ($1)
    `, description)

    err = ts.Execute(w, Dream{dbName})
    if err != nil {
        logger.Error(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
