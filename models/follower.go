package models

//date +%s
type Follower struct {
	// Model
	FollowingUserID string `json:"followingUserID" gorm:"primary_key"` //เราติดตาม
	FollowedUserID  string `json:"followedUserID" gorm:"primary_key"`  //กำลังติดตามใตร

	//User foreignKey Following_user_id,Followed_user_id
	FollowingUser User `gorm:"foreignKey:FollowingUserID"`
	FollowedUser  User `gorm:"foreignKey:FollowedUserID"`
	// User User `gorm:"foreignKey:FollowingUserID,FollowedUserID"`
}
