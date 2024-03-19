package controller

import (
	m "gin/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertMhs(c *gin.Context) {
	db, err := connectGorm()
	if err != nil {
		SendErrorResponse(c, 500, "internal server error")
		return
	}

	nim := c.Query("nim")
	nama := c.Query("nama")
	umurStr := c.Query("umur")

	umur, err := strconv.Atoi(umurStr)
	if err != nil {
		SendErrorResponse(c, 500, "internal server error")
		return
	}

	mhs := m.Mahasiswa{
		Nim:  nim,
		Nama: nama,
		Umur: umur,
	}

	result := db.Create(&mhs)
	err = result.Error
	if err != nil {
		SendErrorResponse(c, 500, "internal server error")
		return
	}
	c.JSON(http.StatusOK, mhs)
}

func GetMhs(c *gin.Context) {
	db, err := connectGorm()
	if err != nil {
		SendErrorResponse(c, 500, "internal server error") //response nya juga bisa ini buat sendiri
		return
	}

	var mhs []m.Mahasiswa
	queryResult := db.Find(&mhs)
	if queryResult.Error != nil {
		SendErrorResponse(c, 500, "internal server error")
		return
	}

	c.JSON(http.StatusOK, mhs) //ini juga response framework gin
}

func UpdateMhs(c *gin.Context) {
	db, err := connectGorm()
	if err != nil {
		SendErrorResponse(c, 500, "internal server error")
		return
	}

	id := c.Param("id")
	var mhs m.Mahasiswa

	umurStr := c.Query("umur")
	umur, err := strconv.Atoi(umurStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "umur harus berupa angka"})
		return
	}

	result := db.Model(&mhs).Where("id = ?", id).Update("umur", umur)
	err = result.Error
	if err != nil {
		SendErrorResponse(c, 500, "internal server error")
		return
	}
	if result.RowsAffected == 0 {
		SendErrorResponse(c, 404, "Data not found")
		return
	}
	c.JSON(http.StatusOK, mhs)
}

func DeleteMhs(c *gin.Context) {
	db, err := connectGorm()
	if err != nil {
		SendErrorResponse(c, 500, "internal server error")
		return
	}

	id := c.Param("id")
	var mhs m.Mahasiswa

	result := db.Delete(&mhs, &id)
	if result.RowsAffected == 0 {
		SendErrorResponse(c, 500, "internal server error")
		return
	}
	c.JSON(http.StatusOK, mhs)
}
