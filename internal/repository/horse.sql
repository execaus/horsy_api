-- GetRelativeHorses
WITH RECURSIVE relatives AS (
    -- стартовая лошадь
    SELECT
        h.*,
        ARRAY[h.id] AS visited_ids
    FROM app.horses h
    WHERE h.id = $1

    UNION ALL

    -- родители и потомки, которых ещё не посещали
    SELECT
        h.*,
        r.visited_ids || h.id
    FROM app.horses h
             JOIN relatives r
                  ON h.id = r.sire
                      OR h.id = r.dam
                      OR h.sire = r.id
                      OR h.dam = r.id
    WHERE NOT h.id = ANY (r.visited_ids)
)
SELECT *
FROM relatives
WHERE id <> $1;
