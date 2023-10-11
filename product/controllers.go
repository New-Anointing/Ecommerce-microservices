package product

import (
	"net/http"
	//"strconv"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type accessData struct {
	DB *gorm.DB
}

// var newProduct []Products

func (r *accessData) createProduct(context *fiber.Ctx) error {
	prod := Products{}
	err := context.BodyParser(&prod)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"Message": "request failed"},
		)
	}
	// w.Header().Set("Content-Type", "application/json")
	// var prod Products
	// _ = json.NewDecoder(r.Body).Decode(&prod)
	// prod.Id = strconv.Itoa(rand.Intn(100000000))
	// newProduct = append(newProduct, prod)
	// json.NewEncoder(w).Encode(prod)

}
