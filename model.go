package tm

import "github.com/txn2/es"

const IdxModel = "models"

// Model
type Model struct {
	// a lowercase under score delimited uniq id
	MachineName string `json:"machine_name" mapstructure:"machine_name"`

	// short human readable display name
	DisplayName string `json:"display_name" mapstructure:"display_name"`

	// a single sentence description
	BriefDescription string `json:"description_brief" mapstructure:"description_brief"`

	// full documentation in markdown
	Description string `json:"description" mapstructure:"description"`

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

// GetModelMapping
func GetModelMapping() es.IndexTemplate {
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
		"index_patterns": []string{"*-" + IdxModel},
		"settings": es.Obj{
			"number_of_shards": 2,
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
