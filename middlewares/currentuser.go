package middlewares

var currentUser string

func SaveCurrentUser(user string) {
	currentUser = user
	println(currentUser)
}

func GetCurrentUser() string {
	return currentUser
}
