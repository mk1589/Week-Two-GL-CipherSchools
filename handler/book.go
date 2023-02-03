package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mk1589/Week-Two-GL-CipherSchools/database"
	"github.com/mk1589/Week-Two-GL-CipherSchools/models"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) GetBooks(in *gin.Context) {
	books, err := database.GetBooks(h.DB) // h.DB is full configured and can give the access of book table (collect the book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, books) // return the book
}
func (h *Handler) SaveBook(in *gin.Context) {
	book := models.Book{}
	err := in.BindJSON(&book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SaveBookByID(h.DB, &book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)
}

/*
h:= Handler{}
h.DB=GetDB()
h.GetBooks()
*/
