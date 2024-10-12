-- migrate:up
CREATE TABLE pcr_phone (
    id           SERIAL PRIMARY KEY,
    model        VARCHAR,
    mac          VARCHAR,
    username     VARCHAR,
    password     VARCHAR,
    display_name VARCHAR,
    UNIQUE (model, mac)
)