-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE escalateservice.update_user_scores()
LANGUAGE plpgsql
AS $$
BEGIN
    -- Crear táboa temporal para acumular puntuacións
    CREATE TEMPORARY TABLE IF NOT EXISTS puntuacions_temp (
        username VARCHAR(255) PRIMARY KEY,
        postScore_average DECIMAL DEFAULT 0.0,
        total_followers INT DEFAULT 0
    ) ON COMMIT DROP;

    -- Calcular media das puntuacions dos posts (corrixido)
    INSERT INTO puntuacions_temp (username, postScore_average)
    SELECT
        username,
        CASE
            WHEN COUNT(*) > 0 THEN SUM(score)::DECIMAL / COUNT(*)
            ELSE 0
        END AS postScore_average
    FROM escalateservice.posts
    GROUP BY username
    ON CONFLICT (username)
    DO UPDATE SET
        postScore_average = EXCLUDED.postScore_average;

    -- Calcular followers (corrixido)
    INSERT INTO puntuacions_temp (username, total_followers)
    SELECT followee, COUNT(*)
    FROM escalateservice.follows
    GROUP BY followee
    ON CONFLICT (username) 
    DO UPDATE SET 
        total_followers = EXCLUDED.total_followers;

    -- Actualizar puntuación final na táboa de users
    UPDATE escalateservice.users u
    SET score =
        800 +
        COALESCE(t.postScore_average, 0.0) +
        COALESCE(t.total_followers, 0) * 10
    FROM puntuacions_temp t
    WHERE u.username = t.username;
    
    -- A táboa temporal elimínase automaticamente (ON COMMIT DROP)
END;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS escalateservice.update_user_scores();
-- +goose StatementEnd