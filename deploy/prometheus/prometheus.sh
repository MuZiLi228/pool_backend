docker run --name prometheus -d \
-p 9090:9090 \
-v ~/prometheus.yaml:/etc/prometheus/prometheus.yml \
prom/prometheus