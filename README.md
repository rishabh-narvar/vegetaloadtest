# Problem 
 - how do we load test APIs that require payload field values to change with every request? (ex OrderAPI)
 - how do we test requests that require dynamic headers? (ex testing shopify app with headers containing hmac of request payload)
 - yet use the power of vegeta? (https://github.com/tsenart/vegeta)

# Vegeta Dynamic LoadTest

 - In Config.yaml define your test criteria.
 - Copy json into payload.json (path is configurable)
 - Run go build on main.go
 - run the main go binary

## Dyamic Fields : current support
 - uuid (string) 
 - timestamp (utc timestamp RFC3339 - https://golang.org/pkg/time/)) (ex. carrier apis need scan times to be different)
 - epoch
 - epochnano

 ### Sample yaml
   ```yaml
   post-request-json-dynamic-fields: 
     guid: uuid #uuid string
     timestamp: timestamp #timestamp
     customer.firstName: uuid
   ```
  
  
## Unit test
- coverage low. But basic yaml and json parsing have unit test coverage

## TBD
 - Benchmark
 - look at benchmarks for viper (not published on GIT) , may induce latency.
   
