CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  name TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

create table if not exists notes (
  id bigserial primary key,
  user_id bigint not null references users(id) on delete cascade,
  body text not null,
  created_at timestamptz not null default now()
);
