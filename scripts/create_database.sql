CREATE TABLE licenses
(
  id character varying(36) NOT NULL,
  version character varying(10),
  created_at timestamp without time zone NOT NULL,
  info json,
  CONSTRAINT licenses_pkey PRIMARY KEY (id),
  CONSTRAINT uk_id_version UNIQUE (id, version)
)