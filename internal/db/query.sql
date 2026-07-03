-- name: ListProjects :many
SELECT name
FROM projects
ORDER BY name ASC;

-- name: ListOngoingProjects :many
SELECT project_name AS name
FROM project_timers
WHERE stop_seconds = -1
ORDER BY name ASC;

-- name: ListStoppedProjects :many
SELECT DISTINCT t1.project_name
FROM project_timers t1
LEFT OUTER JOIN project_timers t2
ON t1.project_name = t2.project_name
AND t2.stop_seconds = -1
WHERE t2.project_name IS NULL
ORDER BY t1.project_name;

-- name: GetProject :many
SELECT p.name, t.start_seconds, t.stop_seconds
FROM projects AS p
JOIN project_timers AS t ON p.name = t.project_name
WHERE p.name = ?
ORDER BY
    t.start_seconds DESC,
    t.stop_seconds DESC;

-- name: StartTimer :exec
INSERT INTO project_timers (
	project_name, start_seconds, stop_seconds
) VALUES (
	?, unixepoch('now'), -1
);

-- name: StopTimer :one
UPDATE project_timers
SET stop_seconds = unixepoch('now')
WHERE project_name = ? AND stop_seconds = -1
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE name = ?;

-- name: DeleteTimers :exec
DELETE FROM project_timers
WHERE project_name = ?;
