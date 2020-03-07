package models

import "time"

type User struct {
	ID        uint       `gorm:"primary_key"`
	UID       string     `json:"-"` //json:"-"でレスポンスに含めないようにする
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index"json:"-"`

	Favorites []Favorite //”1（User)対多(Favorite)”を表現
}
