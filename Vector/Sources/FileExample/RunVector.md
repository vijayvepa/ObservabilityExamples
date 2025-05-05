```shell

vector
```

Example Output:

```shell

2025-05-05T15:31:50.928111Z  INFO vector::app: Log level is enabled. level="info"
2025-05-05T15:31:50.929223Z  INFO vector::app: Loading configs. paths=["/etc/vector/vector.yaml"]
2025-05-05T15:31:50.951039Z  INFO vector::topology::running: Running healthchecks.
2025-05-05T15:31:50.951117Z  INFO vector: Vector has started. debug="false" version="0.46.1" arch="aarch64" revision="9a19e8a 2025-04-14 18:36:30.707862743"
2025-05-05T15:31:50.951126Z  INFO vector::app: API is disabled, enable by setting `api.enabled` to `true` and use commands like `vector top`.
2025-05-05T15:31:50.951182Z  INFO source{component_kind="source" component_id=apache_access_log component_type=file}: vector::sources::file: Starting file server. include=["/etc/vector/*.log"] exclude=[]
2025-05-05T15:31:50.951404Z  INFO vector::topology::builder: Healthcheck passed.
2025-05-05T15:31:50.952120Z  INFO source{component_kind="source" component_id=apache_access_log component_type=file}:file_server: file_source::checkpointer: Attempting to read legacy checkpoint files.
2025-05-05T15:31:50.957716Z  INFO source{component_kind="source" component_id=apache_access_log component_type=file}:file_server: vector::internal_events::file::source: Found new file to watch. file=/etc/vector/ApacheAccessLog.log
{"file":"/etc/vector/ApacheAccessLog.log","host":"3f59f7385fc0","message":"53.126.150.246 - - [01/Oct/2020:11:25:58 -0400] \"GET /disintermediate HTTP/2.0\" 401 20308","source_type":"file","timestamp":"2025-05-05T15:31:50.958956003Z"}
```

