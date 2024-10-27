-- Table: users
CREATE TABLE users
(
    id            CHAR(26) PRIMARY KEY,
    username      VARCHAR(255)        NOT NULL,
    email         VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT                NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT NOW(),
    updated_at    TIMESTAMPTZ DEFAULT NOW()
);

-- gm_places: GoogleMapsAPI から取得した施設情報を格納するテーブル
CREATE TABLE gm_places
(
    id                 VARCHAR(255) PRIMARY KEY,
    place_id           VARCHAR(255) UNIQUE NOT NULL,
    name               VARCHAR(255)        NOT NULL,
    formatted_address  TEXT                NOT NULL,
    icon               VARCHAR(255)        NOT NULL,
    rating             REAL CHECK (rating >= 1.0 AND rating <= 5.0),
    user_ratings_total INT                 NOT NULL,
    price_level        INT CHECK (price_level >= 0 AND price_level <= 4),
    vicinity           TEXT                NOT NULL,
    permanently_closed BOOLEAN             NOT NULL,
    business_status    VARCHAR(255)        NOT NULL,
    location_latitude  FLOAT               NOT NULL,
    location_longitude FLOAT               NOT NULL,
    types              TEXT                NOT NULL, -- カンマ区切りでタイプ情報を格納
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

-- Table: spots
CREATE TABLE spots
(
    id         CHAR(26) PRIMARY KEY,
    gm_id      VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (gm_id) REFERENCES gm_places (place_id) ON DELETE CASCADE
);

-- Table: messages
CREATE TABLE messages
(
    id         CHAR(26) PRIMARY KEY,
    user_id    CHAR(26) NOT NULL,
    spot_id    CHAR(26) NOT NULL,
    photo_url  TEXT     NOT NULL,
    content    TEXT     NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (spot_id) REFERENCES spots (id) ON DELETE CASCADE
);

-- Table: stamps
CREATE TABLE stamps
(
    id         CHAR(26) PRIMARY KEY,
    type       VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Table: reactions
CREATE TABLE reactions
(
    id         CHAR(26) PRIMARY KEY,
    user_id    CHAR(26) NOT NULL,
    message_id CHAR(26) NOT NULL,
    stamp_id   CHAR(26) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (message_id) REFERENCES messages (id) ON DELETE CASCADE,
    FOREIGN KEY (stamp_id) REFERENCES stamps (id) ON DELETE CASCADE
);
