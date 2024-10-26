package postgres

const (
	queryInitContainers = `CREATE TABLE IF NOT EXISTS containers (
		id serial PRIMARY KEY,
		name text NOT NULL,
		document_id bigint NOT NULL DEFAULT -1,
		link_small text NOT NULL,
		link_big text NOT NULL
	  )`
)
