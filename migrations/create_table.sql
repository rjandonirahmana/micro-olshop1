CREATE TABLE products (
    id          VARCHAR(50)                 NOT NULL,
    seller_id   VARCHAR(50)                 NOT NULL,
    name        VARCHAR(50)                 DEFAULT "products1",
    price       INT(10)                     NOT NULL,
    quantity    INT(10)                     DEFAULT 1,
    category_id INT(10)                     NOT NULL,
    description VARCHAR(1024),
    created_at   DATETIME    DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME    DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE miss_indonesia (
  id int(11) NOT NULL AUTO_INCREMENT,
  user_id varchar(25) NOT NULL,
  province_representation_id int(11) NOT NULL,
  reference_from varchar(255) NOT NULL,
  nick_name varchar(100) NOT NULL,
  place_of_birth varchar(100) NOT NULL,
  height int(11) NOT NULL,
  weight int(11) NOT NULL 
);


DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS seller CASCADE;


CREATE TABLE IF NOT EXISTS seller
(
    id           VARCHAR(50)                     NOT NULL,
    name         VARCHAR(50)                 NOT NULL CHECK ( name <> '' ),
    email        VARCHAR(64) UNIQUE          NOT NULL CHECK ( email <> '' ),
    password     VARCHAR(250)                NOT NULL CHECK ( octet_length(password) <> 0 ),
    salt         VARCHAR(100)                NOT NULL CHECK ( salt <> ''),
    avatar       VARCHAR(512),
    phone_number VARCHAR(20),
    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,
    login_date   TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS customer
(
    id           VARCHAR(50)                     NOT NULL,
    name         VARCHAR(50)                 NOT NULL CHECK ( name <> '' ),
    email        VARCHAR(64) UNIQUE          NOT NULL CHECK ( email <> '' ),
    password     VARCHAR(250)                NOT NULL CHECK ( octet_length(password) <> 0 ),
    salt         VARCHAR(100)                NOT NULL CHECK ( salt <> ''),
    avatar       VARCHAR(512),
    phone_number VARCHAR(20),
    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP WITH TIME ZONE             DEFAULT CURRENT_TIMESTAMP,
    login_date   TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_category
(
    id           INT(10)                     NOT NULL,
    name         VARCHAR(30)                 NOT NULL CHECK ( name <> '' ),
    PRIMARY KEY (id)
);


CREATE TABLE product_images
(
    product_id           INT(10)                    NOT NULL,
    name                VARCHAR(30)                 NOT NULL CHECK ( name <> '' ),
    is_primary           INT(10)                    NOT NULL
);


ALTER TABLE products
  ADD FOREIGN KEY (category_id) REFERENCES product_category(id);

ALTER TABLE product_images
  ADD FOREIGN KEY (product_id) REFERENCES products(id);


ALTER TABLE product_images
  MODIFY product_id varchar(50);