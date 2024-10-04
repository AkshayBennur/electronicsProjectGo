package services

import (
	"net/http"
	"electronicsProjectGo/models"
	"electronicsProjectGo/repositories"
	// "strconv"
	// "time"
)

type ProductsService struct {
	productsRepository *repositories.ProductsRepository
}

func NewProductsService(productsRepository *repositories.ProductsRepository) *ProductsService {
	return &ProductsService{
		productsRepository: productsRepository,
	}
}

func (rs ProductsService) CreateProduct(product *models.Product) (*models.Product, *models.ResponseError) {
	responseErr := validateProduct(product)
	if responseErr != nil {
		return nil, responseErr
	}

	return rs.productsRepository.CreateProduct(product)
}

func (rs ProductsService) UpdateProduct(productName string) *models.ResponseError {
	responseErr := validateProductParam(productName)
	if responseErr != nil {
		return responseErr
	}

	return rs.productsRepository.UpdateProduct(productName)
}

func (rs ProductsService) DeleteProduct(productId string) *models.ResponseError {
	responseErr := validateProductParam(productId)
	if responseErr != nil {
		return responseErr
	}

	return rs.productsRepository.DeleteProduct(productId)
}

// func (rs ProductsService) GetProduct(productId string) (*models.Product, *models.ResponseError) {
// 	responseErr := validateProductId(productId)
// 	if responseErr != nil {
// 		return nil, responseErr
// 	}

// 	product, responseErr := rs.productsRepository.GetProduct(productId)
// 	if responseErr != nil {
// 		return nil, responseErr
// 	}

// 	results, responseErr := rs.resultsRepository.GetAllProductsResults(productId)
// 	if responseErr != nil {
// 		return nil, responseErr
// 	}

// 	product.Results = results

// 	return product, nil
// }

func (rs ProductsService) GetProductsBatch(country string, year string) ([]*models.Product, *models.ResponseError) {
	// if country != "" && year != "" {
	// 	return nil, &models.ResponseError{
	// 		Message: "Only one parameter, country or year, can be passed",
	// 		Status:  http.StatusBadRequest,
	// 	}
	// }

	// if country != "" {
	// 	return rs.productsRepository.GetProductsByCountry(country)
	// }

	// if year != "" {
	// 	intYear, err := strconv.Atoi(year)
	// 	if err != nil {
	// 		return nil, &models.ResponseError{
	// 			Message: "Invalid year",
	// 			Status:  http.StatusBadRequest,
	// 		}
	// 	}

	// 	currentYear := time.Now().Year()
	// 	if intYear < 0 || intYear > currentYear {
	// 		return nil, &models.ResponseError{
	// 			Message: "Invalid year",
	// 			Status:  http.StatusBadRequest,
	// 		}
	// 	}

	// 	return rs.productsRepository.GetProductsByYear(intYear)
	// }

	return rs.productsRepository.GetAllProducts()
}

func validateProduct(product *models.Product) *models.ResponseError {

	//remove
	// name VARCHAR(255) NOT NULL,
    // price DECIMAL(10, 2) NOT NULL,
    // category VARCHAR(100) NOT NULL,
    // brand VARCHAR(100) NOT NULL,
    // rating INT DEFAULT 0,
    // selected BOOLEAN DEFAULT FALSE,
    // ordered BOOLEAN DEFAULT FALSE

	if product.Name == "" {
		return &models.ResponseError{
			Message: "Invalid Name",
			Status:  http.StatusBadRequest,
		}
	}

	if product.Price == 0 {
		return &models.ResponseError{
			Message: "Invalid Price",
			Status:  http.StatusBadRequest}
	}


	if product.Category == "" {
		return &models.ResponseError{
			Message: "Invalid Category",
			Status:  http.StatusBadRequest,
		}
	}

	if product.Brand == "" {
		return &models.ResponseError{
			Message: "Invalid Brand",
			Status:  http.StatusBadRequest,
		}
	}

	if product.Rating != 0 {
		return &models.ResponseError{
			Message: "Invalid Rating",	
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}

func validateProductParam(productParam string) *models.ResponseError {
	if productParam == "" {
		return &models.ResponseError{
			Message: "Invalid product Name",
			Status:  http.StatusBadRequest,
		}
	}

	return nil
}
