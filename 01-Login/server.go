package main

import (
	"github.com/taytorious/cavavin/01-Login/routes/callback"
	"github.com/taytorious/cavavin/01-Login/routes/home"
	"github.com/taytorious/cavavin/01-Login/routes/middlewares"
	"github.com/taytorious/cavavin/01-Login/routes/user"
	"github.com/taytorious/cavavin/01-Login/routes/login"
	"github.com/taytorious/cavavin/01-Login/routes/logout"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	log.Print("Server listening on http://localhost:3000/")
	http.ListenAndServe("0.0.0.0:3000", nil)
}
