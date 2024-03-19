package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	m "mux/model"

	"github.com/gorilla/mux"
)

func InsertMhsMux(w http.ResponseWriter, r *http.Request) {
	db, err := connectGorm()
	if err != nil {
		SendErrorResponseMux(w, 500, "error")
		return
	}

	nim := r.URL.Query().Get("nim")
	nama := r.URL.Query().Get("nama")
	umurStr := r.URL.Query().Get("umur")

	umur, err := strconv.Atoi(umurStr)
	if err != nil {
		SendErrorResponseMux(w, 400, "umur gagal di convert")
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
		SendErrorResponseMux(w, 404, "data not found")
		return
	}
	SendSuccesResponseMux(w, 200, "success")
}

func GetMhsMux(w http.ResponseWriter, r *http.Request) {
	db, err := connectGorm()
	if err != nil {
		SendErrorResponseMux(w, 500, "Failed to establish a connection to the database")
		return
	}

	var mhs []m.Mahasiswa
	queryResult := db.Find(&mhs)
	if queryResult.Error != nil {
		SendErrorResponseMux(w, 500, "Failed to establish a connection to the database")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var response m.MhsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = mhs
	json.NewEncoder(w).Encode(response)
}

func UpdateMhsMux(w http.ResponseWriter, r *http.Request) {
	db, err := connectGorm()
	if err != nil {
		SendErrorResponseMux(w, 500, "Failed to establish a connection to the database")
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	var mhs m.Mahasiswa

	umurStr := r.URL.Query().Get("umur")
	umur, err := strconv.Atoi(umurStr)
	if err != nil {
		SendErrorResponseMux(w, 500, "Failed to parse age parameter")
		return
	}

	result := db.Model(&mhs).Where("id = ?", id).Update("umur", umur)
	if result.Error != nil {
		SendErrorResponseMux(w, 500, "Failed to update data")
		return
	}
	if result.RowsAffected == 0 {
		SendErrorResponseMux(w, 404, "Data not found")
		return
	}
	SendSuccesResponseMux(w, 200, "success")
}

func DeleteMhsMux(w http.ResponseWriter, r *http.Request) {
	db, err := connectGorm()
	if err != nil {
		SendErrorResponseMux(w, 500, "Failed to establish a connection to the database")
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	var mhs m.Mahasiswa

	result := db.Delete(&mhs, &id)
	if result.RowsAffected == 0 {
		SendErrorResponseMux(w, 500, "Failed to establish a connection to the database")
		return
	}
	SendSuccesResponseMux(w, 200, "succes")
}
