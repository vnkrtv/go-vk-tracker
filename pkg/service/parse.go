package service

import (
	"database/sql"
	"github.com/lib/pq"
	"time"

	pg "github.com/vnkrtv/go-vk-tracker/pkg/postgres"
	vk "github.com/vnkrtv/go-vk-tracker/pkg/vkapi"
)

func parseVKCountry(country vk.VKCountry) pg.Country {
	return pg.Country{
		ID:    country.ID,
		Title: country.Title,
	}
}

func parseVKCity(city vk.VKCity) pg.City {
	return pg.City{
		ID:    city.ID,
		Title: city.Title,
	}
}

func parseVKUniversity(vkUniversity vk.VKUniversity) (pg.University, pg.Country, pg.City) {
	country := pg.Country{
		ID:    vkUniversity.CountryID,
		Title: "",
	}
	city := pg.City{
		ID:    vkUniversity.CityID,
		Title: "",
	}
	university := pg.University{
		ID:        vkUniversity.ID,
		Name:      vkUniversity.Name,
		CountryID: sql.NullInt32{Int32: vkUniversity.CountryID, Valid: true},
		CityID:    sql.NullInt32{Int32: vkUniversity.CityID, Valid: true},
	}
	return university, country, city
}

func parseVKUniversities(vkUniversities []vk.VKUniversity) ([]pg.University, []pg.Country, []pg.City) {
	universities := make([]pg.University, len(vkUniversities))
	countries := make([]pg.Country, len(vkUniversities))
	cities := make([]pg.City, len(vkUniversities))
	for i := range universities {
		university, country, city := parseVKUniversity(vkUniversities[i])
		universities[i] = university
		countries[i] = country
		cities[i] = city
	}
	return universities, countries, cities
}

func parseVKSchool(vkSchool vk.VKSchool) (pg.School, pg.Country, pg.City) {
	country := pg.Country{
		ID:    vkSchool.CountryID,
		Title: "",
	}
	city := pg.City{
		ID:    vkSchool.CityID,
		Title: "",
	}
	school := pg.School{
		ID:            vkSchool.ID,
		Name:          vkSchool.Name,
		YearFrom:      sql.NullInt32{Int32: vkSchool.YearFrom, Valid: true},
		YearTo:        sql.NullInt32{Int32: vkSchool.YearTo, Valid: true},
		YearGraduated: sql.NullInt32{Int32: vkSchool.YearGraduated, Valid: true},
		Type:          sql.NullString{String: vkSchool.Type, Valid: true},
		CountryID:     sql.NullInt32{Int32: vkSchool.CountryID, Valid: true},
		CityID:        sql.NullInt32{Int32: vkSchool.CityID, Valid: true},
	}
	return school, country, city
}

func parseVKSchools(vkSchools []vk.VKSchool) ([]pg.School, []pg.Country, []pg.City) {
	schools := make([]pg.School, len(vkSchools))
	countries := make([]pg.Country, len(vkSchools))
	cities := make([]pg.City, len(vkSchools))
	for i := range schools {
		school, country, city := parseVKSchool(vkSchools[i])
		schools[i] = school
		countries[i] = country
		cities[i] = city
	}
	return schools, countries, cities
}

func parseVKGroup(group vk.VKGroup) pg.Group {
	return pg.Group{
		ID:           group.ID,
		ScreenName:   group.ScreenName,
		Name:         group.Name,
		MembersCount: group.MembersCount,
		Type:         group.Type,
		IsClosed:     group.IsClosed,
	}
}

func parseVKGroups(vkGroups []vk.VKGroup) []pg.Group {
	groups := make([]pg.Group, len(vkGroups))
	for i := range groups {
		groups[i] = parseVKGroup(vkGroups[i])
	}
	return groups
}

func parseVKPost(post vk.VKPost) pg.Post {
	return pg.Post{
		ID:            post.ID,
		OwnerID:       sql.NullInt32{Int32: post.OwnerID, Valid: true},
		Date:          time.Unix(post.Date, 0),
		Text:          post.Text,
		LikesCount:    post.Likes.Count,
		CommentsCount: post.Comments.Count,
		ViewsCount:    post.Views.Count,
		RepostsCount:  post.Reposts.Count,
		LikedIDs:      pq.Int32Array{}, // coming soon ..
		CommentedIDs:  pq.Int32Array{}, // coming soon ..
	}
}

func parseVKPosts(vkPosts []vk.VKPost) []pg.Post {
	posts := make([]pg.Post, len(vkPosts))
	for i := range posts {
		posts[i] = parseVKPost(vkPosts[i])
	}
	return posts
}

func parseVKPhoto(post vk.VKPhoto) pg.Photo {
	return pg.Photo{
		ID:            post.ID,
		OwnerID:       sql.NullInt32{Int32: post.OwnerID, Valid: true},
		Date:          time.Unix(post.Date, 0),
		Text:          post.Text,
		LikesCount:    post.Likes.Count,
		CommentsCount: 0, // coming soon ..
		RepostsCount:  post.Reposts.Count,
		LikedIDs:      pq.Int32Array{}, // coming soon ..
		CommentedIDs:  pq.Int32Array{}, // coming soon ..
	}
}

