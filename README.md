# go-worker-pattern

A simple Go example of the worker-pool pattern. Multiple workers process jobs concurrently from a buffered queue.

## Run

```bash
go run .
```

The program queues 40 jobs and starts one more worker than the number of available CPU cores. Press **Ctrl+C** to stop the workers.
