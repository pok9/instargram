package models

//date +%s
type Follower struct {
	FollowingUserID uint `json:"followingUserID"`
	FollowedUserID  uint `json:"followedUserID"`

	//User foreignKey Following_user_id,Followed_user_id
	FollowingUser    User `gorm:"foreignKey:FollowingUserID"`
	UsFollowedUserer User `gorm:"foreignKey:FollowedUserID"`
	// User User `gorm:"foreignKey:FollowingUserID,FollowedUserID"`
}
