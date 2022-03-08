package database

import (
	"github.com/fukunokaze/GoMicro/ItemMasterService/database/model"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

type ItemMasterRepository interface {
	CreateItemMaster(model *model.ItemMaster) *model.ItemMaster
}

type itemMasterRepository struct {
	connection *gorm.DB
}

func NewItemMasterRepository() ItemMasterRepository {
	DB, err = gorm.Open(postgres.Open(DbURL(BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
		fmt.Println("connection error")
	}
	DB.AutoMigrate(&model.ItemMaster{})
	return &itemMasterRepository{
		connection: DB,
	}
}

func (repo *itemMasterRepository) CreateItemMaster(model *model.ItemMaster) *model.ItemMaster {
	repo.connection.Create(&model)
	return model
}
