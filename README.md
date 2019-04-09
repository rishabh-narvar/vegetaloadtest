# Vegeta Dynamic LoadTest

 - In the Config.yaml define your test criteria.
 - Run go build on main.go
 - run the main go binary

Unit test
- coverage low. But basic yaml and json parsing have unit test coverage

TBD
 - Benchmark
 - For now a `uuid` is assumed as the dybamic value, but types per field should be configurable  - DONE
   - order_id | guid
   - order_date | date
   - customer_name | string
 - look at benchmarks for viper (not published on GIT) , may induce latency.
   
