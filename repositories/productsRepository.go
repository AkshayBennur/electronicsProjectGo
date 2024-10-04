package repositories

import (
	"database/sql"
	// "fmt"
	"net/http"
	"runners-mysql/models"
	"strconv"
)

type ProductsRepository struct {
	dbHandler   *sql.DB
}

func NewProductsRepository(dbHandler *sql.DB) *ProductsRepository {
	return &ProductsRepository{
		dbHandler: dbHandler,
	}
}

func (rr ProductsRepository) CreateProduct(product *models.Product) (*models.Product, *models.ResponseError) {
	// INSERT INTO products (name, price, category, brand, rating, selected, ordered) VALUES 
	query := `
		INSERT INTO products(name, price, category, brand, rating, selected, ordered)
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	res, err := rr.dbHandler.Exec(query, product.Name, product.Price, product.Category, product.Brand, product.Rating, product.Selected, product.Ordered)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	productId, err := res.LastInsertId()
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	// product.Name, product.Price, product.Category, product.Brand, product.Rating, product.Selected, product.Ordered

	// ID			string	`json:"id"`
    // Name  		string	`json:"name"`
    // Price		int 	`json:"price"`
    // Category	string	`json:"category"`
    // Brand		string	`json:"brand"`
    // Rating		int		`json:"rating"`
    // Selected	bool	`json:"selected"`
    // Ordered		bool	`json:"ordered"`


	return &models.Product{
		ID:        strconv.FormatInt(productId, 10),
		Name: product.Name,
		Price:  product.Price,
		Category:       product.Category,
		Brand:       product.Brand,
		Rating:       product.Rating,
		Selected:       product.Selected,
		Ordered:       product.Ordered,
	}, nil
}

func (rr ProductsRepository) UpdateProduct(productName string) *models.ResponseError {
// 	UPDATE products
// SET selected = TRUE
// WHERE name = ?

// 	query := `
// 		UPDATE products
// 		SET
// 			first_name = ?,
// 			last_name = ?,
// 			age = ?,
// 			country = ?
// 		WHERE id = ?`

		query := `
			UPDATE products
			SET selected = TRUE
			WHERE name = ?`
		

	res, err := rr.dbHandler.Exec(query, productName)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if rowsAffected == 0 {
		return &models.ResponseError{
			Message: "Product not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}

// func (rr ProductsRepository) UpdateProductResults(product *models.Product) *models.ResponseError {
// 	query := `
// 		UPDATE products
// 		SET
// 			personal_best = ?,
// 			season_best = ?
// 		WHERE id = ?`

// 	_, err := rr.transaction.Exec(query, product.PersonalBest, product.SeasonBest, product.ID)
// 	if err != nil {
// 		return &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	return nil
// }

func (rr ProductsRepository) DeleteProduct(productId string) *models.ResponseError {
	query := `DELETE FROM products WHERE id = ?`

	res, err := rr.dbHandler.Exec(query, productId)
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if rowsAffected == 0 {
		return &models.ResponseError{
			Message: "Product not found",
			Status:  http.StatusNotFound,
		}
	}

	return nil
}

// func (rr ProductsRepository) GetProduct(productId string) (*models.Product, *models.ResponseError) {
// 	fmt.Println(productId)
// 	query := `
// 		SELECT *
// 		FROM products
// 		WHERE id = ?`

// 	rows, err := rr.dbHandler.Query(query, productId)
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	defer rows.Close()

// 	var id, firstName, lastName, country string
// 	var personalBest, seasonBest sql.NullString
// 	var age int
// 	var isActive bool
// 	for rows.Next() {
// 		err := rows.Scan(&id, &firstName, &lastName, &age, &isActive, &country, &personalBest, &seasonBest)
// 		if err != nil {
// 			return nil, &models.ResponseError{
// 				Message: err.Error(),
// 				Status:  http.StatusInternalServerError,
// 			}
// 		}
// 	}

// 	if rows.Err() != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	return &models.Product{
// 		ID:           id,
// 		FirstName:    firstName,
// 		LastName:     lastName,
// 		Age:          age,
// 		IsActive:     isActive,
// 		Country:      country,
// 		PersonalBest: personalBest.String,
// 		SeasonBest:   seasonBest.String,
// 	}, nil
// }

func (rr ProductsRepository) GetAllProducts() ([]*models.Product, *models.ResponseError) {
	query := `
	SELECT *
	FROM products`

	rows, err := rr.dbHandler.Query(query)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	defer rows.Close()

	// type Product struct {
		// ID			int64	`json:"id"`
		// Name  		string	`json:"name"`
		// Price		float32 `json:"Price"`
		// Category	string	`json:"category"`
		// brand		string	`json:"brand"`
		// Rating		float32	`json:"rating"`
		// Selected	bool	`json:"selected"`
		// Ordered		bool	`json:"ordered"`
	// }

	products := make([]*models.Product, 0)
	var name, category, brand string
	// var personalBest, seasonBest sql.NullString
	var id string
	var price, rating int
	var selected, ordered bool

	for rows.Next() {
		err := rows.Scan(&id, &name, &price, &category, &brand, &rating, &selected, &ordered)
		if err != nil {
			return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}

		product := &models.Product{
			// ID:           id,
			// FirstName:    firstName,
			// LastName:     lastName,
			// Age:          age,
			// IsActive:     isActive,
			// Country:      country,
			// PersonalBest: personalBest.String,
			// SeasonBest:   seasonBest.String,


			ID:			id,
			Name:  		name,
			Price:		price,
			Category:	category,
			Brand:		brand,
			Rating:		rating,
			Selected:	selected,
			Ordered:	ordered,
		}

		products = append(products, product)
	}

	if rows.Err() != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return products, nil
}

// func (rr ProductsRepository) GetProductsByCountry(country string) ([]*models.Product, *models.ResponseError) {
// 	query := `
// 	SELECT id, first_name, last_name, age, personal_best, season_best
// 	FROM products
// 	WHERE country = ? AND is_active = TRUE
// 	ORDER BY personal_best
// 	LIMIT 10`

// 	rows, err := rr.dbHandler.Query(query, country)
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	defer rows.Close()

// 	products := make([]*models.Product, 0)
// 	var id, firstName, lastName string
// 	var personalBest, seasonBest sql.NullString
// 	var age int

// 	for rows.Next() {
// 		err := rows.Scan(&id, &firstName, &lastName, &age, &personalBest, &seasonBest)
// 		if err != nil {
// 			return nil, &models.ResponseError{
// 				Message: err.Error(),
// 				Status:  http.StatusInternalServerError,
// 			}
// 		}

// 		product := &models.Product{
// 			ID:           id,
// 			FirstName:    firstName,
// 			LastName:     lastName,
// 			Age:          age,
// 			IsActive:     true,
// 			Country:      country,
// 			PersonalBest: personalBest.String,
// 			SeasonBest:   seasonBest.String,
// 		}

// 		products = append(products, product)
// 	}

// 	if rows.Err() != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	return products, nil
// }

// func (rr ProductsRepository) GetProductsByYear(year int) ([]*models.Product, *models.ResponseError) {
// 	query := `
// 	SELECT products.id, products.first_name, products.last_name, products.age, products.is_active, products.country, products.personal_best, results.race_result
// 	FROM products
// 	INNER JOIN (
// 		SELECT product_id, MIN(race_result) as race_result
// 		FROM results
// 		WHERE result_year = ?
// 		GROUP BY product_id) results
// 	ON products.id = results.product_id
// 	ORDER BY results.race_result
// 	LIMIT 10`

// 	rows, err := rr.dbHandler.Query(query, year)
// 	if err != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	defer rows.Close()

// 	products := make([]*models.Product, 0)
// 	var id, firstName, lastName, country string
// 	var personalBest, seasonBest sql.NullString
// 	var age int
// 	var isActive bool

// 	for rows.Next() {
// 		err := rows.Scan(&id, &firstName, &lastName, &age, &isActive, &country, &personalBest, &seasonBest)
// 		if err != nil {
// 			return nil, &models.ResponseError{
// 				Message: err.Error(),
// 				Status:  http.StatusInternalServerError,
// 			}
// 		}

// 		product := &models.Product{
// 			ID:           id,
// 			FirstName:    firstName,
// 			LastName:     lastName,
// 			Age:          age,
// 			IsActive:     isActive,
// 			Country:      country,
// 			PersonalBest: personalBest.String,
// 			SeasonBest:   seasonBest.String,
// 		}

// 		products = append(products, product)
// 	}

// 	if rows.Err() != nil {
// 		return nil, &models.ResponseError{
// 			Message: err.Error(),
// 			Status:  http.StatusInternalServerError,
// 		}
// 	}

// 	return products, nil
// }
