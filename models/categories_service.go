package models

type Barang struct {
	ID        		int    `json:"id"`
	CategoryName	string `json:"category_name"`
	Description		string `json:"description"`
}