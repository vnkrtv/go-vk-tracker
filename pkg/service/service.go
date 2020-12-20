package service

import (
	"github.com/pkg/errors"
	"log"

	pg "../postgres"
	vk "../vkapi"
)

var (
	IncorrectScreenName = errors.New("incorrect group screen name")
	IncorrectUserID = errors.New("incorrect user id")
)

type VKLoader struct {
	db          *pg.Storage
	vkApi       *vk.VKAPi
}

func NewVKLoader(vkToken, apiVersion string, timeout float32, pgUser, pgPass, pgHost, pgPort, pgDBName string) (*VKLoader, error) {
	db, err := pg.OpenConnection(pgUser, pgPass, pgHost, pgPort, pgDBName)
	if err != nil {
		return nil, err
	}
	api, err := vk.NewVKApi(vkToken, apiVersion, timeout)
	if err != nil {
		return nil, err
	}
	return &VKLoader{
		db:          db,
		vkApi:       api,
	}, err
}

func (s *VKLoader) InitDB() error {
	return s.db.CreateSchema()
}

func (s *VKLoader) AddTrackingUsers(usersIDs []int) error {
	for _, userID := range usersIDs {
		if err := s.db.AddTrackingUser(userID); err != nil {
			return err
		}
	}
	return nil
}

func (s *VKLoader) AddTrackingGroups(screenNames []string) error {
	for _, screenName := range screenNames {
		if err := s.db.AddTrackingGroup(screenName); err != nil {
			return err
		}
	}
	return nil
}

func (s *VKLoader) LoadUsersInfo() error {
	userIDs, err := s.db.GetTrackingUsers()
	if err != nil {
		return err
	}
	for _, userID := range userIDs {
		userInfo, err := s.vkApi.GetUserInfo(userID)
		if err != nil {
			log.Printf("Error on getting user info: %s", err)
		} else {
			country := parseVKCountry(userInfo.MainInfo.Country)

			city := parseVKCity(userInfo.MainInfo.City)
		}

	}
	return err
}