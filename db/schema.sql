create table if not exists users (
  id bigserial primary key,
  email text not null unique,
  created_at timestamptz not null default now()
);

create table if not exists notes (
  id bigserial primary key,
  user_id bigint not null references users(id) on delete cascade,
  body text not null,
  created_at timestamptz not null default now()
);
