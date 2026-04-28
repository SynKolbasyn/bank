-- +goose Up
-- +goose statementBegin

CREATE TYPE PAYMENT_STATUS AS ENUM ('pending', 'processing', 'canceled', 'success');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT UUIDV7(),
    email VARCHAR(256) NOT NULL UNIQUE CHECK (email != ''::VARCHAR(256)),
    password_hash VARCHAR(256) NOT NULL CHECK (password_hash != ''::VARCHAR(256)),
    balance NUMERIC(17, 2) NOT NULL DEFAULT 1000000::NUMERIC(17, 2) CHECK (balance >= 0::NUMERIC(17, 2)),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT UUIDV7(),
    sender_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    recipient_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    amount NUMERIC(17, 2) NOT NULL CHECK (amount > 0::NUMERIC(17, 2)),
    status PAYMENT_STATUS NOT NULL DEFAULT 'pending'::PAYMENT_STATUS,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CHECK (sender_id != recipient_id)
);

CREATE FUNCTION UPDATE_updated_at_COLUMN()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_update_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION UPDATE_updated_at_COLUMN();

CREATE TRIGGER trg_payments_update_updated_at
BEFORE UPDATE ON payments
FOR EACH ROW
EXECUTE FUNCTION UPDATE_updated_at_COLUMN();

CREATE INDEX idx_users_email ON users(email);

-- +goose statementEnd


-- +goose Down
-- +goose statementBegin

DROP INDEX idx_users_email;
DROP TRIGGER trg_payments_update_updated_at;
DROP TRIGGER trg_users_update_updated_at;
DROP FUNCTION UPDATE_updated_at_COLUMN;
DROP TABLE payments;
DROP TABLE users;
DROP TYPE PAYMENT_STATUS;

-- +goose statementEnd
