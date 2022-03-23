CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table practices
(
    uuid4_field uuid not null default uuid_generate_v4(),
    str_field varchar not null,
    int_field integer not null,
    is_bool bool not null,
    time_field timestamp(0) not null,
    date_field date not null,
    easy_json jsonb not null
);
