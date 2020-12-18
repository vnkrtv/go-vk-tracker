package vkapi

type VKFriends struct {
	Count int        `json:"count"`
	Items []VKFriend `json:"items"`
}

type VKFriend struct {
	ID           int            `json:"id"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	IsClosed     bool           `json:"is_closed"`
	Sex          int            `json:"sex"`
	Domain       string         `json:"domain"`
	BDate        string         `json:"bdate"`
	Status       string         `json:"status"`
	Verified     int            `json:"verified"`
	Country      VKCountry      `json:"country"`
	City         VKCity         `json:"city"`
	Universities []VKUniversity `json:"universities"`
	Schools      []VKSchool     `json:"schools"`
}

type VKCountry struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type VKCity struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type VKUniversity struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CityID    int    `json:"city"`
	CountryID int    `json:"country"`
}

type VKSchool struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	YearFrom      int    `json:"year_from"`
	YearTo        int    `json:"year_to"`
	YearGraduated int    `json:"year_graduated"`
	Type          int    `json:"type_str"`
	CityID        int    `json:"city"`
	CountryID     int    `json:"country_id"`
}

type VKUser struct {
	ID           int            `json:"id"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	IsClosed     bool           `json:"is_closed"`
	Sex          int            `json:"sex"`
	Domain       string         `json:"domain"`
	BDate        string         `json:"bdate"`
	Status       string         `json:"status"`
	Verified     int            `json:"verified"`
	UniversityID int            `json:"university"`
	HomeTown     string         `json:"home_town"`
	Country      VKCountry      `json:"country"`
	Universities []VKUniversity `json:"universities"`
	Schools      []VKSchool     `json:"schools"`
}

type VKFriend struct {

}

type VKPost struct {
	ID          int             `json:"id"`
	Date        int             `json:"date"`
	PostType    string          `json:"post_type"`
	Text        string          `json:"text"`
	IsPinned    int8            `json:"is_pinned"`
	Comments    struct {
		Count		int	     		 `json:"count"`
	}                           `json:"comments"`
	Likes       struct {
		Count		int  			 `json:"count"`
	}                           `json:"likes"`
	Reposts     struct {
		Count		int  			 `json:"count"`
	}                           `json:"reposts"`
	Views       struct {
		Count		int 			 `json:"count"`
	}                           `json:"views"`
}
