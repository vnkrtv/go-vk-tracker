package postgres

const dbSchema = `
CREATE TABLE IF NOT EXISTS countries
(
    country_id int2
        not null,
    title      text
        not null,

    CONSTRAINT pk_countries PRIMARY KEY (country_id)
);

CREATE TABLE IF NOT EXISTS cities
(
    city_id int2
        not null,
    title   text
        not null,

    CONSTRAINT pk_cities PRIMARY KEY (city_id)
);

CREATE TABLE IF NOT EXISTS universities
(
    university_id int2
        not null,
    name          text
        not null,
    country_id    int2,
    city_id        int2,

    CONSTRAINT pk_universities PRIMARY KEY (university_id),

    CONSTRAINT fk_country FOREIGN KEY (country_id)
        REFERENCES countries (country_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,

    CONSTRAINT fk_city FOREIGN KEY (city_id)
        REFERENCES cities (city_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS schools
(
    school_id      int4
        not null,
    name           text
        not null,
    year_from      int2,
    year_to        int2,
    year_graduated int2,
    type_str       text,
    country_id     int2,
    city_id        int2,

    CONSTRAINT pk_schools PRIMARY KEY (school_id),

    CONSTRAINT fk_country FOREIGN KEY (country_id)
        REFERENCES countries (country_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,

    CONSTRAINT fk_city FOREIGN KEY (city_id)
        REFERENCES cities (city_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS groups
(
    group_id      int4
        not null,
    name          text
        not null,
    screen_name   text
        not null,
    type          text
        not null,
    members_count int4
        not null,
    is_closed     bool
        not null,

    CONSTRAINT pk_groups PRIMARY KEY (group_id)
);

CREATE TABLE IF NOT EXISTS users
(
    user_id         int4
        not null,
    first_name      text
        not null,
    last_name       text
        not null,
    is_closed       bool
        not null,
    sex             int2
        not null,
    domain          text
        not null,
    bdate           text
        not null,
    collect_date    date
        not null,
    status          text
        not null,
    verified        bool
        not null,
    country_id      int2,
    city_id         int2,
    home_town       text
        not null,
    universities    int2[]
        not null,
    schools         int4[]
        not null,
    friends_count   int4,
    friends_ids     int4[],
    followers_count int4,
    followers_ids   int4[],
    posts_count     int4,
    posts_ids       int4[],
    photos_count    int4,
    photos_ids      int4[],
    groups_count    int4,
    groups_ids      int4[],

    CONSTRAINT pk_users PRIMARY KEY (user_id, collect_date),

    CONSTRAINT fk_country FOREIGN KEY (country_id)
        REFERENCES countries (country_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,

    CONSTRAINT fk_city FOREIGN KEY (city_id)
        REFERENCES cities (city_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS photos
(
    photo_id       int4
        not null,
    owner_id       int4,
    likes_count    int4
        not null,
    comments_count int4
        not null,
    reposts_count  int4
        not null,
    liked_ids      int4[]
        not null,
    commented_ids  int4[]
        not null,
    text           text
        not null,
    date           timestamp
        not null,

    CONSTRAINT pk_photos PRIMARY KEY (photo_id),

    CONSTRAINT fk_owner FOREIGN KEY (owner_id)
        REFERENCES users (user_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS posts
(
    post_id        int4
        not null,
    owner_id       int4,
    likes_count    int4
        not null,
    comments_count int4
        not null,
    reposts_count  int4
        not null,
    views_count    int4
        not null,
    liked_ids      int4[]
        not null,
    commented_ids  int4[]
        not null,
    text           text
        not null,
    date           timestamp
        not null,

    CONSTRAINT pk_posts PRIMARY KEY (post_id),

    CONSTRAINT fk_owner FOREIGN KEY (owner_id)
        REFERENCES users (user_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS tracking_groups
(
    screen_name text,

    CONSTRAINT pk_tracking_groups PRIMARY KEY (screen_name)
);

CREATE TABLE IF NOT EXISTS tracking_users
(
    user_id int4,

    CONSTRAINT pk_tracking_users PRIMARY KEY (user_id)
);
`
