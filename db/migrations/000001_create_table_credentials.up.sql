CREATE TABLE credentials(
    id BIGSERIAL NOT NULL,
    username VARCHAR(300) NOT NULL,
    password VARCHAR(300) NOT NULL,
    is_enabled BOOLEAN NOT NULL DEFAULT true,
    user_profile_id BIGINT,
    role VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT credentials_pk PRIMARY KEY(id),
    CONSTRAINT credentials_username_un UNIQUE(username)
)