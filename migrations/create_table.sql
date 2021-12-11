CREATE TABLE IF NOT EXIST products (
    id          UUID PRIMARY KEY            DEFAULT uuid_generate_v4(),
    seller_id   UUID DEFAULT                NOT NULL,
    name        VARCHAR(50)                 NOT NULL DEFAULT `products1`,
    price       INTEGER                     NOT NULL,
    quantity    INTEGER                     NOT NULL DEFAULT 1,
    category_id INTEGER                     NOT NULL,
    description VARCHAR(1024),
    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE    DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS seller CASCADE;


CREATE TABLE IF NOT EXISTS seller
(
    id           VARCHAR(50)                     NOT NULL,
    name         VARCHAR(50)                 NOT NULL CHECK ( first_name <> '' ),
    email        VARCHAR(64) UNIQUE          NOT NULL CHECK ( email <> '' ),
    password     VARCHAR(250)                NOT NULL CHECK ( octet_length(password) <> 0 ),
    salt         VARCHAR(100)                NOT NULL CHECK ( salt <> ''),
    avatar       VARCHAR(512),
    phone_number VARCHAR(20),
    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,
    login_date   TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- type Seller struct {
-- 	ID        string          `bson:"_id" db:"id" json:"id"`
-- 	Name      string          `bson:"name" db:"name" json:"name"`
-- 	Email     string          `bson:"email" db:"email" json:"email"`
-- 	Phone     string          `bson:"phone" db:"phone" json:"phone_number"`
-- 	Password  string          `bson:"password" db:"password" json:"-"`
-- 	Salt      string          `bson:"salt" db:"salt" json:"-"`
-- 	Avatar    string          `bson:"avatar" db:"avatar" json:"avatar"`
-- 	Address   []AddressSeller `bson:"address" db:"address_seller" json:"address"`
-- 	CreatedAt time.Time       `bson:"created_at" db:"created_at" json:"created_at"`
-- 	UpdatedAt time.Time       `bson:"updated_at" db:"updated_at" json:"updated_at"`
-- }

-- CREATE TABLE news
-- (
--     news_id    UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
--     author_id  UUID                     NOT NULL REFERENCES users (user_id),
--     title      VARCHAR(250)             NOT NULL CHECK ( title <> '' ),
--     content    TEXT                     NOT NULL CHECK ( content <> '' ),
--     image_url  VARCHAR(1024) check ( image_url <> '' ),
--     category   VARCHAR(250),
--     created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
--     updated_at TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
-- );