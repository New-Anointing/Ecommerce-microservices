package product

import (
	"net/http"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type accessData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *accessData {
	return &accessData{DB: db}
}

// var newProduct []Products

func (r *accessData) createProduct(context *fiber.Ctx) error {
	prod := Products{}
	err := context.BodyParser(&prod)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"Message": "request failed"},
		)
		return err
	}
	r.DB.Create(&prod)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"Message": "Sorry, Could not create book"},
		)
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"Message": "Product successfully created",
	})
	return nil

}

func (r *accessData) GetAllProduct(context *fiber.Ctx) {
	products := &[]Products{}
	err := r.DB.Find(products).Error
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"Message": "Sorry, could not fetch product",
		})
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"Message": "Product fetched successfully",
		"Data":    products,
	})

}

func (r *accessData) GetProductByID(context *fiber.Ctx) {
	var product Products
	id := context.Params("Id")
	err := r.DB.First(&product, id).Error
	if err != nil { // Fixed the if condition
		context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"Message": "Sorry, there is no data matching the id",
		})
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"Message": "Product fetched successfully",
		"Product": product,
	})

}

func (r *accessData) DeleteProductbyID(context *fiber.Ctx) {
	var product Products
	id := context.Params("Id")
	err := r.DB.Delete(&product, id).Error
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"Message": "Could not delete product",
		})
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"Message": "Product deleted successfully",
	})

}

func (r *accessData) UpdateProductByID(context *fiber.Ctx) {
	id := context.Params("Id")
	var updateProduct Products
	err := r.DB.First(&updateProduct, id).Error
	if err != nil {
		context.Status(http.StatusNotFound).JSON(&fiber.Map{
			"Message": "Could not find product",
		})
	}
	if err := context.BodyParser(&updateProduct); err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"Message": "Could not update",
		})
	}
	if err := r.DB.Save(&updateProduct).Error; err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"Message": "Could not save to db",
		})
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"Message": "Successfully saved",
	})

}
