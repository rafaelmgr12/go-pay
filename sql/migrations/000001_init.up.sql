START TRANSACTION;
CREATE TABLE IF NOT EXISTS balances (
id VARCHAR(36) NOT NULL PRIMARY KEY,
user_id VARCHAR(36) NOT NULL,
amount DECIMAL(10, 2),
created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
COMMIT;

insert into balances (id, user_id, amount) values 
('ab596652-e526-4838-ab50-c0caa3d7488b', 'f2f0e0d1-e37e-45c3-ad06-e6c2a66544fc', 1000), 
('5231d5de-3157-41af-b1f6-950d8e12f0ec', '089557bc-ddf2-4ec5-8077-d8bf09fe3ddc', 1000);
