CREATE TABLE IF NOT EXISTS projects (
  name	TEXT	PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS project_timers (
	project_name	TEXT	NOT NULL REFERENCES projects(name),
	start_seconds	INTEGER	NOT NULL,
	stop_seconds	INTEGER	NOT NULL DEFAULT -1,

	-- Cannot have two active timers for the same project.
	-- This also prevents two recorded timers from stopping at the same time
	-- for the same project, which realistically should never happen anyways.
	PRIMARY KEY (project_name, stop_seconds)
);

-- Upsert the project row before inserting a timer so the FK is satisfied
-- without requiring two round-trips from the application.
CREATE TRIGGER IF NOT EXISTS upsert_project_on_timer
BEFORE INSERT ON project_timers
BEGIN
	INSERT INTO projects (name) VALUES (NEW.project_name)
	ON CONFLICT (name) DO NOTHING;
END;
