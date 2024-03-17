create table roles (
    id int primary key,
    name varchar(50) not null
);

create table users (
    id uuid default gen_random_uuid() not null primary key,
    first_name varchar(50) not null,
    last_name varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(255) not null,
    role_id int default 1 references roles(id),
    created_at   timestamp with time zone default now(),
    updated_at   timestamp with time zone default now()
);

create type gender as enum ('male', 'female');

create table actors (
    id uuid default gen_random_uuid() not null primary key,
    name varchar(50) not null,
    gender gender not null,
    birth_day timestamp without time zone not null
);

create table films(
    id uuid default gen_random_uuid() not null primary key,
    name varchar(100) not null,
    description varchar(1000) not null,
    release_date timestamp without time zone not null,
    rating numeric(3,1) check (rating >= 0 and rating <= 10)
);

create table films_actors (
    id uuid default gen_random_uuid() not null,
    film_id uuid not null references films(id) on delete cascade,
    actor_id uuid not null references actors(id) on delete cascade
);

create table user_tokens (
    id uuid default gen_random_uuid() not null,
    user_id uuid not null references users(id) on delete cascade,
    token varchar(255) not null,
    refresh_token varchar(255) not null,
    expires_at timestamp with time zone default now(),
    created_at timestamp with time zone default now()
);