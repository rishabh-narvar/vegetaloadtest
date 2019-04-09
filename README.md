# Vegeta Dynamic LoadTest

 - In the Config.yaml define your test criteria.
 - Run go build on main.go
 - run the main go binary

## Dyamic Fields : current support
 - uuid (string) 
 - timestamp (utc timestamp RFC3339 - https://golang.org/pkg/time/))

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
   
