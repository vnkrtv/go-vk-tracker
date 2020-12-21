package vkapi

import (
	"time"

	"github.com/go-vk-api/vk"
)

type VKTracker interface {
	GetUser(userID int) (VKUser, error)
	GetFriends(userID int) (VKUsers, error)
	GetFollowers(userID int) (VKUsers, error)
	GetGroups(userID int) (VKGroups, error)
	GetPhotos(userID int) (VKPhotos, error)
	GetUserPosts(userID int) (VKPosts, error)
	GetGroupPosts(userID int) (VKPosts, error)
}

type VKAPi struct {
	api        *vk.Client
	apiVersion string
	Timeout    float32
}

func NewVKApi(token, apiVersion string, timeout float32) (*VKAPi, error) {
	api, err := vk.NewClientWithOptions(
		vk.WithToken(token),
	)
	return &VKAPi{
		api:        api,
		apiVersion: apiVersion,
		Timeout:    timeout,
	}, err
}

func (a *VKAPi) GetUser(userID int32) (VKUser, error) {
	var user VKUser
	fields := `home_town,schools,status,domain,sex,bdate,country,city,contacts,universities`
	err := a.api.CallMethod("users.get", vk.RequestParams{
		"user_ids": userID,
		"fields": fields,
		"v": a.apiVersion,
	}, &user)
	return user, err
}

func (a *VKAPi) GetFriends(userID int32) (VKUsers, error) {
	var response VKUsers
	fields := "home_town,schools,status,domain,sex,bdate,country,city,contacts,universities"
	err := a.api.CallMethod("friends.get", vk.RequestParams{
		"user_ids": userID,
		"fields": fields,
		"v": a.apiVersion,
	}, &response)
	return response, err
}

func (a *VKAPi) GetFollowers(userID int32) (VKUsers, error) {
	var response VKUsers
	fields := "home_town,schools,status,domain,sex,bdate,country,city,contacts,universities"
	err := a.api.CallMethod("users.getFollowers", vk.RequestParams{
		"user_id": userID,
		"fields": fields,
		"v": a.apiVersion,
	}, &response)
	return response, err
}

func (a *VKAPi) GetGroups(userID int32) (VKGroups, error) {
	var response VKGroups
	fields := "id,name,screen_name,members_count"
	err := a.api.CallMethod("groups.get", vk.RequestParams{
		"user_id": userID,
		"extended": 1,
		"fields": fields,
		"v": a.apiVersion,
	}, &response)
	return response, err
}

func (a *VKAPi) GetPhotos(userID int32) (VKPhotos, error) {
	var response VKPhotos
	err := a.api.CallMethod("photos.getAll", vk.RequestParams{
		"owner_id": userID,
		"count": 200,
		"extended": 1,
		"v": a.apiVersion,
	}, &response)
	return response, err
}

func (a *VKAPi) GetUserPosts(userID int32) (VKPosts, error) {
	var response VKPosts
	err := a.api.CallMethod("wall.get", vk.RequestParams{
		"owner_id": userID,
		"count": 100,
		"v": a.apiVersion,
	}, &response)
	return response, err
}

func (a *VKAPi) GetGroupPosts(screenName string) (VKGroupInfo, error) {
	var response VKGroupInfo
	fields := "id,name,screen_name,members_count"
	err := a.api.CallMethod("wall.get", vk.RequestParams{
		"domain": screenName,
		"count": 100,
		"fields": fields,
		"extended": 1,
		"v": a.apiVersion,
	}, &response)
	return response, err
}

func (a *VKAPi) GetUserInfo(userID int32) (*VKUserInfo, error) {
	user, err := a.GetUser(userID)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(a.Timeout) * time.Second)
	friends, err := a.GetFriends(userID)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(a.Timeout) * time.Second)
	followers, err := a.GetFollowers(userID)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(a.Timeout) * time.Second)
	groups, err := a.GetGroups(userID)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(a.Timeout) * time.Second)
	posts, err := a.GetUserPosts(userID)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Duration(a.Timeout) * time.Second)
	photos, err := a.GetPhotos(userID)
	if err != nil {
		return nil, err
	}

	return &VKUserInfo{
		MainInfo:  user,
		Friends:   friends,
		Followers: followers,
		Groups:    groups,
		Posts:     posts,
		Photos:    photos,
	}, err
}

func (a *VKAPi) GetGroupInfo(screenName string) (VKGroup, []VKPost, error) {
	response, err := a.GetGroupPosts(screenName)
	if err != nil {
		return VKGroup{}, nil, err
	}
	return response.Groups[0], response.Items, err
}