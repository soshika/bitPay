package user

type User struct {
	Id 			int64 		`json:"id"`
	APIKey		string 		`json:"api_key"`
	SecretKey	string		`json:"secret_key"`
}

func (user *User) Init(apiKey, secretKey string) {
	user.APIKey 	= apiKey
	user.SecretKey 	= secretKey
}


