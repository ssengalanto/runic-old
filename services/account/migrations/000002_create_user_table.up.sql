CREATE TYPE account.role_enum AS ENUM ('admin', 'moderator', 'user');

CREATE TABLE account.user (
    id UUID UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    active BOOLEAN NOT NULL,
    role account.role_enum NOT NULL,
    last_login_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT user_password_ck CHECK (char_length(password) >= 10)
);
