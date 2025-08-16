-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE public.user_story (
                                   id serial4 NOT NULL,
                                   created_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
                                   updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NULL,
                                   deleted_at timestamptz NULL,
                                   description text DEFAULT '' NULL,
                                   CONSTRAINT user_story_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
