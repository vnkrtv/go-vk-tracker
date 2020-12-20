package vkapi

type VKCountry struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

type VKCity struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

type VKUniversity struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	CityID    int32  `json:"city"`
	CountryID int32  `json:"country"`
}

type VKSchool struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	YearFrom      int32  `json:"year_from"`
	YearTo        int32  `json:"year_to"`
	YearGraduated int32  `json:"year_graduated"`
	Type          string `json:"type_str"`
	CityID        int32  `json:"city"`
	CountryID     int32  `json:"country_id"`
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
	Hometown     string         `json:"home_town"`
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
	OwnerID  int32  `json:"owner_id"`
	Date     int64  `json:"date"`
	PostType string `json:"post_type"`
	Text     string `json:"text"`
	Likes    struct {
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
	OwnerID  int32  `json:"owner_id"`
	Date     int64  `json:"date"`
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
