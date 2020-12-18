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

CREATE TABLE IF NOT EXISTS universities
(
    university_id int2
        not null,
    name          text
        not null,
    country_id    int2,

    CONSTRAINT pk_universities PRIMARY KEY (university_id),

    CONSTRAINT fk_country FOREIGN KEY (country_id)
        REFERENCES countries (country_id)
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

    CONSTRAINT pk_schools PRIMARY KEY (school_id),

    CONSTRAINT fk_country FOREIGN KEY (country_id)
        REFERENCES countries (country_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS groups
(
    group_id    int4
        not null,
    name        text
        not null,
    screen_name text
        not null,
    type        text,
    is_closed   bool,

    CONSTRAINT pk_groups PRIMARY KEY (group_id)
);

CREATE TABLE IF NOT EXISTS photos
(
    photo_id       int4
        not null,
    likes_count    int4
        not null,
    comments_count int4
        not null,
    liked_ids      int4[]
        not null,
    commented_ids  int4[]
        not null,

    CONSTRAINT pk_photos PRIMARY KEY (photo_id)
);

CREATE TABLE IF NOT EXISTS posts
(
    post_id        int4
        not null,
    likes_count    int4
        not null,
    comments_count int4
        not null,
    liked_ids      int4[]
        not null,
    commented_ids  int4[]
        not null,
    text           text,

    CONSTRAINT pk_posts PRIMARY KEY (post_id)
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
    university_id   int2,
    country_id      int2,
    home_town       text
        not null,
    universities    int2[]
        not null,
    schools         int4[]
        not null,
    friends_count   int4
        not null,
    friends_ids     int4[]
        not null,
    followers_count int4
        not null,
    followers_ids   int4[]
        not null,
    posts_count     int4
        not null,
    posts_ids       int4[]
        not null,
    photos_count    int4
        not null,
    photos_ids      int4[]
        not null,
    groups_count    int4
        not null,
    groups_ids      int4[]
        not null,

    CONSTRAINT pk_users PRIMARY KEY (user_id, collect_date),

    CONSTRAINT fk_country FOREIGN KEY (country_id)
        REFERENCES countries (country_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,

    CONSTRAINT fk_university FOREIGN KEY (university_id)
        REFERENCES universities (university_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);
`
