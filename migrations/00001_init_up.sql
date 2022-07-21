BEGIN;

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public, extensions;
SET default_tablespace = '';
SET default_with_oids = FALSE;

SET SCHEMA 'public';

-- EXTENSIONS --

-- CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- TABLES --
DROP TABLE IF EXISTS paragraphs;
DROP TABLE IF EXISTS links;
DROP TABLE IF EXISTS chapters;
DROP TABLE IF EXISTS regulations;

CREATE TABLE regulations (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL CHECK (NAME != ''),
    abbreviation TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE chapters (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL CHECK (NAME != ''),
    order_num INT NOT NULL CHECK (order_num >= 0),
    num TEXT,
    r_id integer REFERENCES regulations
);

CREATE TABLE paragraphs (
    id SERIAL PRIMARY KEY,
    paragraph_id INT NOT NULL CHECK (paragraph_id >= 0),
    order_num INT NOT NULL CHECK (order_num >= 0),
    is_html BOOLEAN NOT NULL,
    is_table BOOLEAN NOT NULL,
    is_nft BOOLEAN NOT NULL,
    class TEXT,
    content TEXT NOT NULL CHECK (content != ''),
    c_id integer REFERENCES chapters
);

CREATE TABLE links (
    id INT NOT NULL UNIQUE,
    paragraph_num INT NOT NULL CHECK (paragraph_num >= 0),
    c_id integer REFERENCES chapters
);

-- DATA --


COMMIT;

SELECT * FROM regulations;
SELECT * FROM chapters;
SELECT * FROM paragraphs;
SELECT * FROM links;

SELECT id,c_id,paragraph_num FROM "links" ORDER BY c_id;
SELECT id,c_id,paragraph_num FROM "links" WHERE c_id = 2 ORDER BY c_id;
SELECT order_num, is_html, is_nft, is_table, class, content, c_id FROM "paragraphs" WHERE c_id = 1 ORDER BY order_num;