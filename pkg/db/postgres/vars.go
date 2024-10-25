package postgres

const (
	queryInitContainers = `CREATE TABLE IF NOT EXISTS containers (
		id bigint NOT NULL,
		name text NOT NULL,
		document_id bigint NULL,
		link_small text NOT NULL,
		link_bit text NOT NULL,
		PRIMARY KEY (id)
	  )`
)
