# This is an application to monitor data metrics


## DataBase Schema:

### Restaraunt:
```
RestarauntId: BIGINT
RestarauntName: TEXT
```

### Food:
- FoodId: BIGINT
- FoodName: TEXT
- RestarauntId: BIGINT

### Order:
- OrderId: BIGINT
- RestarauntId: BIGINT
- UserId: BIGINT
- Address: TEXT

### OrderedFood:
- OrderedFoodId: BIGINT
- FoodId: BIGINT
- OrderId: BIGINT

### User:
- UserId: BIGINT
- UserName: TEXT


## APIs:

### Create a new metric:
- POST /api/v1/metrics
    - Request Body:
        - description
        - query
    - Response:
        - metricId
        - description
        - query

### Update a metric:
- PUT /api/v1/metrics/{metricId}
    - Request Body:
        - description (optional)
        - query (optional)
    - Response:
        - metricId
        - description
        - query

### Delete a metric:
- DELETE /api/v1/metrics/{metricId}

### Execute a query:
- POST /api/v1/metrics/{metricId}
    - Request Body:
        - // Parameters for the query
    - Response:
        - result


