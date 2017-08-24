--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.4
-- Dumped by pg_dump version 9.6.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: cargo; Type: TABLE; Schema: public; Owner: paramey
--

CREATE TABLE cargo (
    id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE cargo OWNER TO paramey;

--
-- Name: cargo_id_seq; Type: SEQUENCE; Schema: public; Owner: paramey
--

CREATE SEQUENCE cargo_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE cargo_id_seq OWNER TO paramey;

--
-- Name: cargo_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paramey
--

ALTER SEQUENCE cargo_id_seq OWNED BY cargo.id;


--
-- Name: general_cargo; Type: TABLE; Schema: public; Owner: paramey
--

CREATE TABLE general_cargo (
    cargo_id integer NOT NULL,
    voyage_number integer NOT NULL
);


ALTER TABLE general_cargo OWNER TO paramey;

--
-- Name: port; Type: TABLE; Schema: public; Owner: paramey
--

CREATE TABLE port (
    id integer NOT NULL,
    name text
);


ALTER TABLE port OWNER TO paramey;

--
-- Name: ship_id_seq; Type: SEQUENCE; Schema: public; Owner: paramey
--

CREATE SEQUENCE ship_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE ship_id_seq OWNER TO paramey;

--
-- Name: ship; Type: TABLE; Schema: public; Owner: paramey
--

CREATE TABLE ship (
    id integer DEFAULT nextval('ship_id_seq'::regclass) NOT NULL,
    name character varying(30) NOT NULL
);


ALTER TABLE ship OWNER TO paramey;

--
-- Name: ship_schedule; Type: TABLE; Schema: public; Owner: paramey
--

CREATE TABLE ship_schedule (
    starting_point integer NOT NULL,
    final_destination integer NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    ship integer NOT NULL,
    voyage_number integer NOT NULL,
    CONSTRAINT orders_check CHECK ((starting_point <> final_destination)),
    CONSTRAINT ship_schedule_check CHECK ((start_date < end_date))
);


ALTER TABLE ship_schedule OWNER TO paramey;

--
-- Name: untitled_table_id_seq; Type: SEQUENCE; Schema: public; Owner: paramey
--

CREATE SEQUENCE untitled_table_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE untitled_table_id_seq OWNER TO paramey;

--
-- Name: untitled_table_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paramey
--

ALTER SEQUENCE untitled_table_id_seq OWNED BY port.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: paramey
--

CREATE TABLE users (
    id integer NOT NULL,
    name character varying(10) NOT NULL,
    password character varying(10) NOT NULL,
    permission character varying(10) DEFAULT '1'::character varying NOT NULL
);


ALTER TABLE users OWNER TO paramey;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: paramey
--

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users_id_seq OWNER TO paramey;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: paramey
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: cargo id; Type: DEFAULT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY cargo ALTER COLUMN id SET DEFAULT nextval('cargo_id_seq'::regclass);


--
-- Name: port id; Type: DEFAULT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY port ALTER COLUMN id SET DEFAULT nextval('untitled_table_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Data for Name: cargo; Type: TABLE DATA; Schema: public; Owner: paramey
--

COPY cargo (id, name) FROM stdin;
1	Oil
2	Tree
3	Gas
4	Food
5	Water
6	Gold
7	Weapon
8	Silver
11	Milk
13	Milky Way
\.


--
-- Name: cargo_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paramey
--

SELECT pg_catalog.setval('cargo_id_seq', 16, true);


--
-- Data for Name: general_cargo; Type: TABLE DATA; Schema: public; Owner: paramey
--

COPY general_cargo (cargo_id, voyage_number) FROM stdin;
3	1
1	1
5	1
1	3
2	3
3	3
4	3
\.


--
-- Data for Name: port; Type: TABLE DATA; Schema: public; Owner: paramey
--

COPY port (id, name) FROM stdin;
1	Kerch
2	Sevastopol
3	St. Petersburg
4	London
5	Amsterdam
6	Sydney
7	Melbourne
\.


--
-- Data for Name: ship; Type: TABLE DATA; Schema: public; Owner: paramey
--

COPY ship (id, name) FROM stdin;
1	Avrora
2	Gollandec
3	Pobeda
4	Beda
5	Hersones
7	Moby Dick
8	Titanic
\.


--
-- Name: ship_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paramey
--

SELECT pg_catalog.setval('ship_id_seq', 8, true);


--
-- Data for Name: ship_schedule; Type: TABLE DATA; Schema: public; Owner: paramey
--

COPY ship_schedule (starting_point, final_destination, start_date, end_date, ship, voyage_number) FROM stdin;
2	1	2017-07-15	2017-08-15	3	1
4	7	2017-08-01	2017-09-01	5	3
\.


--
-- Name: untitled_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paramey
--

SELECT pg_catalog.setval('untitled_table_id_seq', 7, true);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: paramey
--

COPY users (id, name, password, permission) FROM stdin;
1	Director	123456dir	1
3	Auditor	123456aud	3
2	Client	123456cli	2
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: paramey
--

SELECT pg_catalog.setval('users_id_seq', 3, true);


--
-- Name: cargo cargo_name_key; Type: CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY cargo
    ADD CONSTRAINT cargo_name_key UNIQUE (name);


--
-- Name: cargo cargo_pkey; Type: CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY cargo
    ADD CONSTRAINT cargo_pkey PRIMARY KEY (id);


--
-- Name: ship ship_name_key; Type: CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY ship
    ADD CONSTRAINT ship_name_key UNIQUE (name);


--
-- Name: ship ship_pkey; Type: CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY ship
    ADD CONSTRAINT ship_pkey PRIMARY KEY (id);


--
-- Name: ship_schedule ship_schedule_voyage_number_key; Type: CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY ship_schedule
    ADD CONSTRAINT ship_schedule_voyage_number_key UNIQUE (voyage_number);


--
-- Name: port untitled_table_name_key; Type: CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY port
    ADD CONSTRAINT untitled_table_name_key UNIQUE (name);


--
-- Name: port untitled_table_pkey; Type: CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY port
    ADD CONSTRAINT untitled_table_pkey PRIMARY KEY (id);


--
-- Name: users users_name_key; Type: CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_name_key UNIQUE (name);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: general_cargo general_cargo_cargo_fkey; Type: FK CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY general_cargo
    ADD CONSTRAINT general_cargo_cargo_fkey FOREIGN KEY (cargo_id) REFERENCES cargo(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: general_cargo general_cargo_voyage_number_fkey; Type: FK CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY general_cargo
    ADD CONSTRAINT general_cargo_voyage_number_fkey FOREIGN KEY (voyage_number) REFERENCES ship_schedule(voyage_number) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: ship_schedule ship_schedule_final_destination_fkey; Type: FK CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY ship_schedule
    ADD CONSTRAINT ship_schedule_final_destination_fkey FOREIGN KEY (final_destination) REFERENCES port(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: ship_schedule ship_schedule_ship_fkey; Type: FK CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY ship_schedule
    ADD CONSTRAINT ship_schedule_ship_fkey FOREIGN KEY (ship) REFERENCES ship(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: ship_schedule ship_schedule_starting_point_fkey; Type: FK CONSTRAINT; Schema: public; Owner: paramey
--

ALTER TABLE ONLY ship_schedule
    ADD CONSTRAINT ship_schedule_starting_point_fkey FOREIGN KEY (starting_point) REFERENCES port(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

