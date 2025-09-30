begin;
truncate table notes restart identity cascade;
truncate table users restart identity cascade;

insert into users (email) values
  ('test1@example.com'),
  ('test2@example.com');

insert into notes (user_id, body) values
  (1, 'hello'),
  (1, 'world'),
  (2, 'note-2-1');
commit;