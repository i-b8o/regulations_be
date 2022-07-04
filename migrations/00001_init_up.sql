BEGIN

SET STATEMENT_TIMEOUtBEGIN;

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

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- TABLES --
DROP TABLE IF EXISTS regulations;
DROP TABLE IF EXISTS chapters;
DROP TABLE IF EXISTS paragraph;

CREATE TABLE regulations (
    regulation_id INT GENERATED ALWAYS AS IDENTITY,
    regulation_name TEXT NOT NULL,
    PRIMARY KEY(regulation_id)
);

CREATE TABLE chapters (
    chapter_id INT GENERATED ALWAYS AS IDENTITY,
    chapter_name TEXT NOT NULL,
    chapter_num TEXT NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    PRIMARY KEY(chapter_id)
    CONSTRAINT fk_regulation
        FOREIGN KEY(regulation_id) 
	    REFERENCES chapters(regulation_id)
);

CREATE TABLE paragraphs (
    paragraph_id NOT NULL,
    paragraph_class TEXT,
    paragraph_text TEXT NOT NULL,
    PRIMARY KEY(paragraph_id),
    CONSTRAINT fk_chapter
        FOREIGN KEY(chapter_id) 
	    REFERENCES chapters(chapter_id)
);

-- DATA --


COMMIT;