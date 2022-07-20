--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4 (Debian 14.4-1.pgdg110+1)
-- Dumped by pg_dump version 14.4 (Debian 14.4-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: booking_list; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.booking_list (
    id integer NOT NULL,
    table_id integer NOT NULL,
    users_id integer NOT NULL,
    table_num integer NOT NULL,
    booking_date date NOT NULL,
    booking_time time without time zone NOT NULL
);


ALTER TABLE public.booking_list OWNER TO postgres;

--
-- Name: booking_list_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.booking_list_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.booking_list_id_seq OWNER TO postgres;

--
-- Name: booking_list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.booking_list_id_seq OWNED BY public.booking_list.id;


--
-- Name: restaurants_list; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.restaurants_list (
    id integer NOT NULL,
    rest_name character varying(255) NOT NULL,
    av_time time without time zone NOT NULL,
    av_check integer NOT NULL
);


ALTER TABLE public.restaurants_list OWNER TO postgres;

--
-- Name: restaurants_list_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.restaurants_list_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.restaurants_list_id_seq OWNER TO postgres;

--
-- Name: restaurants_list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.restaurants_list_id_seq OWNED BY public.restaurants_list.id;


--
-- Name: restaurants_tables; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.restaurants_tables (
    id integer NOT NULL,
    restaurant_id integer NOT NULL,
    table_id integer NOT NULL
);


ALTER TABLE public.restaurants_tables OWNER TO postgres;

--
-- Name: restaurants_tables_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.restaurants_tables_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.restaurants_tables_id_seq OWNER TO postgres;

--
-- Name: restaurants_tables_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.restaurants_tables_id_seq OWNED BY public.restaurants_tables.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: table_list; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.table_list (
    id integer NOT NULL,
    table_num integer NOT NULL,
    max_persons integer NOT NULL,
    free_tables integer NOT NULL
);


ALTER TABLE public.table_list OWNER TO postgres;

--
-- Name: table_list_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.table_list_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.table_list_id_seq OWNER TO postgres;

--
-- Name: table_list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.table_list_id_seq OWNED BY public.table_list.id;


--
-- Name: user_list; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_list (
    id integer NOT NULL,
    users_name character varying(255) NOT NULL,
    phone character varying(255) NOT NULL
);


ALTER TABLE public.user_list OWNER TO postgres;

--
-- Name: user_list_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_list_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_list_id_seq OWNER TO postgres;

--
-- Name: user_list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_list_id_seq OWNED BY public.user_list.id;


--
-- Name: booking_list id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_list ALTER COLUMN id SET DEFAULT nextval('public.booking_list_id_seq'::regclass);


--
-- Name: restaurants_list id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.restaurants_list ALTER COLUMN id SET DEFAULT nextval('public.restaurants_list_id_seq'::regclass);


--
-- Name: restaurants_tables id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.restaurants_tables ALTER COLUMN id SET DEFAULT nextval('public.restaurants_tables_id_seq'::regclass);


--
-- Name: table_list id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.table_list ALTER COLUMN id SET DEFAULT nextval('public.table_list_id_seq'::regclass);


--
-- Name: user_list id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_list ALTER COLUMN id SET DEFAULT nextval('public.user_list_id_seq'::regclass);


--
-- Data for Name: booking_list; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.booking_list (id, table_id, users_id, table_num, booking_date, booking_time) FROM stdin;
\.


--
-- Data for Name: restaurants_list; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.restaurants_list (id, rest_name, av_time, av_check) FROM stdin;
1	Каравелла	00:30:00	2000
2	Молодость	00:15:00	1000
3	Мясо и Салат	01:00:00	1500
\.


--
-- Data for Name: restaurants_tables; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.restaurants_tables (id, restaurant_id, table_id) FROM stdin;
1	1	1
2	1	2
3	1	3
4	2	4
5	3	5
6	3	6
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schema_migrations (version, dirty) FROM stdin;
1	f
\.


--
-- Data for Name: table_list; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.table_list (id, table_num, max_persons, free_tables) FROM stdin;
1	6	4	6
2	2	3	2
3	2	2	2
4	3	3	3
5	2	8	2
6	4	3	4
\.


--
-- Data for Name: user_list; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_list (id, users_name, phone) FROM stdin;
\.


--
-- Name: booking_list_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.booking_list_id_seq', 1, false);


--
-- Name: restaurants_list_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.restaurants_list_id_seq', 3, true);


--
-- Name: restaurants_tables_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.restaurants_tables_id_seq', 6, true);


--
-- Name: table_list_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.table_list_id_seq', 6, true);


--
-- Name: user_list_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_list_id_seq', 1, false);


--
-- Name: booking_list booking_list_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.booking_list
    ADD CONSTRAINT booking_list_id_key UNIQUE (id);


--
-- Name: restaurants_list restaurants_list_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.restaurants_list
    ADD CONSTRAINT restaurants_list_id_key UNIQUE (id);


--
-- Name: restaurants_tables restaurants_tables_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.restaurants_tables
    ADD CONSTRAINT restaurants_tables_id_key UNIQUE (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: table_list table_list_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.table_list
    ADD CONSTRAINT table_list_id_key UNIQUE (id);


--
-- Name: user_list user_list_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_list
    ADD CONSTRAINT user_list_id_key UNIQUE (id);


--
-- Name: user_list user_list_phone_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_list
    ADD CONSTRAINT user_list_phone_key UNIQUE (phone);


--
-- PostgreSQL database dump complete
--

