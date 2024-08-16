// NOTE: Goではディレクトリがパッケージに対応する
package model

import "time"

type User struct {
	ID        uint 		  `json:"id" gorm:"primary_key"` // NOTE: jsonに変換時に自動的に小文字に変換させる
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
}
