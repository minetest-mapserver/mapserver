
# Prometheus monitoring

The mapserver exposes a prometheus endpoint at the `/metrics` path

An overview of the collected metrics:
* tiledb get/set histogram
* cache hit/misses
* mapblock parse duration histogram
* mapblock render duration histogram
* tile render duration histogram
* various counters
