CREATE TABLE IF NOT EXISTS Category (
	id_category SERIAL,
	name VARCHAR(20) NOT NULL,
	nsfw BOOLEAN NOT NULL,
	PRIMARY KEY (id_category)
);

CREATE TABLE IF NOT EXISTS Board (
	id_board SERIAL,
	code VARCHAR(3) UNIQUE NOT NULL,
	name VARCHAR(20) NOT NULL,
	id_category INT NOT NULL,
	PRIMARY KEY (id_board),
	FOREIGN KEY (id_category) REFERENCES Category (id_category)
);

CREATE TABLE IF NOT EXISTS Thread (
	id_thread SERIAL,
	filepath TEXT,
	subject VARCHAR(20),
	username VARCHAR(20),
	timestampp BIGINT,
	commenta TEXT,
	reply_count SMALLINT NOT NULL,
	image_count SMALLINT NOT NULL,
	is_archived BOOLEAN NOT NULL,
	is_pinned BOOLEAN NOT NULL,
	code VARCHAR(3) NOT NULL,
	PRIMARY KEY (id_thread),
	FOREIGN KEY (code) REFERENCES Board (code)
);

CREATE TABLE IF NOT EXISTS Reply (
	id_reply SERIAL,
	filepath TEXT,
	username VARCHAR(20),
	timestampp BIGINT,
	commenta TEXT,
	id_thread INT NOT NULL,
	PRIMARY KEY (id_reply),
	FOREIGN KEY (id_thread) REFERENCES Thread (id_thread)
);