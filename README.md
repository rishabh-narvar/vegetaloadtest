# Vegeta Dynamic LoadTest

 - In the Config.yaml define your test criteria.
 - Run go build on main.go
 - run the go binary

Unit test
- coverage low. 

Benchmark
- TBD


- TBD
 - Optimize header formation
 - For now a `uuid` is assumed as the dybamic value, but types per field should be configurable
   - order_id | guid
   - order_date | date
   - customer_name | string
 - look at benchmarks for viper (not published on GIT) , may induce latency.
   
