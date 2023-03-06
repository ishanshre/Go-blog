CREATE TABLE "users" (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    is_admin BOOLEAN DEFAULT 'f',
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    last_login TIMESTAMPTZ
);

CREATE TABLE "posts" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    pic VARCHAR(255),
    content VARCHAR(10000),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE "tags" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

CREATE TABLE "tag_post" (
    post_id INTEGER,
    tag_id INTEGER,
    PRIMARY KEY (post_id, tag_id),
    CONSTRAINT fk_post FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

CREATE TABLE "comments" (
    id SERIAL PRIMARY KEY,
    content VARCHAR(1000) NOT NULL,
    rating INTEGER NOT NULL CHECK (rating >0 AND rating < 6),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    post_id INTEGER,
    user_id INTEGER,
    CONSTRAINT fk_post FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);