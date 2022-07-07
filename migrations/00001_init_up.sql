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
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL CHECK (NAME != ''),
    abbreviation TEXT,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE chapters (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL CHECK (NAME != ''),
    num TEXT,

    r_id integer REFERENCES regulations
);

CREATE TABLE paragraphs (
    id SERIAL PRIMARY KEY,
    num INT NOT NULL CHECK (num >= 0),
    class TEXT,
    content TEXT NOT NULL CHECK (content != ''),
    c_id integer REFERENCES chapters
);

-- DATA --


COMMIT;