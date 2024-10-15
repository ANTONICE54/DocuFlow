package models

import "time"

type Subcategory struct {
	ID         uint
	CategoryID uint
	Name       string
	CreatedAt  time.Time
}

type CreateSubcategoryRequest struct {
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
}

type CreateSubcategoryResponse struct {
	ID         uint   `json:"id"`
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
}

type ListSubcategoryRequest struct {
	CategoryID uint `json:"category_id"`
}

type ElementOfSubcategoryList struct {
	ID         uint   `json:"id"`
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
}

type ListSubcategoryResponse struct {
	SubcategoryList []ElementOfSubcategoryList `json:"subcategory_list"`
}

type UpdateSubcategoryRequest struct {
	Name string `json:"name"`
}

type UpdateSubcategoryResponse struct {
	ID         uint   `json:"id"`
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
}
