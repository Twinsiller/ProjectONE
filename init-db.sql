CREATE USER postgres WITH PASSWORD 'qwerty';
CREATE DATABASE ProjectONE OWNER postgres;
\c ProjectONE;

CREATE TABLE public.authors (
    id serial PRIMARY KEY,
    nickname character varying(50) NOT NULL,
    hash_password text NOT NULL,
    status boolean NOT NULL,
    access_level integer NOT NULL,
    firstname character varying(50),
    lastname character varying(50),
    created_at timestamp(0) with time zone DEFAULT now()
);

CREATE TABLE public.posts (
    id serial PRIMARY KEY,
    id_author integer NOT NULL,
    title character varying(50) NOT NULL,
    description text,
    date_publication timestamp with time zone DEFAULT now(),
    date_last_modified timestamp with time zone DEFAULT now(),
    likes integer DEFAULT 0,
    CONSTRAINT fk_author FOREIGN KEY (id_author)
        REFERENCES public.authors (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE public.comments (
    id serial PRIMARY KEY,
    id_author integer NOT NULL,
    comment_text text NOT NULL,
    date_publication timestamp with time zone DEFAULT now(),
    date_last_modified timestamp with time zone DEFAULT now(),
    id_post integer NOT NULL,
    CONSTRAINT fk_author_comment FOREIGN KEY (id_author)
        REFERENCES public.authors (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_post FOREIGN KEY (id_post)
        REFERENCES public.posts (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);