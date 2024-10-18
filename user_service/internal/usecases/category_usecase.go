package usecases

import (
	"auth_service/internal/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type CategoryUC struct {
	httpClient *http.Client
	baseURL    string
}

func NewCategoryUC(client *http.Client, baseURL string) *CategoryUC {
	return &CategoryUC{
		httpClient: client,
		baseURL:    baseURL,
	}
}

func (uc *CategoryUC) CreateDefaultCategories(userID uint) {
	for _, category := range models.DefaultCategories {
		body := models.CreateCategoryRequest{
			UserID: userID,
			Name:   category.Name,
		}

		jsonBody, err := json.Marshal(body)

		if err != nil {
			log.Printf("Failed to create category %v due to error: %v", category.Name, err.Error())
		}

		req, err := http.NewRequest("POST", uc.baseURL+"/category", bytes.NewBuffer(jsonBody))
		if err != nil {
			log.Printf("Failed to create request for %v due to error: %v", category.Name, err.Error())
		}

		resp, err := uc.httpClient.Do(req)
		if err != nil {
			log.Printf("Failed to send request for %v due to error: %v", category.Name, err.Error())
		}
		defer resp.Body.Close()

		var categoryResp models.CreateCategoryResponse

		if err := json.NewDecoder(resp.Body).Decode(&categoryResp); err != nil {
			log.Printf("Failed to unmarchal data for %v due to error: %v", category.Name, err.Error())
		}

		for _, subcategory := range category.SubcategoryList {
			subcategoryBody := models.CreateSubcategoryRequest{
				CategoryID: categoryResp.ID,
				Name:       subcategory.Name,
			}

			jsonSubcategoryBody, err := json.Marshal(subcategoryBody)
			if err != nil {
				log.Printf("Failed to create category %v due to error: %v", subcategory.Name, err.Error())
			}

			req, err := http.NewRequest("POST", uc.baseURL+"/subcategory", bytes.NewBuffer(jsonSubcategoryBody))
			if err != nil {
				log.Printf("Failed to create request for %v due to error: %v", subcategory.Name, err.Error())
			}

			resp, err := uc.httpClient.Do(req)
			if err != nil {
				log.Printf("Failed to send request for %v due to error: %v", subcategory.Name, err.Error())
			}
			defer resp.Body.Close()

		}

	}

}
