package models

type Follower struct {
	Id         uint `gorm:"primaryKey"`
	UserId     uint `json:"user_id"`
	FollowerId uint `json:"follower_id"`
	User       User `gorm:"foreignKey:UserId"`
	Follower   User `gorm:"foreignKey:FollowerId"`
}
