CREATE TABLE users
(
    id       BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    login    varchar(128) NOT NULL,
    email    varchar(128) NOT NULL,
    password varchar(255) NOT NULL
);
CREATE TABLE categories
(
    id       BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name    varchar(128) NOT NULL
);
CREATE TABLE spendings
(
    id       BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    price    INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    user_id BIGINT NOT NULL,
    categories_id BIGINT,

    CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT categorie_id_fk FOREIGN KEY (categories_id) REFERENCES categories (id)
);