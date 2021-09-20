DROP TABLE IF EXISTS urlshort;

CREATE TABLE urlshort (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  url TEXT NOT NULL
);

CREATE INDEX shortened_index ON urlshort (shortened);

INSERT INTO 
    urlshort (url, name)
VALUES
    ('https://www.google.com','goog'),
    ('https://www.yahoo.com','yahoo'),
    ('https://www.bing.com','bing');