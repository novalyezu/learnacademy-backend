CREATE TABLE IF NOT EXISTS public.communities (
  id character varying(255) COLLATE pg_catalog."default" NOT NULL,
  user_id character varying(255) COLLATE pg_catalog."default" NOT NULL,
  name character varying(255) COLLATE pg_catalog."default" NOT NULL,
  slug character varying(255) COLLATE pg_catalog."default" NOT NULL,
  thumbnail text COLLATE pg_catalog."default" NOT NULL,
  short_description character varying(255) COLLATE pg_catalog."default" NOT NULL,
  description text COLLATE pg_catalog."default" NOT NULL,
  is_public boolean NOT NULL,
  total_member integer NOT NULL DEFAULT 0,
  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now(),
  CONSTRAINT communities_pkey PRIMARY KEY (id),
  CONSTRAINT slug UNIQUE (slug),
  CONSTRAINT user_id FOREIGN KEY (user_id)
      REFERENCES public.users (id) MATCH SIMPLE
      ON UPDATE CASCADE
      ON DELETE CASCADE
);
