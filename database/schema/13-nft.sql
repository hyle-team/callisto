-- +migrate Up
CREATE TABLE nft_events
(
    id                  SERIAL,
    event_type          TEXT  NOT NULL,
    nft_address         TEXT  NOT NULL,
    owner               TEXT  NOT NULL,
    new_owner           TEXT,
    validator           TEXT,
    new_validator       TEXT,
    amount              COIN[] NOT NULL DEFAULT '{}'
);
CREATE INDEX nft_events_id_index ON nft_events (id);

CREATE TABLE nfts
(
    address             TEXT    NOT NULL UNIQUE,
    owner               TEXT    NOT NULL,
    available_amount    COIN[]  NOT NULL DEFAULT '{}',
    vesting_period      INTEGER,
    reward_per_period   COIN[] NOT NULL DEFAULT '{}',
    last_vesting_time   INTEGER,
    vesting_counter     SMALLINT,
    denom               TEXT
);
CREATE INDEX nft_address_index ON nfts(address);



-- +migrate Down
DROP TABLE nft_events;
DROP TABLE nfts;
