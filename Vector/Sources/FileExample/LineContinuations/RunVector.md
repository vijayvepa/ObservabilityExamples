```shell

vector
```

Example output:

```shell
2025-05-05T15:47:20.601897Z  INFO vector::app: Log level is enabled. level="info"
2025-05-05T15:47:20.606440Z  INFO vector::app: Loading configs. paths=["/etc/vector/vector.yaml"]
2025-05-05T15:47:20.615108Z  INFO vector::topology::running: Running healthchecks.
2025-05-05T15:47:20.615180Z  INFO vector: Vector has started. debug="false" version="0.46.1" arch="aarch64" revision="9a19e8a 2025-04-14 18:36:30.707862743"
2025-05-05T15:47:20.615186Z  INFO vector::app: API is disabled, enable by setting `api.enabled` to `true` and use commands like `vector top`.
2025-05-05T15:47:20.615221Z  INFO vector::topology::builder: Healthcheck passed.
2025-05-05T15:47:20.615271Z  INFO source{component_kind="source" component_id=line_continued_logs component_type=file}: vector::sources::file: Starting file server. include=["/etc/vector/*.log"] exclude=[]
2025-05-05T15:47:20.617682Z  INFO source{component_kind="source" component_id=line_continued_logs component_type=file}:file_server: file_source::checkpointer: Attempting to read legacy checkpoint files.
2025-05-05T15:47:20.627812Z  INFO source{component_kind="source" component_id=line_continued_logs component_type=file}:file_server: vector::internal_events::file::source: Found new file to watch. file=/etc/vector/LineContinuation.log
{"file":"/etc/vector/LineContinuation.log","host":"2884d7bfe5a3","message":"First line\\\nsecond line\\","source_type":"file","timestamp":"2025-05-05T15:47:21.636328836Z"}
```