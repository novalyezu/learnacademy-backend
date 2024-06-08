CREATE TABLE IF NOT EXISTS public.users(
  id character varying(255) COLLATE pg_catalog."default" NOT NULL,
  firstname character varying(255) COLLATE pg_catalog."default" NOT NULL,
  lastname character varying(255) COLLATE pg_catalog."default",
  email character varying(255) COLLATE pg_catalog."default" NOT NULL,
  password character varying(255) COLLATE pg_catalog."default" NOT NULL,
  created_at timestamp without time zone NOT NULL DEFAULT now(),
  updated_at timestamp without time zone NOT NULL DEFAULT now(),
  CONSTRAINT users_pkey PRIMARY KEY (id),
  CONSTRAINT email UNIQUE (email)
);
