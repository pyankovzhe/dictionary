-- +goose Up
-- +goose StatementBegin
CREATE table cards (
  id serial primary key,
  original varchar,
  translation text,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
