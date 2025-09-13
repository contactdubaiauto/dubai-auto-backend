-- 0001	-	add status to users table
alter table users add column status int not null default 1;
alter table temp_users add column status int not null default 1;
alter table users add column type int not null default 1;
alter table temp_users add column type int not null default 1;
