```shell

vector
```

Example output:
```shell

2025-05-05T15:43:12.418139Z  INFO vector::app: Log level is enabled. level="info"
2025-05-05T15:43:12.418656Z  INFO vector::app: Loading configs. paths=["/etc/vector/vector.yaml"]
2025-05-05T15:43:12.431519Z  INFO vector::topology::running: Running healthchecks.
2025-05-05T15:43:12.431597Z  INFO vector: Vector has started. debug="false" version="0.46.1" arch="aarch64" revision="9a19e8a 2025-04-14 18:36:30.707862743"
2025-05-05T15:43:12.431592Z  INFO vector::topology::builder: Healthcheck passed.
2025-05-05T15:43:12.431603Z  INFO vector::app: API is disabled, enable by setting `api.enabled` to `true` and use commands like `vector top`.
2025-05-05T15:43:12.431646Z  INFO source{component_kind="source" component_id=ruby_logs component_type=file}: vector::sources::file: Starting file server. include=["/etc/vector/*.log"] exclude=[]
2025-05-05T15:43:12.432419Z  INFO source{component_kind="source" component_id=ruby_logs component_type=file}:file_server: file_source::checkpointer: Attempting to read legacy checkpoint files.
2025-05-05T15:43:12.440070Z  INFO source{component_kind="source" component_id=ruby_logs component_type=file}:file_server: vector::internal_events::file::source: Found new file to watch. file=/etc/vector/RubyMultilineLog.log
{"file":"/etc/vector/RubyMultilineLog.log","host":"1712ba38984e","message":"foobar.rb:6:in `/': divided by 0 (ZeroDivisionError)\n\tfrom foobar.rb:6:in `bar'\n\tfrom foobar.rb:2:in `foo'","source_type":"file","timestamp":"2025-05-05T15:43:13.444156472Z"}
```