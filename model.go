// Copyright 2019 txn2
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tm

import (
	"strings"

	"github.com/txn2/es/v2"
)

const IdxModel = "models"

// Model
type Model struct {
	// MachineName is a lowercase under score delimited uniq id
	MachineName string `json:"machine_name" mapstructure:"machine_name"`

	// short human readable display name
	DisplayName string `json:"display_name" mapstructure:"display_name"`

	// a single sentence description
	BriefDescription string `json:"description_brief" mapstructure:"description_brief"`

	// full documentation in markdown
	Description string `json:"description" mapstructure:"description"`

	// default value expressed as a string
	DefaultValue string `json:"default_value" mapstructure:"default_value"`

	// integer, float, date, binary, text and keyword
	DataType string `json:"data_type" mapstructure:"data_type"`

	// used for data formats
	Format string `json:"format" mapstructure:"format"`

	// named parsers
	Parsers []string `json:"parsers" mapstructure:"parsers"`

	// belongs to a class of models
	TypeClass string `json:"type_class" mapstructure:"type_class"`

	// groups models
	Group string `json:"group" mapstructure:"group"`

	// false to ignore inbound parsing
	Parse bool `json:"parse" mapstructure:"parse"`

	// used by parsers of element ordered inbound data
	Index int `json:"index" mapstructure:"index"`

	// children of this model
	Fields []Model `json:"fields" mapstructure:"fields"`
}

// fieldProps
func fieldProps(fields []Model) map[string]interface{} {
	parts := make(map[string]interface{})

	for _, field := range fields {
		if len(field.Fields) > 0 {
			parts[field.MachineName] = es.Obj{
				"type":       field.DataType,
				"properties": fieldProps(field.Fields),
			}
			continue
		}

		if field.Format != "" {
			parts[field.MachineName] = es.Obj{
				"type":   field.DataType,
				"format": field.Format,
			}
			continue
		}

		parts[field.MachineName] = es.Obj{"type": field.DataType}
	}

	return parts
}

// MakeModelTemplateMapping creates a template for modeled data
// coming in from rxtx.
func MakeModelTemplateMapping(account string, model *Model) es.IndexTemplate {

	name := account + "-data-" + model.MachineName
	idxPattern := account + "-data-" + model.MachineName + "-*"

	// CONVENTION: if the account ends in an underscore "_" then
	// it is a system model (SYSTEM_IdxModel)
	if strings.HasSuffix(account, "_") {
		name = account + IdxModel + "-data-" + model.MachineName
		idxPattern = "*-data-" + model.MachineName + "-*"
	}

	payloadProps := fieldProps(model.Fields)

	template := es.Obj{
		"index_patterns": []string{idxPattern},
		"settings": es.Obj{
			"index": es.Obj{
				"number_of_shards": 1, // @TODO allow this to be configured
			},
		},
		"mappings": es.Obj{
			// _doc is the standard until deprecated, logstash uses "doc"
			// messages come into elasticsearch via txn2/rxtx->txn2/rtBeat->logstash
			"doc": es.Obj{
				"_source": es.Obj{
					"enabled": true,
				},
				"properties": es.Obj{
					// txn2/rtbeat sends txn2/rxtx messages as rxtxMsg
					"rxtxMsg": es.Obj{
						"properties": es.Obj{
							"seq":      es.Obj{"type": "long"},
							"producer": es.Obj{"type": "text"},
							"key":      es.Obj{"type": "text"},
							"uuid":     es.Obj{"type": "text"},
							"label":    es.Obj{"type": "text"},
							"payload": es.Obj{
								"properties": payloadProps,
							},
						},
					},
				},
			},
		},
	}

	return es.IndexTemplate{
		Name:     name,
		Template: template,
	}
}

// GetModelsTemplateMapping
func GetModelsTemplateMapping() es.IndexTemplate {
	properties := es.Obj{
		"machine_name": es.Obj{
			"type": "text",
		},
		"display_name": es.Obj{
			"type": "text",
		},
		"description_brief": es.Obj{
			"type": "text",
		},
		"description": es.Obj{
			"type": "text",
		},
		"default_value": es.Obj{
			"type": "keyword",
		},
		"data_type": es.Obj{
			"type": "keyword",
		},
		"format": es.Obj{
			"type": "text",
		},
		"parsers": es.Obj{
			"type": "text",
		},
		"type_class": es.Obj{
			"type": "keyword",
		},
		"group": es.Obj{
			"type": "keyword",
		},
		"parse": es.Obj{
			"type": "boolean",
		},
		"index": es.Obj{
			"type": "integer",
		},
		"fields": es.Obj{
			"type": "nested",
		},
	}

	template := es.Obj{
		"index_patterns": []string{"*-" + IdxModel, "*_" + IdxModel},
		"settings": es.Obj{
			"index": es.Obj{
				"number_of_shards": 1,
			},
		},
		"mappings": es.Obj{
			"_doc": es.Obj{
				"_source": es.Obj{
					"enabled": true,
				},
				"properties": properties,
			},
		},
	}

	return es.IndexTemplate{
		Name:     IdxModel,
		Template: template,
	}
}
