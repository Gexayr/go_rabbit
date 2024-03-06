// models/user.go
package models

// User представляет структуру данных пользователя.
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    // Другие поля, связанные с пользователем, могут быть добавлены сюда
}
