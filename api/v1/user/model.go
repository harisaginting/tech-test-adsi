package user

type ResponseList struct{
	Items 	[]User 		`json:"items"`
	Total 	int 		`json:"total"`
}

type User struct{
	ID 			string `json:"id",gorm:"primaryKey"`
	FirstName 	string `json:"first_name"`
	Meta1 		string `json:"metadata1,omitempty"`
	Meta2 		string `json:"-"`
}