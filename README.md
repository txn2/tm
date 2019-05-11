![tm](https://raw.githubusercontent.com/txn2/tm/master/mast.jpg)
[![tm Release](https://img.shields.io/github/release/txn2/tm.svg)](https://github.com/txn2/tm/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/txn2/tm)](https://goreportcard.com/report/github.com/txn2/tm)
[![GoDoc](https://godoc.org/github.com/txn2/tm?status.svg)](https://godoc.org/github.com/txn2/tm)
[![Docker Container Image Size](https://shields.beevelop.com/docker/image/image-size/txn2/tm/latest.svg)](https://hub.docker.com/r/txn2/tm/)
[![Docker Container Layers](https://shields.beevelop.com/docker/image/layers/txn2/tm/latest.svg)](https://hub.docker.com/r/txn2/tm/)

**tm** is used for defining the strucutre and Elasticsearch indexing rules for [Message]s
sent to Elasticsearch from [rxtx] through [rtBeat](https://github.com/txn2/rtbeat) with
the key `rxtxMsg`. **tm** [Model]s define the properties of the [rxtx] `payload`.

The **tm** library defines a type of generic nested meta-data [Model]. The **tm** server creates a services for the storage,
retrieval and searching of [Model]s associated with a [txn2/provision](https://github.com/txn2/provision) [Account].

A [Model] consists of a record stored in the Elasticsearch index **ACCOUNT-models** and a corresponding Elasticsearch
template (**_template/ACCOUNT-data-MODEL**) representing the index pattern **ACCOUNT-data-MODEL-\***.


[Account]: https://godoc.org/github.com/txn2/provision#Account
[Message]: https://godoc.org/github.com/txn2/rxtx/rtq#Message
[rxtx]: https://github.com/txn2/rxtx

## Configuration

Configuration is inherited from [txn2/micro](https://github.com/txn2/micro#configuration). The
following configuration is specific to **tm**:

| Flag      | Environment Variable | Description                                                    |
|:----------|:---------------------|:---------------------------------------------------------------|
| -esServer | ELASTIC_SERVER       | Elasticsearch Server (default "http://elasticsearch:9200")     |
| -mode     | MODE                 | Protected or internal modes. ("internal" = token check bypass) |

## Routes

| Method | Route Pattern                           | Description                                          |
|:-------|:----------------------------------------|:-----------------------------------------------------|
| POST   | [/model/:account](#upsert-model)        | Upset a model into an account.                       |
| GET    | [/model/:account/:id](#get-model)       | Get a model by account and id.                       |
| POST   | [searchModels/:account](#search-models) | Search for models in an account with a Lucene query. |

## Local Development

The project includes a Docker Compose file with Elasticsearch, Kibana and Cerebro:
```bash
docker-compose up
```

Run the source in token bypass mode and pointed to Elasticsearch exposed on localhost port 9200:
```bash
go run ./cmd/tm.go --mode=internal --esServer=http://localhost:9200
```

## Examples

The following examples assume mode is set to internal and will not check a Bearer token for
proper permissions.

#### Upsert Model

Upserting a [Model] will result in an [Ack] with a [Result] payload.

The following creates a model called **test** and will result in a record with the id **test**
in the **xorg-models** index. A mapping template will be also be generated and stored
at **_template/xorg-data-test**:
```bash
curl -X POST \
  http://localhost:8080/model/xorg \
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
				}				
		    ]
		}
	]
}'
```

### Get Model

Getting a [Model] will result in a [ModelResultAck].

```bash
curl http://localhost:8080/model/xorg/test
```

### Search Models

Searching for [Model]s will result in a [ModelSearchResultsAck].

```bash
curl -X POST \
  http://localhost:8080/searchModels/xorg \
  -d '{
  "query": {
    "match_all": {}
  }
}'
```

[Ack]: https://godoc.org/github.com/txn2/ack#Ack
[Result]: https://godoc.org/github.com/txn2/es#Result
[Model]: https://godoc.org/github.com/txn2/tm#Model
[ModelSearchResultsAck]: https://godoc.org/github.com/txn2/tm#ModelSearchResultsAck
[ModelResultAck]: https://godoc.org/github.com/txn2/tm#ModelResultAck

## Release Packaging

Build test release:
```bash
goreleaser --skip-publish --rm-dist --skip-validate
```

Build and release:
```bash
GITHUB_TOKEN=$GITHUB_TOKEN goreleaser --rm-dist
```