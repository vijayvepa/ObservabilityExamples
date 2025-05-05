```shell

vector
```

Example output:

```shell
2025-05-05T15:54:39.438169Z  INFO vector::app: Log level is enabled. level="info"
2025-05-05T15:54:39.438699Z  INFO vector::app: Loading configs. paths=["/etc/vector/vector.yaml"]
2025-05-05T15:54:39.449432Z  INFO vector::topology::running: Running healthchecks.
2025-05-05T15:54:39.449510Z  INFO vector: Vector has started. debug="false" version="0.46.1" arch="aarch64" revision="9a19e8a 2025-04-14 18:36:30.707862743"
2025-05-05T15:54:39.449517Z  INFO vector::app: API is disabled, enable by setting `api.enabled` to `true` and use commands like `vector top`.
2025-05-05T15:54:39.449656Z  INFO vector::topology::builder: Healthcheck passed.
2025-05-05T15:54:39.449709Z  INFO source{component_kind="source" component_id=elasticsearch_logs component_type=file}: vector::sources::file: Starting file server. include=["/etc/vector/*.log"] exclude=[]
2025-05-05T15:54:39.452965Z  INFO source{component_kind="source" component_id=elasticsearch_logs component_type=file}:file_server: file_source::checkpointer: Attempting to read legacy checkpoint files.
2025-05-05T15:54:39.458118Z  INFO source{component_kind="source" component_id=elasticsearch_logs component_type=file}:file_server: vector::internal_events::file::source: Found new file to watch. file=/etc/vector/Elasticsearch.log
{"file":"/etc/vector/Elasticsearch.log","host":"0f834cb1201b","message":"[2015-08-24 11:49:14,389][ INFO ][env                      ] [Letha] using [1] data paths, mounts [[/\n(/dev/disk1)]], net usable_space [34.5gb], net total_space [118.9gb], types [hfs]\n","source_type":"file","timestamp":"2025-05-05T15:54:39.459274011Z"}
{"file":"/etc/vector/Elasticsearch.log","host":"0f834cb1201b","message":"[2015-08-24 11:49:14,389][ INFO ][env                      ] [Letha] using [1] data paths, mounts [[/\n(/dev/disk1)]], net usable_space [34.5gb], net total_space [118.9gb], types [hfs]\n","source_type":"file","timestamp":"2025-05-05T15:54:40.462506178Z"}

```