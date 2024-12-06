# This is an application to monitor data metrics


## Food Delivery DataBase:

### Schema:

```
restaraunts:
    RestarauntId: BIGINT
    RestarauntName: TEXT

foods:
    FoodId: BIGINT
    FoodName: TEXT
    RestarauntId: BIGINT

orders: (Partitioned based on city)
    OrderId: BIGINT
    RestarauntId: BIGINT
    UserId: BIGINT
    City: TEXT
    OrdereTime: TIMESTAMP
    Cost: INT

ordered_foods: (Partitioned based on city)
    OrderedFoodId: BIGINT
    FoodId: BIGINT
    OrderId: BIGINT
    City: TEXT

users:
    UserId: BIGINT
    UserName: TEXT
```

### Indexes:
```
foods:
    FoodName
    RestarauntId

orders:
    RestarauntId
    City
    OrderTime

ordered_foods:
    FoodId
    OrderId
    City
```


## Metrics DataBase:

### Schema:

```
queries:
    QueryId: INT
    QueryDescription: TEXT
    Query: TEXT

query_parameters:
    ParameterId: INT
    QueryId: INT
    ParameterName: TEXT
    ParameterType: TEXT
    Order: INT
```


## APIs:

### Create a new metric:
```
POST /api/v1/metrics
    Request Body:
        description
        query
        queryParameters
    Response:
        metricId
        description
        query
```

### Fetch all metrics:
```
GET /api/v1/metrics
    Request Body:
    Response:
        metricId
        description
        query
        parameters
```

### Fetch metric:
```
GET /api/v1/metrics/{metricId}
    Request Body:
    Response:
        metricId
        description
        query
        parameters
```

### Execute a query:
```
POST /api/v1/metrics/{metricId}
    Request Body:
        // Parameters for the metric
    Response:
        result
```


<!-- ### Update a metric:
```
PUT /api/v1/metrics/{metricId}
    Request Body:
        description (optional)
        query (optional)
    Response:
        metricId
        description
        query
``` -->

<!-- ### Delete a metric:
```
DELETE /api/v1/metrics/{metricId}
``` -->

