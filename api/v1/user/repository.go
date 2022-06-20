package user

import "log"
import "gorm.io/gorm"
import "github.com/harisaginting/ginting/db/table"
import "github.com/harisaginting/ginting/pkg/utils/helper"

type Repository struct {
	db *gorm.DB
}

func ProviderRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (repo *Repository) FindAll() (users []User) {
	var table []table.User 
	qx 		:= repo.db
	qx.Find(&table)
	if qx.Error != nil {
		log.Println("FindAllByCustomer: ", qx.Error)
	}
	helper.AdjustStructToStruct(table,&users)
	return
}