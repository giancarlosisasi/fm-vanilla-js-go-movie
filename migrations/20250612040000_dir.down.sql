-- Delete all data from tables while preserving structure
TRUNCATE TABLE "public"."actors" CASCADE;
TRUNCATE TABLE "public"."genres" CASCADE;
TRUNCATE TABLE "public"."keywords" CASCADE;
TRUNCATE TABLE "public"."movie_cast" CASCADE;
TRUNCATE TABLE "public"."movie_genres" CASCADE;
TRUNCATE TABLE "public"."movie_keywords" CASCADE;
TRUNCATE TABLE "public"."movies" CASCADE;
TRUNCATE TABLE "public"."user_movies" CASCADE;
TRUNCATE TABLE "public"."users" CASCADE;