-- gm_places: GoogleMapsAPI から取得した施設情報を格納するテーブル
CREATE TABLE gm_places
(
    id                 SERIAL PRIMARY KEY,
    place_id           VARCHAR(255) UNIQUE NOT NULL,
    name               VARCHAR(255),
    formatted_address  TEXT,
    icon               VARCHAR(255),
    rating             REAL CHECK (rating >= 1.0 AND rating <= 5.0),
    user_ratings_total INT,
    price_level        INT CHECK (price_level >= 0 AND price_level <= 4),
    vicinity           TEXT,
    permanently_closed BOOLEAN,
    business_status    VARCHAR(255),
    location_latitude  FLOAT,
    location_longitude FLOAT,
    types              TEXT, -- カンマ区切りでタイプ情報を格納
    created_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- photos テーブル: 写真情報
CREATE TABLE gm_place_photos
(
    id              SERIAL PRIMARY KEY,
    place_id        VARCHAR(255) REFERENCES gm_places (place_id) ON DELETE CASCADE,
    photo_reference VARCHAR(255) NOT NULL,
    height          INT,
    width           INT
);
