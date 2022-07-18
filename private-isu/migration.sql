set persist max_connections=256;

alter table comments add index idx_user_id (user_id);
alter table comments add index idx_post_id (post_id, created_at desc);

alter table posts add index idx_created_at (created_at desc);
alter table posts add index idx_user_id (user_id);
-- alter table posts drop index idx_user_id;
