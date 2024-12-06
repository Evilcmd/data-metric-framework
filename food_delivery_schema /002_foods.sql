-- +goose Up
create table foods(FoodId INT PRIMARY KEY, FoodName Text, RestarauntId INT, FOREIGN KEY(RestarauntId) REFERENCES restaraunts(RestarauntId) on delete cascade);

CREATE INDEX on foods (FoodName);
CREATE INDEX on foods (RestarauntId);

-- +goose Down
DROP TABLE foods CASCADE;