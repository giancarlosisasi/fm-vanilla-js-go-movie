BEGIN;

DROP EXTENSION IF EXISTS "uuid-ossp";
DROP TABLE IF EXISTS "public"."actors";
DROP TABLE IF EXISTS "public"."genres";
DROP TABLE IF EXISTS "public"."keywords";
DROP TABLE IF EXISTS "public"."movie_cast";
DROP TABLE IF EXISTS "public"."movie_genres";
DROP TABLE IF EXISTS "public"."movie_keywords";
DROP TABLE IF EXISTS "public"."movies";
DROP TABLE IF EXISTS "public"."user_movies";
DROP TABLE IF EXISTS "public"."users";
DROP SEQUENCE IF EXISTS "actors_id_seq";
DROP SEQUENCE IF EXISTS "genres_id_seq";
DROP SEQUENCE IF EXISTS "keywords_id_seq";
DROP SEQUENCE IF EXISTS "movie_cast_id_seq";
DROP SEQUENCE IF EXISTS "movie_genres_id_seq";
DROP SEQUENCE IF EXISTS "movie_keywords_id_seq";
DROP SEQUENCE IF EXISTS "movies_id_seq";
DROP SEQUENCE IF EXISTS "user_movies_id_seq";
DROP SEQUENCE IF EXISTS "users_id_seq";

COMMIT;