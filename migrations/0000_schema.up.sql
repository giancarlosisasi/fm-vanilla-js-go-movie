BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SEQUENCE IF NOT EXISTS actors_id_seq;

CREATE TABLE "public"."actors" (
    "id" int4 NOT NULL DEFAULT nextval('actors_id_seq'::regclass),
    "first_name" text NOT NULL,
    "last_name" text NOT NULL,
    "image_url" text,
    PRIMARY KEY ("id")
);

CREATE SEQUENCE IF NOT EXISTS genres_id_seq;

CREATE TABLE "public"."genres" (
    "id" int4 NOT NULL DEFAULT nextval('genres_id_seq'::regclass),
    "name" text NOT NULL,
    PRIMARY KEY ("id")
);

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS keywords_id_seq;

-- Table Definition
CREATE TABLE "public"."keywords" (
    "id" int4 NOT NULL DEFAULT nextval('keywords_id_seq'::regclass),
    "word" text NOT NULL,
    PRIMARY KEY ("id")
);


-- Table Definition
CREATE TABLE "public"."movie_cast" (
    "movie_id" int4 NOT NULL,
    "actor_id" int4 NOT NULL,
    PRIMARY KEY ("movie_id","actor_id")
);

CREATE TABLE "public"."movie_genres" (
    "movie_id" int4 NOT NULL,
    "genre_id" int4 NOT NULL,
    PRIMARY KEY ("movie_id","genre_id")
);

CREATE TABLE "public"."movie_keywords" (
    "movie_id" int4 NOT NULL,
    "keyword_id" int4 NOT NULL,
    PRIMARY KEY ("movie_id","keyword_id")
);

CREATE SEQUENCE IF NOT EXISTS movies_id_seq;

-- Table Definition
CREATE TABLE "public"."movies" (
    "id" int4 NOT NULL DEFAULT nextval('movies_id_seq'::regclass),
    "tmdb_id" int4,
    "title" text NOT NULL,
    "tagline" text,
    "release_year" int4,
    "overview" text,
    "score" float4,
    "popularity" float4,
    "language" text,
    "poster_url" text,
    "trailer_url" text,
    PRIMARY KEY ("id")
);

-- Table Definition
CREATE TABLE "public"."user_movies" (
    "user_id" int4 NOT NULL,
    "movie_id" int4 NOT NULL,
    "relation_type" text NOT NULL CHECK (relation_type = ANY (ARRAY['favorite'::text, 'watchlist'::text])),
    "time_added" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("user_id","movie_id","relation_type")
);

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "name" text NOT NULL,
    "email" text NOT NULL,
    "password_hashed" text NOT NULL,
    "last_login" timestamp,
    "time_created" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "time_confirmed" timestamp,
    "time_deleted" timestamp,
    PRIMARY KEY ("id")
);

-- Indices
CREATE UNIQUE INDEX genres_name_key ON public.genres USING btree (name);


-- Indices
CREATE UNIQUE INDEX keywords_word_key ON public.keywords USING btree (word);
ALTER TABLE "public"."movie_cast" ADD FOREIGN KEY ("actor_id") REFERENCES "public"."actors"("id") ON DELETE CASCADE;
ALTER TABLE "public"."movie_cast" ADD FOREIGN KEY ("movie_id") REFERENCES "public"."movies"("id") ON DELETE CASCADE;


-- Indices
CREATE INDEX movie ON public.movie_cast USING btree (movie_id);
CREATE INDEX actor ON public.movie_cast USING btree (actor_id);
ALTER TABLE "public"."movie_genres" ADD FOREIGN KEY ("movie_id") REFERENCES "public"."movies"("id") ON DELETE CASCADE;
ALTER TABLE "public"."movie_genres" ADD FOREIGN KEY ("genre_id") REFERENCES "public"."genres"("id") ON DELETE CASCADE;
ALTER TABLE "public"."movie_keywords" ADD FOREIGN KEY ("movie_id") REFERENCES "public"."movies"("id") ON DELETE CASCADE;
ALTER TABLE "public"."movie_keywords" ADD FOREIGN KEY ("keyword_id") REFERENCES "public"."keywords"("id") ON DELETE CASCADE;


-- Indices
CREATE UNIQUE INDEX unique_tmdb_id ON public.movies USING btree (tmdb_id);
ALTER TABLE "public"."user_movies" ADD FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE;
ALTER TABLE "public"."user_movies" ADD FOREIGN KEY ("movie_id") REFERENCES "public"."movies"("id") ON DELETE CASCADE;


-- Indices
CREATE UNIQUE INDEX users_email_key ON public.users USING btree (email);

COMMIT;