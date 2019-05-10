![tm](https://raw.githubusercontent.com/txn2/tm/master/mast.jpg)
[![tm Release](https://img.shields.io/github/release/txn2/tm.svg)](https://github.com/txn2/tm/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/txn2/tm)](https://goreportcard.com/report/github.com/txn2/tm)
[![GoDoc](https://godoc.org/github.com/txn2/tm?status.svg)](https://godoc.org/github.com/txn2/tm)
[![Docker Container Image Size](https://shields.beevelop.com/docker/image/image-size/txn2/tm/latest.svg)](https://hub.docker.com/r/txn2/tm/)
[![Docker Container Layers](https://shields.beevelop.com/docker/image/layers/txn2/tm/latest.svg)](https://hub.docker.com/r/txn2/tm/)


TXN2 types model API

## Configuration

Configuration is inherited from [txn2/micro](https://github.com/txn2/micro#configuration). The
following configuration is specific to **tm**:

| Flag          | Environment Variable | Description                                                |
|:--------------|:---------------------|:-----------------------------------------------------------|
| -esServer     | ELASTIC_SERVER       | Elasticsearch Server (default "http://elasticsearch:9200") |


## Examples

The following creates model called **test** and will result in a record with the id **test** in the **xorg-models** index. A mapping template will be also be generated and stored at **_template/xorg-data-test**:
```bash
curl -X POST \
  http://localhost:8080/model/xorg \
  -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{
    "machine_name": "test",
    "display_name": "",
    "description_brief": "",
    "description": "",
    "data_type": "",
    "format": "",
    "parsers": null,
    "type_class": "",
    "group": "",
    "parse": false,
    "index": 0,
    "fields": [
    	{
		    "machine_name": "event_type",
		    "display_name": "Event Type",
		    "description_brief": "",
		    "description": "",
		    "data_type": "keyword",
		    "format": "",
		    "parsers": null,
		    "type_class": "",
		    "group": "",
		    "parse": false,
		    "index": 0
		},
    	{
		    "machine_name": "gps_utc_time",
		    "display_name": "GPS UTC Time",
		    "description_brief": "",
		    "description": "",
		    "data_type": "date",
		    "format": "yyyyMMddHHmmss",
		    "parsers": null,
		    "type_class": "",
		    "group": "",
		    "parse": false,
		    "index": 0
		},
		{
		    "machine_name": "location",
		    "display_name": "",
		    "description_brief": "",
		    "description": "",
		    "data_type": "nested",
		    "format": "",
		    "parsers": null,
		    "type_class": "",
		    "group": "",
		    "parse": false,
		    "index": 0,
		    "fields": [
    	    	{
				    "machine_name": "lat",
				    "display_name": "",
				    "description_brief": "",
				    "description": "",
				    "data_type": "float",
				    "format": "",
				    "parsers": null,
				    "type_class": "",
				    "group": "",
				    "parse": false,
				    "index": 0
				},
    	    	{
				    "machine_name": "lon",
				    "display_name": "",
				    "description_brief": "",
				    "description": "",
				    "data_type": "float",
				    "format": "",
				    "parsers": null,
				    "type_class": "",
				    "group": "",
				    "parse": false,
				    "index": 0
				},				
		    ]
		}
	]
}'
```


## Release Packaging

Build test release:
```bash
goreleaser --skip-publish --rm-dist --skip-validate
```

Build and release:
```bash
GITHUB_TOKEN=$GITHUB_TOKEN goreleaser --rm-dist
```