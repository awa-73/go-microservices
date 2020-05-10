package main

//restful services

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shop.com/v1/handlers"
	"time"
	"github.com/go-openapi/runtime/middleware"
)


func main()  {
	l:=log.New(os.Stdout,"product-api", log.LstdFlags)
	//server

	sm:=mux.NewRouter()
	getRouter:=sm.Methods("GET").Subrouter()
	putRouter:=sm.Methods("PUT").Subrouter()
	postRouter:=sm.Methods("POST").Subrouter()
	//handlers
	products:= handlers.NewProducts(l)
	getRouter.HandleFunc("/",products.GetProducts)
	putRouter.HandleFunc("/{id:[0-9]+}", products.UpdateProduct)
	postRouter.Use(products.MiddlewareProductValidation)
	postRouter.HandleFunc("/", products.AddProduct)

	ops:=middleware.RedocOpts{SpecURL:"/swagger.yaml"}
	sh:= middleware.Redoc(ops,nil)
	getRouter.Handle("/docs",sh)
	getRouter.Handle("/swagger.yaml",http.FileServer(http.Dir("./")))
	//server params
	s:=http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		WriteTimeout: 1*time.Second,
		ReadTimeout: 1*time.Second,
	}
	go func() {
		err:=s.ListenAndServe()
		if err!=nil {
			l.Fatal(err)
		}
	}()
	//shutdown the server gracefully
	//disconnect all transactions without losing data or affecting client connections and data
	//in transmission
	sigChan:=make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig:=<-sigChan
	l.Println("Received terminate gracefully shutdown", sig)
	tc,_:=context.WithTimeout(context.Background(), 30*time.Second)
	_ = s.Shutdown(tc)
}