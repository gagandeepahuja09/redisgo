* Make sure to initialize your repo before starting.
* go mod init github.com/gagandeepAhuja/redisgo
* https://redis.uptrace.dev/guide/pipelines.html#speeding-up-redis-with-pipelines
* Redis pipelines are a great way to optimize the performance.
* We can execute multiple commands using a single client-server round trip.

* Redis also supports: Geohash, Bloom filter, Hyperloglog.

* Do read about redis labs as well.

* Result method gives the value and Err for the command.
* Error method gives the Err for the command.

* Read this for a deep dive: https://redis.uptrace.dev/

* Map: Declaring using var vs declaring with make syntax. What's the difference?

* Improvements and Optimizations required in code:
* Move to using pipeline or txn.
* Time window related part, threshold, etc should be moved to env config.
* Strategy pattern to decide on whether to use Redis or some other strategy.
* Separate package which decides which key to be used where.
* Handling conversion b/w int and string in a better way.
* Adding load tests. You can use k6 scripts.

* NOTE:
* Do revise and read design patterns, context, panic.