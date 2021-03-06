package service

import (
	"log"

	pg "github.com/vnkrtv/go-vk-tracker/pkg/postgres"
	vk "github.com/vnkrtv/go-vk-tracker/pkg/vkapi"
)

type VKLoader struct {
	db    *pg.Storage
	vkApi *vk.VKAPi
}

func NewVKLoader(vkToken, apiVersion string, timeout int32, pgUser, pgPass, pgHost, pgPort, pgDBName string) (*VKLoader, error) {
	db, err := pg.OpenConnection(pgUser, pgPass, pgHost, pgPort, pgDBName)
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to PostgreSQL(%s:%s@%s:%s/%s)", pgUser, pgPass, pgHost, pgPort, pgDBName)
	api, err := vk.NewVKApi(vkToken, apiVersion, timeout)
	if err != nil {
		return nil, err
	}
	log.Printf("Open VKApi client connection (api_version=%s,timeout=%d milliseconds)", apiVersion, timeout)
	return &VKLoader{
		db:    db,
		vkApi: api,
	}, err
}

func NewVKLoaderFromCfg(cfg Config) (*VKLoader, error) {
	return NewVKLoader(cfg.VKToken, cfg.VKApiVersion, cfg.Timeout,
		cfg.PGUser, cfg.PGPass, cfg.PGHost, cfg.PGPort, cfg.PGName)
}

func (s *VKLoader) InitDB() error {
	return s.db.CreateSchema()
}

func (s *VKLoader) Sleep(millisecondNum int32) {
	s.vkApi.Sleep(millisecondNum)
}

func (s *VKLoader) AddTrackingUsers(usersIDs []int32) error {
	for _, userID := range usersIDs {
		if err := s.db.AddTrackingUser(userID); err != nil {
			return err
		}
		log.Printf("Start tracking user[id=%d]", userID)
	}
	return nil
}

func (s *VKLoader) AddTrackingGroups(screenNames []string) error {
	for _, screenName := range screenNames {
		if err := s.db.AddTrackingGroup(screenName); err != nil {
			return err
		}
		log.Printf("Start tracking group[screen_name=%s]", screenName)
	}
	return nil
}

func (s *VKLoader) LoadGroupsInfo() error {
	groupScreenNames, err := s.db.GetTrackingGroups()
	if err != nil {
		return err
	}
	for _, screenName := range groupScreenNames {
		vkGroup, vkPosts, err := s.vkApi.GetGroupInfo(screenName)
		s.Sleep(s.vkApi.Timeout)
		if err != nil {
			log.Printf("Error on getting group[screen_name=%s] info: %s", screenName, err)
			log.Println()
		} else {
			group := parseVKGroup(vkGroup)
			if err := s.db.InsertGroup(group); err != nil {
				log.Printf("Error on inserting group in db: %s", err)
			}

			posts := parseVKPosts(vkPosts)
			if err := s.db.InsertPosts(posts); err != nil {
				log.Printf("Error on inserting post in db: %s", err)
			}
			log.Printf("Get full group[screen_name=%s] info", screenName)
			log.Println()
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
		s.Sleep(s.vkApi.Timeout)
		if err != nil {
			log.Printf("Error on getting user[id=%d] info: %s", userID, err)
			log.Println()
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

			if err = s.db.InsertUser(user, true); err != nil {
				log.Printf("Error on inserting user in db: %s", err)
			}

			posts := parseVKPosts(userInfo.Posts.Items)
			if err = s.db.InsertPosts(posts); err != nil {
				log.Printf("Error on inserting post in db: %s", err)
			}

			photos := parseVKPhotos(userInfo.Photos.Items)
			if err = s.db.InsertPhotos(photos); err != nil {
				log.Printf("Error on inserting photo in db: %s", err)
			}

			for _, vkFriend := range userInfo.Friends.Items {
				friend, country, city := parseVKUser(vkFriend)
				if err = s.db.InsertCountry(country); err != nil {
					log.Printf("Error on inserting country in db: %s", err)
				}
				if err = s.db.InsertCity(city); err != nil {
					log.Printf("Error on inserting city in db: %s", err)
				}

				universities, countries, cities := parseVKUniversities(vkFriend.Universities)
				if err = s.db.InsertCountries(countries); err != nil {
					log.Printf("Error on inserting country in db: %s", err)
				}
				if err = s.db.InsertCities(cities); err != nil {
					log.Printf("Error on inserting city in db: %s", err)
				}
				if err = s.db.InsertUniversities(universities); err != nil {
					log.Printf("Error on inserting university in db: %s", err)
				}

				schools, countries, cities := parseVKSchools(vkFriend.Schools)
				if err = s.db.InsertCountries(countries); err != nil {
					log.Printf("Error on inserting country in db: %s", err)
				}
				if err = s.db.InsertCities(cities); err != nil {
					log.Printf("Error on inserting city in db: %s", err)
				}
				if err = s.db.InsertSchools(schools); err != nil {
					log.Printf("Error on inserting school in db: %s", err)
				}

				if err = s.db.InsertUser(friend, false); err != nil {
					log.Printf("Error on inserting friend in db: %s", err)
				}
			}
			log.Printf("Upsert user[id=%d] %d friends in db", user.ID, len(userInfo.Friends.Items))

			for _, vkFollower := range userInfo.Followers.Items {
				follower, country, city := parseVKUser(vkFollower)
				if err = s.db.InsertCountry(country); err != nil {
					log.Printf("Error on inserting country in db: %s", err)
				}
				if err = s.db.InsertCity(city); err != nil {
					log.Printf("Error on inserting city in db: %s", err)
				}

				universities, countries, cities := parseVKUniversities(vkFollower.Universities)
				if err = s.db.InsertCountries(countries); err != nil {
					log.Printf("Error on inserting country in db: %s", err)
				}
				if err = s.db.InsertCities(cities); err != nil {
					log.Printf("Error on inserting city in db: %s", err)
				}
				if err = s.db.InsertUniversities(universities); err != nil {
					log.Printf("Error on inserting university in db: %s", err)
				}

				schools, countries, cities := parseVKSchools(vkFollower.Schools)
				if err = s.db.InsertCountries(countries); err != nil {
					log.Printf("Error on inserting country in db: %s", err)
				}
				if err = s.db.InsertCities(cities); err != nil {
					log.Printf("Error on inserting city in db: %s", err)
				}
				if err = s.db.InsertSchools(schools); err != nil {
					log.Printf("Error on inserting school in db: %s", err)
				}

				if err = s.db.InsertUser(follower, false); err != nil {
					log.Printf("Error on inserting follower in db: %s", err)
				}
			}
			log.Printf("Upsert user[id=%d] %d followers in db", user.ID, len(userInfo.Followers.Items))

			log.Printf("Get user[id=%d,domain=%s,first_name=%s,second_name=%s] info",
				userInfo.MainInfo.ID, userInfo.MainInfo.FirstName, userInfo.MainInfo.LastName, userInfo.MainInfo.Domain)
			log.Println()
		}

	}
	return err
}
