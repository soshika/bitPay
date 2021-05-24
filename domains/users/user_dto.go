package users

type User struct {
	Id 			int64 		`json:"id"`
	APIKey		string 		`json:"api_key"`
	SecretKey	string		`json:"secret_key"`
}


