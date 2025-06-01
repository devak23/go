--user
DROP TABLE IF EXISTS user;
CREATE TABLE user (
  id       INTEGER PRIMARY KEY AUTOINCREMENT,
  username VARCHAR(100),
  password VARCHAR(1000),
  email    VARCHAR(100)
);
INSERT INTO "user" VALUES (1, 'suraj', 'suraj', 'sapatil@live.com');

--category
DROP TABLE IF EXISTS category;
CREATE TABLE category (
  id   INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(1000) NOT NULL,
  user_id REFERENCES user (id)
);
INSERT INTO "category" VALUES (1, 'TaskApp', 1);

--status
DROP TABLE IF EXISTS status;
CREATE TABLE status (
  id     INTEGER PRIMARY KEY AUTOINCREMENT,
  status VARCHAR(50) NOT NULL
);
INSERT INTO "status" VALUES (1, 'COMPLETE');
INSERT INTO "status" VALUES (2, 'PENDING');
INSERT INTO "status" VALUES (3, 'DELETED');


--task
DROP TABLE IF EXISTS task;
CREATE TABLE task (
  id               INTEGER PRIMARY KEY AUTOINCREMENT,
  title            VARCHAR(100),
  content          TEXT,
  created_date     timestamp,
  last_modified_at timestamp,
  finish_date      timestamp,
  priority         INTEGER,
  cat_id REFERENCES category (id),
  task_status_id REFERENCES status (id),
  due_date         timestamp,
  user_id REFERENCES user (id),
  is_deleted       VARCHAR(1),
  hide             INT
);

INSERT INTO "task" VALUES (1, 'Publish on github', 'Publish the source of tasks and picsort on github', '2015-11-12 15:30:59', '2015-11-21 14:19:22', NULL, 3, 1, 1, NULL, 1, 'N', 0);

INSERT INTO "task" VALUES (4, 'gofmtall', 'The idea is to run gofmt -w file.go on every go file in the listing, *Edit turns out this is is difficult to do in golang **Edit barely 3 line bash script. ', '2015-11-12 16:58:31', '2015-11-14 10:42:14', NULL, 3, 1, 1, NULL, 1, 'N' ,0);

DROP TABLE IF EXISTS comments;
CREATE TABLE comments (
  id      INTEGER PRIMARY KEY AUTOINCREMENT,
  content ntext,
  taskID REFERENCES task(id),
  created DATETIME,
  user_id REFERENCES user (id)
);

DROP TABLE IF EXISTS files;
CREATE TABLE files (
  name         VARCHAR(1000) NOT NULL,
  autoName     VARCHAR(255) NOT null,
  user_id REFERENCES user (id),
  created_date timestamp
);