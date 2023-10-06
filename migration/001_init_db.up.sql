CREATE TABLE burrow
(
    id           bigserial                            NOT NULL,
    uuid         UUID      DEFAULT uuid_generate_v4() NOT NULL,
    name         character varying(50),
    depth        float,
    width        float,
    occupied     BOOLEAN,
    age          INT,

    created_date timestamp DEFAULT now(),
    changed_date timestamp,
    deleted_date timestamp,
    CONSTRAINT burrow_pk PRIMARY KEY (id)
);
