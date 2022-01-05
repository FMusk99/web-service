/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"webserver/app/config"
	"webserver/app/entity"
	"webserver/app/infrastructure/repository"
	"webserver/app/usecase/book"
	pb "webserver/grpc"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type Reader interface {
	GetBook(id entity.ID) (*entity.Book, error)
	SearchBooks(query string) ([]*entity.Book, error)
	ListBooks() ([]*entity.Book, error)
}

//Writer book writer
type Writer interface {
	CreateBook(title string, author string, pages int, quantity int) (entity.ID, error)
	UpdateBook(e *entity.Book) error
	DeleteBook(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedBookServer
	book Repository
}

func Newserver(r Repository) *server {
	return &server{
		book: r,
	}
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetBook(ctx context.Context, in *pb.BookParams) (*pb.BookResponse, error) {
	log.Printf("Received: %v", in.GetKeyword())
	return &pb.BookResponse{Message: "Hello " + in.GetKeyword()}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) CreateBook(ctx context.Context, in *pb.BookRequest) (*pb.BookResponse, error) {
	log.Printf("Received: %v", in.GetTitle())
	rs, err := s.book.CreateBook(in.GetTitle(), in.GetAuthor(), int(in.GetPages()), int(in.GetQuantity()))
	if err != nil {
		panic(err)
	}
	fmt.Print("Resulted query: ", rs)
	return &pb.BookResponse{Message: "Created a book successfully."}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) SearchBooks(ctx context.Context, in *pb.BookParams) (*pb.BookResponse, error) {
	log.Printf("Received: %v", in.GetKeyword())
	return &pb.BookResponse{Message: "Hello " + in.GetKeyword()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	dataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_DATABASE)
	db, err := sql.Open("postgres", dataSourceName)
	fmt.Print(dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully created connection to database")
	if err != nil {
		log.Fatal(err.Error())
	}
	// defer db.Close()
	repo := repository.NewBookSQL(db)
	bookService := book.NewService(repo)
	pb.RegisterBookServer(s, Newserver(bookService))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
