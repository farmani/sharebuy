# sharebuy

## Description

This is a simple script to calculate the number of shares to buy based on the amount of money you have and the price of the stock.

## Usage

```bash
$ python3 sharebuy.py
```

## Database Optimization

```shell
$ psql "postgresql://sharebuy-user:sharebuy-pass@localhost:5432/sharebuy?sslmode=disable"
```


https://pgtune.leopard.in.ua/

```pgsql
# DB Version: 15
# OS Type: linux
# DB Type: web
# Total Memory (RAM): 1 GB
# CPUs num: 4
# Connections num: 128
# Data Storage: ssd

ALTER SYSTEM SET
 max_connections = '128';
ALTER SYSTEM SET
 shared_buffers = '256MB';
ALTER SYSTEM SET
 effective_cache_size = '768MB';
ALTER SYSTEM SET
 maintenance_work_mem = '64MB';
ALTER SYSTEM SET
 checkpoint_completion_target = '0.9';
ALTER SYSTEM SET
 wal_buffers = '7864kB';
ALTER SYSTEM SET
 default_statistics_target = '100';
ALTER SYSTEM SET
 random_page_cost = '1.1';
ALTER SYSTEM SET
 effective_io_concurrency = '200';
ALTER SYSTEM SET
 work_mem = '1MB';
ALTER SYSTEM SET
 min_wal_size = '1GB';
ALTER SYSTEM SET
 max_wal_size = '4GB';
ALTER SYSTEM SET
 max_worker_processes = '4';
ALTER SYSTEM SET
 max_parallel_workers_per_gather = '2';
ALTER SYSTEM SET
 max_parallel_workers = '4';
ALTER SYSTEM SET
 max_parallel_maintenance_workers = '2';
 ```

postgresql.conf
```text
# DB Version: 15
# OS Type: linux
# DB Type: web
# Total Memory (RAM): 1 GB
# CPUs num: 4
# Connections num: 128
# Data Storage: ssd

max_connections = 128
shared_buffers = 256MB
effective_cache_size = 768MB
maintenance_work_mem = 64MB
checkpoint_completion_target = 0.9
wal_buffers = 7864kB
default_statistics_target = 100
random_page_cost = 1.1
effective_io_concurrency = 200
work_mem = 1MB
min_wal_size = 1GB
max_wal_size = 4GB
max_worker_processes = 4
max_parallel_workers_per_gather = 2
max_parallel_workers = 4
max_parallel_maintenance_workers = 2
```