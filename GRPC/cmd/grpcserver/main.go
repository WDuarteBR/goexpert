package main

import (
	"database/sql"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wduartebr/goexpert/grpc/internal/database"
	"github.com/wduartebr/goexpert/grpc/internal/pb"
	"github.com/wduartebr/goexpert/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()                               // criei um servidor
	pb.RegisterCategorySeviceServer(grpcServer, categoryService) // anexando o servico ao servidor
	reflection.Register(grpcServer)                              // aqui fiz esse refection para realizar testes com evans

	list, err := net.Listen("tcp", ":50051") // criando uma porta tcp
	if err != nil {
		panic(err)
	}

	if err = grpcServer.Serve(list); err != nil {
		panic(err)
	}

}
