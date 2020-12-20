package postgres

import (
	"database/sql"
	"time"
)

type Country struct {
	ID    int32  `db:"country_id"`
	Title string `db:"title"`
}

type City struct {
	ID    int32  `db:"city_id"`
	Title string `db:"title"`
}

type University struct {
	ID        int32         `db:"university_id"`
	Name      string        `db:"name"`
	CountryID sql.NullInt32 `db:"country_id"`
	CityID    sql.NullInt32 `db:"city_id"`
}

type School struct {
	ID            int            `db:"school_id"`
	Name          string         `db:"name"`
	YearFrom      sql.NullInt32  `db:"year_from"`
	YearTo        sql.NullInt32  `db:"year_to"`
	YearGraduated sql.NullInt32  `db:"year_graduated"`
	Type          sql.NullString `db:"type_str"`
	CountryID     sql.NullInt32  `db:"country_id"`
	CityID        sql.NullInt32  `db:"city_id"`
}

type Group struct {
	ID           int    `db:"group_id"`
	ScreenName   string `db:"screen_name"`
	Name         string `db:"name"`
	MembersCount int    `db:"members_count"`
	Type         string `db:"type"`
	IsClosed     bool   `db:"is_closed"`
}

type Photo struct {
	ID            int           `db:"photo_id"`
	OwnerID       sql.NullInt32 `db:"owner_id"`
	Date          time.Time     `db:"date"`
	Text          string        `db:"text"`
	LikesCount    int           `db:"likes_count"`
	CommentsCount int           `db:"comments_count"`
	RepostsCount  int           `db:"reposts_count"`
	LikedIDs      []int         `db:"liked_ids"`
	CommentedIDs  []int         `db:"commented_ids"`
}

type Post struct {
	ID            int           `db:"post_id"`
	OwnerID       sql.NullInt32 `db:"owner_id"`
	Date          time.Time     `db:"date"`
	Text          string        `db:"text"`
	LikesCount    int           `db:"likes_count"`
	CommentsCount int           `db:"comments_count"`
	ViewsCount    int           `db:"views_count"`
	RepostsCount  int           `db:"reposts_count"`
	LikedIDs      []int         `db:"liked_ids"`
	CommentedIDs  []int         `db:"commented_ids"`
}

type User struct {
	ID             int           `db:"user_id"`
	FirstName      string        `db:"first_name"`
	LastName       string        `db:"last_name"`
	IsClosed       bool          `db:"is_closed"`
	Sex            int           `db:"sex"`
	Domain         string        `db:"domain"`
	BDate          string        `db:"bdate"`
	CollectDate    time.Time     `db:"collect_date"`
	Status         string        `db:"status"`
	Verified       bool          `db:"verified"`
	CountryID      sql.NullInt32 `db:"country_id"`
	CityID         sql.NullInt32 `db:"city_id"`
	HomeTown       string        `db:"home_town"`
	Universities   []int32       `db:"universities"`
	Schools        []int         `db:"schools"`
	FriendsCount   int           `db:"friends_count"`
	FriendsIDs     []int         `db:"friends_ids"`
	FollowersCount int           `db:"followers_count"`
	FollowersIDs   []int         `db:"followers_ids"`
	PostsCount     int           `db:"posts_count"`
	PostsIDs       []int         `db:"posts_ids"`
	PhotosCount    int           `db:"photos_count"`
	PhotosIDs      []int         `db:"photos_ids"`
	GroupsCount    int           `db:"groups_count"`
	GroupsIDs      []int         `db:"groups_ids"`
}
