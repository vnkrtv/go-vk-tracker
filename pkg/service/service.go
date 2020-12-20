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
			user, country, city := parseTrackingUser(*userInfo)
			if err = s.db.InsertCountry(country); err != nil {
				log.Printf("Error on inserting country in db: %s", err)
			}
			if err = s.db.InsertCity(city); err != nil {
				log.Printf("Error on inserting city in db: %s", err)
			}

			universities, countries, cities := parseVKUniversities(userInfo.MainInfo.Universities)
			if err = s.db.InsertCountries(countries); err != nil {
				log.Printf("Error on inserting country in db: %s", err)
			}
			if err = s.db.InsertCities(cities); err != nil {
				log.Printf("Error on inserting city in db: %s", err)
			}
			if err = s.db.InsertUniversities(universities); err != nil {
				log.Printf("Error on inserting university in db: %s", err)
			}

			schools, countries, cities := parseVKSchools(userInfo.MainInfo.Schools)
			if err = s.db.InsertCountries(countries); err != nil {
				log.Printf("Error on inserting country in db: %s", err)
			}
			if err = s.db.InsertCities(cities); err != nil {
				log.Printf("Error on inserting city in db: %s", err)
			}
			if err = s.db.InsertSchools(schools); err != nil {
				log.Printf("Error on inserting school in db: %s", err)
			}

			groups := parseVKGroups(userInfo.Groups.Items)
			if err = s.db.InsertGroups(groups); err != nil {
				log.Printf("Error on inserting group in db: %s", err)
			}

			if err = s.db.InsertUser(user); err != nil {
				log.Printf("Error on inserting user in db: %s", err)
			}

			for _, friend := range parseVK

		}

	}
	return err
}