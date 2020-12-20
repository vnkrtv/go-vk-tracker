package vkapi

import (
	"fmt"
	"strconv"

	"github.com/go-vk-api/vk"
)

type VKAPi struct {
	api        *vk.Client
	apiVersion string
}

func NewVKApi(token, apiVersion string) (*VKAPi, error) {
	api, err := vk.NewClientWithOptions(
		vk.WithToken(token),
	)
	return &VKAPi{
		api: api,
		apiVersion: apiVersion,
	}, err
}

func (a *VKAPi) GetUser(userID int) (VKUser, error) {
	var user VKUser
	fields := `home_town,schools,status,domain,sex,bdate,country,city,contacts,universities`
	err := a.api.CallMethod("users.get", vk.RequestParams{
		"user_ids": userID,
		"fields": fields,
		"v": a.apiVersion,
	}, &user)
	return user, err
}

func (a *VKAPi) GetFriends(userID int) (VKUsers, error) {
	var response VKUsers
	fields := "home_town,schools,status,domain,sex,bdate,country,city,contacts,universities"
	err := a.api.CallMethod("friends.get", vk.RequestParams{
		"user_ids": userID,
		"fields": fields,
		"v": a.apiVersion,
	}, &response)
	return response, err
}

func (a *VKAPi) GetFollowers(userID int) (VKUsers, error) {
	var response VKUsers
	fields := "home_town,schools,status,domain,sex,bdate,country,city,contacts,universities"
	err := a.api.CallMethod("users.getFollowers", vk.RequestParams{
		"user_id": userID,
		"fields": fields,
		"v": a.apiVersion,
	}, &response)
	return response, err
}

func (a *VKAPi) GetGroups(userID int) (VKGroups, error) {
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

func (a *VKAPi) GetGroupsPosts(groupsScreenNames []string, postsCount int) (map[string]VKWall, error) {
	groups := make([]string, len(groupsScreenNames))
	for i, str := range groupsScreenNames {
		groups[i] = fmt.Sprintf("%s", strconv.Quote(str))
		if i != len(groups) - 1{
			groups[i] += ","
		}
	}
	var response []VKWall
	code := `
        var domains = %s;
		var res = [];
		var i = 0;
		while (i < domains.length) {
			var posts = API.wall.get({
				domain: domains[i], 
				count: %d,
				offset: 1
			});
			res.push(posts);
			i = i + 1; 
		}
		return res;`
	err := a.api.CallMethod("execute", vk.RequestParams{
		"code": fmt.Sprintf(code, groups, postsCount),
	}, &response)
	wallMap := make(map[string]VKWall, len(groups))
	for i, wall := range response {
		wallMap[groupsScreenNames[i]] = wall
	}
	return wallMap, err
}