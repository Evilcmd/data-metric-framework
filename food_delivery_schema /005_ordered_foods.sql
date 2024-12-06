-- +goose Up
CREATE TABLE ordered_foods (
    OrderedFoodId BIGINT,
    FoodId INT,
    OrderId BIGINT,
    City TEXT,
    PRIMARY KEY (OrderedFoodId, City),
    FOREIGN KEY (FoodId) REFERENCES foods(FoodId) ON DELETE CASCADE,
    FOREIGN KEY (OrderId, City) REFERENCES orders(OrderId, City) ON DELETE CASCADE
)PARTITION BY LIST (City);

CREATE TABLE ordered_food_bangalore PARTITION OF ordered_foods FOR VALUES IN ('Bangalore');
CREATE TABLE ordered_food_delhi PARTITION OF ordered_foods FOR VALUES IN ('Delhi');
CREATE TABLE ordered_food_chennai PARTITION OF ordered_foods FOR VALUES IN ('Chennai');
CREATE TABLE ordered_food_hyderbad PARTITION OF ordered_foods FOR VALUES IN ('Hyderbad');
CREATE TABLE ordered_food_gurugram PARTITION OF ordered_foods FOR VALUES IN ('Gurugram');
CREATE TABLE ordered_food_kolkata PARTITION OF ordered_foods FOR VALUES IN ('Kolkata');
CREATE TABLE ordered_food_mumbai PARTITION OF ordered_foods FOR VALUES IN ('Mumbai');
CREATE TABLE ordered_food_other PARTITION OF ordered_foods DEFAULT;

CREATE INDEX on ordered_foods (FoodId);
CREATE INDEX on ordered_foods (OrderId);
CREATE INDEX on ordered_foods (City);

-- +goose Down
DROP TABLE ordered_foods;