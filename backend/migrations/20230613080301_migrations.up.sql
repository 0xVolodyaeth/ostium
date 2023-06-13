CREATE TABLE Bets (
    id BIGINT PRIMARY KEY,
    long_address CHAR(42) NOT NULL,
    short_address CHAR(42),
    amount BIGINT NOT NULL,
    expiration BIGINT NOT NULL,
    created_at BIGINT NOT NULL,
    opening_price BIGINT NOT NULL,
    is_active BOOLEAN NOT NULL,
    withdrawn BOOLEAN NOT NULL,
    winner CHAR(42) NOT NULL
);

