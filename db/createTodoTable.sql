GRANT ALL PRIVILEGES ON DATABASE docker TO docker;

CREATE TABLE todo (
	ID INT PRIMARY KEY NOT NULL,
	VALUE TEXT NOT NULL,
	CHECKED BOOLEAN NOT NULL DEFAULT false
);
