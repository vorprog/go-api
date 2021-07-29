CREATE EXTERNAL TABLE IF NOT EXISTS default.go_server (
  `container_id` string,
  `time` timestamp,
  `container_name` string,
  `log` string,
  `source` string
) ROW FORMAT SERDE 'org.openx.data.jsonserde.JsonSerDe' WITH SERDEPROPERTIES ('serialization.format' = '1') LOCATION 's3://vorprog-logs/go-server/' TBLPROPERTIES ('has_encrypted_data' = 'false')
