
```shell
echo "Hello World" | vector
```

Example Output:
```shell

2025-05-05T15:17:07.985705Z  INFO vector::app: Loading configs. paths=["/etc/vector/vector.yaml"]
2025-05-05T15:17:07.988377Z  INFO vector::sources::file_descriptors: Capturing stdin.
2025-05-05T15:17:07.990881Z  INFO vector::topology::running: Running healthchecks.
2025-05-05T15:17:07.991017Z  INFO vector: Vector has started. debug="false" version="0.46.1" arch="aarch64" revision="9a19e8a 2025-04-14 18:36:30.707862743"
2025-05-05T15:17:07.991032Z  INFO vector::app: API is disabled, enable by setting `api.enabled` to `true` and use commands like `vector top`.
2025-05-05T15:17:07.991057Z  INFO vector::topology::builder: Healthcheck passed.
2025-05-05T15:17:07.991202Z  INFO vector_common::shutdown: All sources have finished.
2025-05-05T15:17:07.991210Z  INFO vector_common::shutdown: All sources have finished.
2025-05-05T15:17:07.991217Z  INFO vector::app: All sources have finished.
2025-05-05T15:17:07.991224Z  INFO vector: Vector has stopped.
Hello World
```