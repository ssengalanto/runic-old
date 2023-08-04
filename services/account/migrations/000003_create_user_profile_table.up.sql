CREATE TABLE account.user_profile (
    id UUID UNIQUE NOT NULL,
    user_id UUID UNIQUE NOT NULL,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    avatar VARCHAR NOT NULL,
    bio VARCHAR NOT NULL,
    date_of_birth TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_profile_pkey PRIMARY KEY (id),
    CONSTRAINT user_profile_user_id_fkey FOREIGN KEY(user_id) REFERENCES account.user(id) ON DELETE CASCADE
);

comment on column account.user_profile.user_id is 'One-to-one relationship with user table.';
