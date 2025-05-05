
```shell
echo "Hello World" | vector
```
Example output:
```shell
2025-05-05T15:14:37.298123Z  INFO vector::app: Log level is enabled. level="info"
2025-05-05T15:14:37.301125Z  INFO vector::app: Loading configs. paths=["/etc/vector/vector.yaml"]
2025-05-05T15:14:37.306880Z  INFO vector::sources::file_descriptors: Capturing stdin.
2025-05-05T15:14:37.306886Z  INFO vector::topology::running: Running healthchecks.
2025-05-05T15:14:37.306997Z  INFO vector: Vector has started. debug="false" version="0.46.1" arch="aarch64" revision="9a19e8a 2025-04-14 18:36:30.707862743"
2025-05-05T15:14:37.307004Z  INFO vector::app: API is disabled, enable by setting `api.enabled` to `true` and use commands like `vector top`.
2025-05-05T15:14:37.307107Z  INFO vector::topology::builder: Healthcheck passed.
{"host":"72b632e4ac46","message":"Hello World","source_type":"stdin","timestamp":"2025-05-05T15:14:37.307221510Z"}
2025-05-05T15:14:37.307589Z  INFO vector_common::shutdown: All sources have finished.
2025-05-05T15:14:37.307595Z  INFO vector_common::shutdown: All sources have finished.
2025-05-05T15:14:37.307603Z  INFO vector::app: All sources have finished.
2025-05-05T15:14:37.307610Z  INFO vector: Vector has stopped.
```