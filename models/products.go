package models

import "time"

type Product struct {
	TitleFa            string    `json:"title_fa"`
	TitleEn            string    `json:"title_en"`
	Description        string    `json:"description"`
	Specifications     string    `json:"specifications"`
	Score              string    `json:"score"`
	WarrantyDetails    string    `json:"warranty_details"`
	AvailabilityStatus string    `json:"availability_status"`
	Brand              string    `json:"brand"`
	Color              string    `json:"color"`
	InStockCount       string    `json:"in_stock_count"`
	MainCategory       string    `json:"main_category"`
	SubCategory        string    `json:"sub_category"`
	Link               string    `json:"link"`
	EncodedLink        string    `json:"encoded_link"`
	ImageLink          string    `json:"image_link"`
	PreviousPrice      string    `json:"previous-price"`
	CurrentPrice       string    `json:"current-price"`
	Discount           string    `json:"discount"`
	SellerName         string    `json:"seller_name"`
	SellerScore        string    `json:"seller_score"`
	Time               time.Time `json:"time"`
}
