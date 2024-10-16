package models

import "time"

type Category struct {
	ID              uint
	UserID          uint
	Name            string
	SubcategoryList []Subcategory
	CreatedAt       time.Time
}

type CreateCategoryRequest struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

type CreateCategoryResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateCategoryResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

type ListCategoryRequest struct {
	UserID uint `json:"user_id"`
}

type ElementOfCategoryList struct {
	ID              uint                       `json:"id"`
	UserID          uint                       `json:"user_id"`
	Name            string                     `json:"name"`
	SubcategoryList []ElementOfSubcategoryList `json:"subcategory_list"`
}

type ListCategoryResponse struct {
	CategoryList []ElementOfCategoryList `json:"category_list"`
}
