CREATE TABLE balance
(
    id          BIGSERIAL NOT NULL PRIMARY KEY,
    player_name VARCHAR   NOT NULL,
    balance     INTEGER   NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS balance_player ON balance (player_name);
