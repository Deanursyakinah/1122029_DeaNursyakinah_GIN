package model

type Mahasiswa struct {
	Id   int    `json:"id"`
	Nim  string `json:"nim"`
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type MhsResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []Mahasiswa `json:"data"`
}
