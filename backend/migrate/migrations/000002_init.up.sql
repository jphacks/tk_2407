-- Table: users
CREATE TABLE users (
                       id CHAR(26) PRIMARY KEY,
                       username VARCHAR(255) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       created_at TIMESTAMPTZ DEFAULT NOW(),
                       updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Table: spots
CREATE TABLE spots (
                       id CHAR(26) PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       description TEXT,
                       photo_url TEXT,
                       google_map_place_id VARCHAR(255),
                       latitude DOUBLE PRECISION,
                       longitude DOUBLE PRECISION,
                       created_at TIMESTAMPTZ DEFAULT NOW(),
                       updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Table: messages
CREATE TABLE messages (
                          id CHAR(26) PRIMARY KEY,
                          user_id CHAR(26) NOT NULL,
                          spot_id CHAR(26) NOT NULL,
                          photo_url TEXT,
                          content TEXT,
                          created_at TIMESTAMPTZ DEFAULT NOW(),
                          updated_at TIMESTAMPTZ DEFAULT NOW(),
                          FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
                          FOREIGN KEY (spot_id) REFERENCES spots (id) ON DELETE CASCADE
);

-- Table: stamps
CREATE TABLE stamps (
                        id CHAR(26) PRIMARY KEY,
                        type VARCHAR(255) NOT NULL,
                        created_at TIMESTAMPTZ DEFAULT NOW(),
                        updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Table: reactions
CREATE TABLE reactions (
                           id CHAR(26) PRIMARY KEY,
                           user_id CHAR(26) NOT NULL,
                           message_id CHAR(26) NOT NULL,
                           stamp_id CHAR(26) NOT NULL,
                           created_at TIMESTAMPTZ DEFAULT NOW(),
                           FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
                           FOREIGN KEY (message_id) REFERENCES messages (id) ON DELETE CASCADE,
                           FOREIGN KEY (stamp_id) REFERENCES stamps (id) ON DELETE CASCADE
);
