package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/azad-ali786/bookmanagementsystem/pkg/models"
	"github.com/azad-ali786/bookmanagementsystem/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, err := json.Marshal(b)
	utils.ErrorHandler(err)
	w.Header().Set("Content-Type", "pkglication/json")	
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	utils.ErrorHandler(err)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	utils.ErrorHandler(err)
	bookDetails, _ := models.GetBookById(Id)
	res, err := json.Marshal(bookDetails)
	utils.ErrorHandler(err)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	utils.ErrorHandler(err)
	deletedBook := models.DeleteBook(Id)
	res, err := json.Marshal(deletedBook)
	utils.ErrorHandler(err)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	utils.ErrorHandler(err)
	bookDetails, db := models.GetBookById(Id)
	if bookDetails.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if bookDetails.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if bookDetails.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
    _ = models.DeleteBook(Id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}