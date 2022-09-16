package Infor

type Admins struct {
	Useradmin     string `json:"useradmin" binding:"required" gorm : "primary key"`
	Passwordadmin string `json:"passwordadmin" binding:"required"`
}
