CREATE TABLE IF NOT EXISTS engines
  ( id BIGSERIAL PRIMARY KEY,
    user_id int NOT NULL REFERENCES users(id),
    name varchar NOT NULL,
    sort varchar NOT NULL,
    start_date timestamp with time zone,
    end_date timestamp with time zone,
    UNIQUE (user_id, name)
  );

INSERT INTO engines SELECT * FROM preference_sets;

DROP TABLE preference_sets;