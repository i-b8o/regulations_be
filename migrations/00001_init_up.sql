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
DROP TABLE IF EXISTS chapters;
DROP TABLE IF EXISTS regulations;

CREATE TABLE regulations (
    regulation_id SERIAL PRIMARY KEY,
    regulation_name TEXT NOT NULL
);

CREATE TABLE chapters (
    chapter_id SERIAL PRIMARY KEY,
    chapter_name TEXT NOT NULL,
    chapter_num TEXT NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    r_id integer REFERENCES regulations
);

CREATE TABLE paragraphs (
    paragraph_id INT NOT NULL PRIMARY KEY,
    paragraph_class TEXT,
    paragraph_text TEXT NOT NULL,
    c_id integer REFERENCES chapters
);

-- DATA --


COMMIT;