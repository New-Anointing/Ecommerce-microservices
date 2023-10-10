package db

import "ecommerce-microservices/product"

func migrate() error {
	err := DB.AutoMigrate(
		&product.Products{},
	)
	return err
}
