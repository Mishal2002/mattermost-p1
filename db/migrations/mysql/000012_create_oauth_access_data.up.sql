CREATE TABLE IF NOT EXISTS OAuthAccessData (
  Token varchar(26) NOT NULL,
  RefreshToken varchar(26) DEFAULT NULL,
  RedirectUri text,
  PRIMARY KEY (Token),
  KEY idx_oauthaccessdata_refresh_token (RefreshToken)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET @preparedStatement = (SELECT IF(
    (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
        WHERE table_name = 'OAuthAccessData'
        AND table_schema = DATABASE()
        AND column_name = 'ClientId'
    ) > 0,
    'SELECT 1',
    'ALTER TABLE OAuthAccessData ADD ClientId varchar(26) DEFAULT "";'
));

PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;

SET @preparedStatement = (SELECT IF(
    (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.statistics
        WHERE table_name = 'OAuthAccessData'
        AND table_schema = DATABASE()
        AND column_name = 'ClientId'
    ) > 0,
    'SELECT 1',
    'CREATE INDEX idx_oauthaccessdata_client_id ON OAuthAccessData(ClientId);'
));

PREPARE createIndexIfNotExists FROM @preparedStatement;
EXECUTE createIndexIfNotExists;
DEALLOCATE PREPARE createIndexIfNotExists;

SET @preparedStatement = (SELECT IF(
    (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
        WHERE table_name = 'OAuthAccessData'
        AND table_schema = DATABASE()
        AND column_name = 'UserId'
    ) > 0,
    'SELECT 1',
    'ALTER TABLE OAuthAccessData ADD UserId varchar(26) DEFAULT "";'
));

PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;

SET @preparedStatement = (SELECT IF(
    (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.statistics
        WHERE table_name = 'OAuthAccessData'
        AND table_schema = DATABASE()
        AND column_name = 'UserId'
    ) > 0,
    'SELECT 1',
    'CREATE INDEX idx_oauthaccessdata_user_id ON OAuthAccessData(UserId);'
));

PREPARE createIndexIfNotExists FROM @preparedStatement;
EXECUTE createIndexIfNotExists;
DEALLOCATE PREPARE createIndexIfNotExists;

SET @preparedStatement = (SELECT IF(
    (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS
        WHERE table_name = 'OAuthAccessData'
        AND table_schema = DATABASE()
        AND constraint_name = 'ClientId'
    ) > 0,
    'SELECT 1',
    'ALTER TABLE OAuthAccessData ADD CONSTRAINT ClientId UNIQUE (ClientId, UserId);'
));

PREPARE createUniqueKeyIfNotExists FROM @preparedStatement;
EXECUTE createUniqueKeyIfNotExists;
DEALLOCATE PREPARE createUniqueKeyIfNotExists;

SET @preparedStatement = (SELECT IF(
    (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
        WHERE table_name = 'OAuthAccessData'
        AND table_schema = DATABASE()
        AND column_name = 'ExpiresAt'
    ) > 0,
    'SELECT 1',
    'ALTER TABLE OAuthAccessData ADD ExpiresAt bigint(20) DEFAULT 0;'
));

PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;

SET @preparedStatement = (SELECT IF(
    (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.statistics
        WHERE table_name = 'OAuthAccessData'
        AND table_schema = DATABASE()
        AND index_name = 'idx_oauthaccessdata_auth_code'
    ) > 0,
    'DROP INDEX idx_oauthaccessdata_auth_code ON OAuthAccessData;',
    'SELECT 1'
));

PREPARE removeIndexIfExists FROM @preparedStatement;
EXECUTE removeIndexIfExists;
DEALLOCATE PREPARE removeIndexIfExists;

SET @preparedStatement = (SELECT IF(
    (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
        WHERE table_name = 'OAuthAccessData'
        AND table_schema = DATABASE()
        AND column_name = 'AuthCode'
    ) > 0,
    'ALTER TABLE OAuthAccessData DROP COLUMN AuthCode;',
    'SELECT 1'
));

PREPARE alterIfExists FROM @preparedStatement;
EXECUTE alterIfExists;
DEALLOCATE PREPARE alterIfExists;

SET @preparedStatement = (SELECT IF(
    (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS
        WHERE table_name = 'OAuthAccessData'
        AND table_schema = DATABASE()
        AND column_name = 'Scope'
    ) > 0,
    'SELECT 1',
    'ALTER TABLE OAuthAccessData ADD Scope varchar(128) DEFAULT "user";'
));

PREPARE alterIfNotExists FROM @preparedStatement;
EXECUTE alterIfNotExists;
DEALLOCATE PREPARE alterIfNotExists;

SET @preparedStatement = (SELECT IF(
    (
        SELECT COUNT(*) FROM INFORMATION_SCHEMA.statistics
        WHERE table_name = 'OAuthAccessData'
        AND table_schema = DATABASE()
        AND index_name = 'ClientId_2'
    ) > 0,
    'DROP INDEX ClientId_2 ON OAuthAccessData;',
    'SELECT 1'
));

PREPARE removeIndexIfExists FROM @preparedStatement;
EXECUTE removeIndexIfExists;
DEALLOCATE PREPARE removeIndexIfExists;
