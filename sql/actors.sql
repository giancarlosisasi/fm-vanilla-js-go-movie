-- name: GetActorByID :one
-- Get an actor by their ID
SELECT * FROM actors WHERE id = $1;