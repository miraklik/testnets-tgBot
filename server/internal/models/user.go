package models

type User struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

type ChangePassword struct {
	Username    string `bson:"username" json:"username"`
	OldPassword string `bson:"password" json:"old_password"`
	NewPassword string `bson:"password" json:"new_password"`
}
