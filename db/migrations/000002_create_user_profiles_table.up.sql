CREATE TABLE user_profiles(
    id BIGSERIAL NOT NULL,
    profile_name VARCHAR(300) NULL,
    phone_number VARCHAR(30) NOT NULL,
    email VARCHAR(100) NOT NULL,
    gender VARCHAR(10) NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    CONSTRAINT user_profiles_pk PRIMARY KEY(id),
    CONSTRAINT user_profiles_phone_number_un UNIQUE(phone_number),
    CONSTRAINT user_profiles_email_un UNIQUE(email)
)