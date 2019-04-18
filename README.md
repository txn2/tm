![tm](https://raw.githubusercontent.com/txn2/tm/master/mast.jpg)

WIP: TXN2 types model API


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