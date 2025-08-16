-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE public.demo (
                                 id serial4 NOT NULL,
                                 "name" varchar DEFAULT '' NULL,
                                 CONSTRAINT demo_pk PRIMARY KEY (id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
