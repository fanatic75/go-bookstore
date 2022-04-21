package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fanatic75/go-bookstore/pkg/models"
	"github.com/fanatic75/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	res, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Parsing failed")
		return
	}
	book, _ := models.GetBookById(bookId)
	res, er := json.Marshal(book)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	utils.ParseBody(r, &book)
	models.CreateBook(&book)
	res, er := json.Marshal(book)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Failed during parsing book id")
		return
	}
	book := models.DeleteBook(bookId)
	res, er := json.Marshal(book)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
