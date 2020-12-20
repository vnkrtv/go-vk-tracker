package vkapi

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
	Country      VKCountry      `json:"country"`
	City         VKCity         `json:"city"`
	Universities []VKUniversity `json:"universities"`
	Schools      []VKSchool     `json:"schools"`
}

type VKUsers struct {
	Count int      `json:"count"`
	Items []VKUser `json:"items"`
}

type VKGroup struct {
	ID           int    `json:"id"`
	ScreenName   string `json:"screen_name"`
	Name         string `json:"name"`
	MembersCount int    `json:"members_count"`
	Type         string `json:"type"`
	IsClosed     bool   `json:"is_closed"`
}

type VKGroups struct {
	Count int       `json:"count"`
	Items []VKGroup `json:"items"`
}

type VKPhoto struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Date     int    `json:"date"`
	PostType string `json:"post_type"`
	Text     string `json:"text"`
	Likes struct {
		Count int `json:"count"`
	} `json:"likes"`
	Reposts struct {
		Count int `json:"count"`
	} `json:"reposts"`
}

type VKPhotos struct {
	Count int       `json:"count"`
	Items []VKPhoto `json:"items"`
}

type VKPost struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Date     int    `json:"date"`
	PostType string `json:"post_type"`
	Text     string `json:"text"`
	Comments struct {
		Count int `json:"count"`
	} `json:"comments"`
	Likes struct {
		Count int `json:"count"`
	} `json:"likes"`
	Reposts struct {
		Count int `json:"count"`
	} `json:"reposts"`
	Views struct {
		Count int `json:"count"`
	} `json:"views"`
}

type VKPosts struct {
	Count int      `json:"count"`
	Items []VKPost `json:"items"`
}

type VKGroupInfo struct {
	Count  int       `json:"count"`
	Items  []VKPost  `json:"items"`
	Groups []VKGroup `json:"groups"`
}

type VKUserInfo struct {
	MainInfo  VKUser
	Friends   VKUsers
	Followers VKUsers
	Groups    VKGroups
	Posts     VKPosts
	Photos    VKPhotos
}
