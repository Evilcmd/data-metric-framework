-- +goose Up
CREATE TABLE orders (
    OrderId BIGINT,
    RestarauntId INT,
    UserId INT,
    City TEXT,
    OrderTime TIMESTAMP,
    Cost INT,
    PRIMARY KEY (OrderId, City),
    FOREIGN KEY (RestarauntId) REFERENCES restaraunts(RestarauntId) ON DELETE CASCADE,
    FOREIGN KEY (UserId) REFERENCES users(UserId) ON DELETE CASCADE
) PARTITION BY LIST (City);

CREATE TABLE orders_bangalore PARTITION OF orders FOR VALUES IN ('Bangalore');
CREATE TABLE orders_delhi PARTITION OF orders FOR VALUES IN ('Delhi');
CREATE TABLE orders_chennai PARTITION OF orders FOR VALUES IN ('Chennai');
CREATE TABLE orders_hyderabad PARTITION OF orders FOR VALUES IN ('Hyderbad');
CREATE TABLE orders_gurugram PARTITION OF orders FOR VALUES IN ('Gurugram');
CREATE TABLE orders_kolkata PARTITION OF orders FOR VALUES IN ('Kolkata');
CREATE TABLE orders_mumbai PARTITION OF orders FOR VALUES IN ('Mumbai');
CREATE TABLE orders_other PARTITION OF orders DEFAULT;

CREATE INDEX on orders (OrderTime);
CREATE INDEX on orders (City);
CREATE INDEX on orders (RestarauntId);

-- +goose Down
DROP TABLE orders CASCADE;