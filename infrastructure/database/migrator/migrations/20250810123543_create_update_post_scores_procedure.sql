-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE escalateservice.update_post_scores()
LANGUAGE plpgsql
AS $$
BEGIN
    -- Crear táboa temporal para acumular puntuacións
    CREATE TEMPORARY TABLE IF NOT EXISTS puntuacions_temp (
        post_id VARCHAR(255) PRIMARY KEY,
        total_likes INT DEFAULT 0,
        total_superlikes INT DEFAULT 0,
        total_revisiones INT DEFAULT 0
    ) ON COMMIT DROP;

    -- Calcular likes normais (👍 = 1 punto)
    INSERT INTO puntuacions_temp (post_id, total_likes)
    SELECT post_id, COUNT(*)
    FROM escalateservice.likeposts
    GROUP BY post_id
    ON CONFLICT (post_id) DO UPDATE SET total_likes = EXCLUDED.total_likes;

    -- Calcular superlikes (❤️ = 10 puntos)
    INSERT INTO puntuacions_temp (post_id, total_superlikes)
    SELECT post_id, COUNT(*)
    FROM escalateservice.superlikeposts
    GROUP BY post_id
    ON CONFLICT (post_id) DO UPDATE SET total_superlikes = EXCLUDED.total_superlikes;

    -- Calcular puntuación de revisións
    INSERT INTO puntuacions_temp (post_id, total_revisiones)
    SELECT
        post_id,
        SUM(CASE
            WHEN rating = 0 THEN -200  -- 👎
            WHEN rating = 1 THEN -100  -- 👎
            WHEN rating = 2 THEN 0     -- 😐
            WHEN rating = 3 THEN 100   -- ✅
            WHEN rating = 4 THEN 200   -- ✅
            WHEN rating = 5 THEN 300   -- ✅
        END) AS total_revisiones
    FROM escalateservice.reviews
    GROUP BY post_id
    ON CONFLICT (post_id) DO UPDATE SET total_revisiones = EXCLUDED.total_revisiones;

    -- Actualizar puntuación final na táboa de posts
    UPDATE escalateservice.posts p
    SET score =
        COALESCE(t.total_likes, 0) +
        COALESCE(t.total_superlikes * 10, 0) +
        COALESCE(t.total_revisiones, 0)
    FROM puntuacions_temp t
    WHERE p.post_id = t.post_id;

    -- A táboa temporal elimínase automaticamente ao rematar a transacción (ON COMMIT DROP)
END;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS escalateservice.update_post_scores();
-- +goose StatementEnd