func parseVKPhotos(vkPhotos []vk.VKPhoto) []pg.Photo {
	photos := make([]pg.Photo, len(vkPhotos))
	for i := range photos {
		photos[i] = parseVKPhoto(vkPhotos[i])
	}
	return photos
}

func parseVKUser(vkUser vk.VKUser) (pg.User, pg.Country, pg.City) {
	country := pg.Country{
		ID:    vkUser.Country.ID,
		Title: vkUser.Country.Title,
	}
	city := pg.City{
		ID:    vkUser.City.ID,
		Title: vkUser.City.Title,
	}
	user := pg.User{
		ID:             vkUser.ID,
		FirstName:      vkUser.FirstName,
		LastName:       vkUser.LastName,
		IsClosed:       vkUser.IsClosed,
		Sex:            vkUser.Sex,
		Domain:         vkUser.Domain,
		BDate:          vkUser.BDate,
		CollectDate:    time.Now(),
		Status:         vkUser.Status,
		Verified:       vkUser.Verified == 1,
		CountryID:      sql.NullInt32{Int32: vkUser.Country.ID, Valid: true},
		CityID:         sql.NullInt32{Int32: vkUser.City.ID, Valid: true},
		HomeTown:       vkUser.Hometown,
		Universities:   []int32{},
		Schools:        pq.Int32Array{},
		FriendsCount:   -1,
		FriendsIDs:     pq.Int32Array{},
		FollowersCount: -1,
		FollowersIDs:   pq.Int32Array{},
		PostsCount:     -1,
		PostsIDs:       pq.Int32Array{},
		PhotosCount:    -1,
		PhotosIDs:      pq.Int32Array{},
		GroupsCount:    -1,
		GroupsIDs:      pq.Int32Array{},
	}
	return user, country, city
}

func parseTrackingUser(userInfo vk.VKUserInfo) (pg.User, pg.Country, pg.City) {
	schoolsIDs := make(pq.Int32Array, len(userInfo.MainInfo.Schools))
	for i := range schoolsIDs {
		schoolsIDs[i] = userInfo.MainInfo.Schools[i].ID
	}

	universitiesIDs := make(pq.Int32Array, len(userInfo.MainInfo.Universities))
	for i := range universitiesIDs {
		universitiesIDs[i] = userInfo.MainInfo.Universities[i].ID
	}

	friendsIDs := make(pq.Int32Array, userInfo.Friends.Count)
	for i := range friendsIDs {
		friendsIDs[i] = userInfo.Friends.Items[i].ID
	}

	followersIDs := make(pq.Int32Array, userInfo.Followers.Count)
	for i := range followersIDs {
		followersIDs[i] = userInfo.Followers.Items[i].ID
	}

	postsIDs := make(pq.Int32Array, userInfo.Posts.Count)
	for i := range postsIDs {
		postsIDs[i] = userInfo.Posts.Items[i].ID
	}

	photosIDs := make(pq.Int32Array, userInfo.Photos.Count)
	for i := range photosIDs {
		photosIDs[i] = userInfo.Photos.Items[i].ID
	}

	groupsIDs := make(pq.Int32Array, userInfo.Groups.Count)
	for i := range groupsIDs {
		groupsIDs[i] = userInfo.Groups.Items[i].ID
	}

	country := pg.Country{
		ID:    userInfo.MainInfo.Country.ID,
		Title: userInfo.MainInfo.Country.Title,
	}
	city := pg.City{
		ID:    userInfo.MainInfo.City.ID,
		Title: userInfo.MainInfo.City.Title,
	}
	user := pg.User{
		ID:             userInfo.MainInfo.ID,
		FirstName:      userInfo.MainInfo.FirstName,
		LastName:       userInfo.MainInfo.LastName,
		IsClosed:       userInfo.MainInfo.IsClosed,
		Sex:            userInfo.MainInfo.Sex,
		Domain:         userInfo.MainInfo.Domain,
		BDate:          userInfo.MainInfo.BDate,
		CollectDate:    time.Now(),
		Status:         userInfo.MainInfo.Status,
		Verified:       userInfo.MainInfo.Verified == 1,
		CountryID:      sql.NullInt32{Int32: userInfo.MainInfo.Country.ID, Valid: true},
		CityID:         sql.NullInt32{Int32: userInfo.MainInfo.City.ID, Valid: true},
		HomeTown:       userInfo.MainInfo.Hometown,
		Universities:   universitiesIDs,
		Schools:        schoolsIDs,
		FriendsCount:   userInfo.Friends.Count,
		FriendsIDs:     friendsIDs,
		FollowersCount: userInfo.Followers.Count,
		FollowersIDs:   followersIDs,
		PostsCount:     userInfo.Posts.Count,
		PostsIDs:       postsIDs,
		PhotosCount:    userInfo.Photos.Count,
		PhotosIDs:      photosIDs,
		GroupsCount:    userInfo.Groups.Count,
		GroupsIDs:      groupsIDs,
	}
	return user, country, city
}
