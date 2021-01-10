package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type VKStorage interface {
	CreateSchema() error
}

type Storage struct {
	db *sqlx.DB
}

func OpenConnection(user, password, host, port, dbName string) (*Storage, error) {
	conStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName)
	db, err := sqlx.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return &Storage{db: db}, err
}

func (s *Storage) CreateSchema() error {
	_, err := s.db.Exec(dbSchema)
	return err
}

func (s *Storage) GetTrackingUsers() ([]int32, error) {
	var userIDs []int32
	err := s.db.Select(&userIDs, "SELECT user_id FROM tracking_users")
	return userIDs, err
}

func (s *Storage) GetTrackingGroups() ([]string, error) {
	var screenNames []string
	err := s.db.Select(&screenNames, "SELECT screen_name FROM tracking_groups")
	return screenNames, err
}

func (s *Storage) AddTrackingUser(userID int32) error {
	sql := `
		INSERT INTO 
			tracking_users (user_id) 
		VALUES 
			($1)
		ON CONFLICT (user_id)
    		DO NOTHING`
	_, err := s.db.Exec(sql, userID)
	return err
}

func (s *Storage) AddTrackingGroup(screenName string) error {
	sql := `
		INSERT INTO 
			tracking_groups (screen_name) 
		VALUES 
			($1)
		ON CONFLICT (screen_name)
    		DO NOTHING`
	_, err := s.db.Exec(sql, screenName)
	return err
}

func (s *Storage) InsertCountry(country Country) error {
	sql := `
		INSERT INTO 
			countries (country_id, title) 
		VALUES 
			(:country_id, :title)
		ON CONFLICT (country_id)
    		DO NOTHING`
	_, err := s.db.NamedExec(sql, &country)
	return err
}

func (s *Storage) InsertCountries(countries []Country) error {
	for _, country := range countries {
		if err := s.InsertCountry(country); err != nil {
			return err
		}
	}
	log.Printf("Upsert %d countries in db", len(countries))
	return nil
}

func (s *Storage) InsertCity(city City) error {
	sql := `
		INSERT INTO 
			cities (city_id, title) 
		VALUES 
			(:city_id, :title)
		ON CONFLICT (city_id)
    		DO NOTHING`
	_, err := s.db.NamedExec(sql, &city)
	return err
}

func (s *Storage) InsertCities(cities []City) error {
	for _, city := range cities {
		if err := s.InsertCity(city); err != nil {
			return err
		}
	}
	log.Printf("Upsert %d cities in db", len(cities))
	return nil
}

func (s *Storage) InsertUniversity(university University) error {
	sql := `
		INSERT INTO 
			universities (university_id, name, country_id, city_id) 
		VALUES 
			(:university_id, :name, :country_id, :city_id)
		ON CONFLICT (university_id)
    		DO NOTHING`
	_, err := s.db.NamedExec(sql, &university)
	return err
}

func (s *Storage) InsertUniversities(universities []University) error {
	for _, university := range universities {
		if err := s.InsertUniversity(university); err != nil {
			return err
		}
	}
	log.Printf("Upsert %d universities in db", len(universities))
	return nil
}

func (s *Storage) InsertSchool(school School) error {
	sql := `
		INSERT INTO 
			schools (school_id, name, year_from, year_to, 
			         year_graduated, type_str, country_id, city_id) 
		VALUES 
			(:school_id, :name, :year_from, :year_to, 
			 :year_graduated, :type_str, :country_id, :city_id)
		ON CONFLICT (school_id)
    		DO NOTHING`
	_, err := s.db.NamedExec(sql, &school)
	return err
}

func (s *Storage) InsertSchools(schools []School) error {
	for _, school := range schools {
		if err := s.InsertSchool(school); err != nil {
			return err
		}
	}
	log.Printf("Upsert %d schools in db", len(schools))
	return nil
}

func (s *Storage) InsertPhoto(photo Photo) error {
	sql := `
		INSERT INTO 
			photos (photo_id, owner_id, date, text, likes_count, 
			       comments_count, reposts_count, liked_ids, commented_ids) 
		VALUES 
			(:photo_id, :owner_id, :date, :text, :likes_count, 
			 :comments_count, :reposts_count, :liked_ids, :commented_ids)
		ON CONFLICT (photo_id)
		DO UPDATE SET
			text = :text, likes_count = :likes_count, comments_count = :comments_count, 
		    reposts_count = :reposts_count, liked_ids = :liked_ids, commented_ids = :commented_ids`
	_, err := s.db.NamedExec(sql, &photo)
	return err
}

