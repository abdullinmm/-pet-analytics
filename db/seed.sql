BEGIN;
TRUNCATE TABLE notes RESTART IDENTITY CASCADE;
TRUNCATE TABLE users RESTART IDENTITY CASCADE;

INSERT INTO users (email, name) VALUES
  ('test1@example.com', 'Test One'),
  ('test2@example.com', 'Test Two');

INSERT INTO notes (user_id, body) VALUES
  (1, 'hello'),
  (1, 'world'),
  (2, 'note-2-1');
COMMIT;
