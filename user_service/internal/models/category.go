package models

type Category struct {
	Name            string
	SubcategoryList []Subcategory
}

type Subcategory struct {
	Name string
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

type CreateSubcategoryRequest struct {
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
}

type CreateSubcategoryResponse struct {
	ID         uint   `json:"id"`
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
}

var DefaultCategories = []Category{
	{
		Name: "Договори",
		SubcategoryList: []Subcategory{
			{
				Name: "Договір поставки",
			},
			{
				Name: "Договір послуг",
			},
			{
				Name: "Ліцензійний договір",
			},
		},
	},
	{
		Name: "Рахунки",
		SubcategoryList: []Subcategory{
			{
				Name: "Рахунок-фактура",
			},
			{
				Name: "Операційні рахунки",
			},
		},
	},
	{
		Name: "Видаткові",
		SubcategoryList: []Subcategory{
			{
				Name: "Видаткові накладної",
			},
			{
				Name: "Актив виконаних робіт",
			},
		},
	},
}
