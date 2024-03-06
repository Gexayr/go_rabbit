// models/game.go
package models

// Game представляет структуру данных игры.
type Game struct {
    ID       int    `json:"id"`
    State    string `json:"state"`
    Win      float64 `json:"win"`
    // Другие поля, связанные с игрой, могут быть добавлены сюда
}