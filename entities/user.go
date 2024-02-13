package entities

type User struct {
	Id int `json:"-" db:"id"`

	Username string `json:"username" binding:"required"`

	//Password string `json:"password" binding:"required"`
	//
	//Name string `json:"name" binding:"required"`
	//
	//Age string `json:"age" binding:"required"`
	//
	//Weight int `json:"weight" binding:"required"`
	//
	//Height int `json:"height" binding:"required"`
	//
	//Image string `json:"image"`
}
