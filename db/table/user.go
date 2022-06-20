package table

import "log"
import "gorm.io/gorm"

type User struct{
	ID 			string `json:"id",gorm:"primaryKey"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Meta1 		string `json:"metadata1,omitempty"`
	Meta2 		string `json:"-"`
}

func (User) TableName() string {
    return "user"
}

func MigrateUser(db *gorm.DB){	
	if db.Migrator().HasTable(&User{}) == false{	
		log.Println("migrate table user")
		db.AutoMigrate(&User{})
	}
}