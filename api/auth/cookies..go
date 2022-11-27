package auth

// // If use session cookie - Unused ATM replaced by JWT

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/securecookie"
// )

// var cookieHandler = securecookie.New(
// 	securecookie.GenerateRandomKey(64),
// 	securecookie.GenerateRandomKey(32))

// func GetSession(request *http.Request) (role string, err error) {
// 	cookie, err := request.Cookie("role")
// 	if err != nil {
// 		return "", err
// 	}
// 	cookieValue := make(map[string]string)

// 	err = cookieHandler.Decode("role", cookie.Value, &cookieValue)
// 	if err != nil {
// 		return "", err
// 	}
// 	role = cookieValue["role"]
// 	log.Println("Role get: ", role)
// 	return role, nil
// }

// func SetSession(yourName string, response http.ResponseWriter) error {
// 	value := map[string]string{
// 		"role": yourName,
// 	}

// 	encoded, err := cookieHandler.Encode("role", value)
// 	if err != nil {
// 		return err
// 	}

// 	cookie := &http.Cookie{
// 		Name:   "role",
// 		Value:  encoded,
// 		Path:   "/",
// 		MaxAge: 3600,
// 	}
// 	http.SetCookie(response, cookie)
// 	log.Println("Role set: ", yourName)
// 	return nil
// }