func (s *Storage) InsertPhotos(photos []Photo) error {
	for _, photo := range photos {
		if err := s.InsertPhoto(photo); err != nil {
			return err
		}
	}
	if len(photos) > 0 {
		log.Printf("Upsert %d photos[owner_id=%d] in db", len(photos), photos[0].OwnerID.Int32)
	}
	return nil
}

func (s *Storage) InsertPost(post Post) error {
	sql := `
		INSERT INTO 
			posts (post_id, owner_id, date, text, likes_count, views_count, 
			       comments_count, reposts_count, liked_ids, commented_ids) 
		VALUES 
			(:post_id, :owner_id, :date, :text, :likes_count, :views_count, 
			 :comments_count, :reposts_count, :liked_ids, :commented_ids)
		ON CONFLICT (post_id)
		DO UPDATE SET
			text = :text, likes_count = :likes_count, views_count = :views_count, comments_count = :comments_count, 
		    reposts_count = :reposts_count, liked_ids = :liked_ids, commented_ids = :commented_ids`
	_, err := s.db.NamedExec(sql, &post)
	return err
}

func (s *Storage) InsertPosts(posts []Post) error {
	for _, post := range posts {
		if err := s.InsertPost(post); err != nil {
			return err
		}
	}
	if len(posts) > 0 {
		log.Printf("Upsert %d posts[owner_id=%d] in db", len(posts), posts[0].OwnerID.Int32)
	}
	return nil
}

func (s *Storage) InsertGroup(group Group) error {
	sql := `
		INSERT INTO 
			groups (group_id, screen_name, name, members_count, type, is_closed) 
		VALUES 
			(:group_id, :screen_name, :name, :members_count, :type, :is_closed)
		ON CONFLICT (group_id)
    		DO UPDATE SET
    			name = :name, members_count = :members_count, is_closed = :is_closed`
	_, err := s.db.NamedExec(sql, &group)
	return err
}

func (s *Storage) InsertGroups(groups []Group) error {
	for _, group := range groups {
		if err := s.InsertGroup(group); err != nil {
			return err
		}
	}
	log.Printf("Upsert %d groups in db", len(groups))
	return nil
}

func (s *Storage) InsertUser(user User, loadedUser bool) error {
	var sql, params string
	sql = `
		INSERT INTO 
			users (user_id, first_name, last_name, is_closed, sex, domain, bdate, city_id,
			       collect_date, status, verified, country_id, home_town, universities,
			       schools, friends_count, friends_ids, followers_count, followers_ids,
			       posts_count, posts_ids, photos_count, photos_ids, groups_count, groups_ids) 
		VALUES 
			(:user_id, :first_name, :last_name, :is_closed, :sex, :domain, :bdate, :city_id,
			 :collect_date, :status, :verified, :country_id, :home_town, :universities,
			 :schools, :friends_count, :friends_ids, :followers_count, :followers_ids,
			 :posts_count, :posts_ids, :photos_count, :photos_ids, :groups_count, :groups_ids)
		ON CONFLICT (user_id, collect_date)
    		DO UPDATE SET
    			%s`
	if loadedUser {
		params = `
				first_name = :first_name, last_name = :last_name, is_closed = :is_closed, city_id = :city_id,
    			domain = :domain, bdate = :bdate, status = :status, verified = :verified,
    		    country_id = :country_id, home_town = :home_town, universities = :universities,
    		    schools = :schools, friends_count = :friends_count, friends_ids = :friends_ids,
    		    followers_count = :followers_count, followers_ids = :followers_ids, posts_count = :posts_count,
    		    posts_ids = :posts_ids, photos_count = :photos_count, photos_ids = :photos_ids, 
    		    groups_count = :groups_count, groups_ids = :groups_ids`
	} else {
		params = `
				first_name = :first_name, last_name = :last_name, status = :status`
	}
	_, err := s.db.NamedExec(fmt.Sprintf(sql, params), &user)
	return err
}

func (s *Storage) InsertUsers(users []User) error {
	for _, user := range users {
		if err := s.InsertUser(user, false); err != nil {
			return err
		}
	}
	log.Printf("Upsert %d users in db", len(users))
	return nil
}
