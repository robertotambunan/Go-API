package main

import "time"

type Task struct {
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Tasks []Task

type ReturnAPI struct {
	Product []Product `json:"products"`
}
type Product struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	City           string `json:"city"`
	Price          int    `json:"price"`
	Category       string `json:"category"`
	SellerUsername string `json:"seller_username"`
	SellerName     string `json:"seller_name"`
	Province       string `json:"province"`
	Url            string `json:"url"`
	Weight         int    `json:"weight"`
	Stock          int    `json:"stock"`
}
