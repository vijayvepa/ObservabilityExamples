# Prometheus Http Request Counter

- [Reference](https://prometheus.io/docs/tutorials/instrumenting_http_server_in_go/)

## Example Output

- Server Metrics
![alt text](image.png)
- [Prometheus Metrics](http://localhost:9000/graph?g0.expr=ping_requests_total&g0.tab=0&g0.display_mode=lines&g0.show_exemplars=0&g0.range_input=1h)

![alt text](image-1.png)

- Grafana

![alt text](image-2.png)