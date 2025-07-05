package handler

import (
	"net/http"
	"fmt"
)

type Order struct{

}

func (o *Order) Create(w http.ResponseWriter, r *http.Request){
	fmt.Println("Creating order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request){
	fmt.Println("Listing all orders")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Getting order by ID")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Updating order by ID")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete order by ID")
}