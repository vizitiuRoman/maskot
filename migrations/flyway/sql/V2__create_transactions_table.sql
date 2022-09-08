CREATE TABLE transactions
(
    id                     BIGSERIAL  NOT NULL PRIMARY KEY,
    withdraw               INTEGER    NOT NULL,
    deposit                INTEGER    NOT NULL,
    is_rollback            BOOLEAN    NOT NULL DEFAULT false,
    currency               VARCHAR(3) NOT NULL,
    player_name            VARCHAR    NOT NULL,
    transaction_ref        VARCHAR    NOT NULL,
    game_round_ref         VARCHAR,
    game_id                VARCHAR,
    reason                 VARCHAR,
    session_id             VARCHAR,
    session_alternative_id VARCHAR,
    spin_details_bet_type  VARCHAR,
    spin_details_win_type  VARCHAR,
    created_at             TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at             TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS transactions_uniq_idx ON transactions (transaction_ref);
