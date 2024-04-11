DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS transfers;
-- accounts 被 entries 和 transfers 依赖，所以要最后删除
DROP TABLE IF EXISTS accounts;