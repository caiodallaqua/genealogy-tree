// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ascendants": {
            "get": {
                "description": "get all ascendants by person ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ascendants / Descendants"
                ],
                "summary": "Get Ascendants",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetAscendantsRes"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    }
                }
            }
        },
        "/ascendants-and-children": {
            "get": {
                "description": "give person ID. get all ascendants and children of the leaf ascendants",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ascendants / Descendants"
                ],
                "summary": "Get Ascendants and Immediate Children",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetAscendantsAndChildrenRes"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    }
                }
            }
        },
        "/ascendants-and-descendants": {
            "get": {
                "description": "get all ascendants and descendants by person ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ascendants / Descendants"
                ],
                "summary": "Get Ascendants and Descendants",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetAscendantsAndDescendantsRes"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    }
                }
            }
        },
        "/parent-relationship": {
            "post": {
                "description": "add relationship given two ids (of parent and child)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Relation"
                ],
                "summary": "Post Parent Relationship",
                "parameters": [
                    {
                        "description": "Parent relationship data",
                        "name": "ParentRelationship",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PostParentRelationship"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.PostParentRelationshipRes"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    }
                }
            },
            "delete": {
                "description": "removes relationship between parent and child given their ids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Relation"
                ],
                "summary": "Delete Parent Relationship",
                "parameters": [
                    {
                        "description": "Parent relationship",
                        "name": "ParentRelationship",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DelParentRelationship"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.DelParentRelationshipRes"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    }
                }
            }
        },
        "/person": {
            "get": {
                "description": "get person by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person"
                ],
                "summary": "Get Person",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Person ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetPersonRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    }
                }
            },
            "post": {
                "description": "add person given name and optionally birth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person"
                ],
                "summary": "Post Person",
                "parameters": [
                    {
                        "description": "Person data",
                        "name": "PostPerson",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PostPerson"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.PostPersonRes"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    }
                }
            },
            "delete": {
                "description": "removes person and its relationships given the person id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Person"
                ],
                "summary": "Delete Person",
                "parameters": [
                    {
                        "description": "Removes person from database",
                        "name": "DelPerson",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DelPerson"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.DelPersonRes"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest.GenTreeError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.DelParentRelationship": {
            "type": "object",
            "required": [
                "child_id",
                "parent_id"
            ],
            "properties": {
                "child_id": {
                    "type": "integer"
                },
                "parent_id": {
                    "type": "integer"
                }
            }
        },
        "models.DelParentRelationshipRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "string": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "models.DelPerson": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.DelPersonRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "string": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "models.GetAscendantsAndChildrenRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "ascendantsAndChildren": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PersonData"
                            }
                        },
                        "relationships": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.RelationData"
                            }
                        }
                    }
                }
            }
        },
        "models.GetAscendantsAndDescendantsRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "ascendantsAndDescendants": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PersonData"
                            }
                        },
                        "relationships": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.RelationData"
                            }
                        }
                    }
                }
            }
        },
        "models.GetAscendantsRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "ascendants": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PersonData"
                            }
                        },
                        "relationships": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.RelationData"
                            }
                        }
                    }
                }
            }
        },
        "models.GetPersonRes": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.PersonData"
                }
            }
        },
        "models.PersonData": {
            "type": "object",
            "properties": {
                "birth": {
                    "type": "string",
                    "example": "2020-12-09T16:09:53+00:00"
                },
                "id": {
                    "type": "integer",
                    "example": 42
                },
                "name": {
                    "type": "string",
                    "example": "Aristeu"
                }
            }
        },
        "models.PostParentRelationship": {
            "type": "object",
            "required": [
                "child_id",
                "parent_id"
            ],
            "properties": {
                "child_id": {
                    "type": "integer"
                },
                "parent_id": {
                    "type": "integer"
                }
            }
        },
        "models.PostParentRelationshipRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "string": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "models.PostPerson": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "birth": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.PostPersonRes": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "id": {
                            "type": "integer",
                            "example": 42
                        }
                    }
                }
            }
        },
        "models.RelationData": {
            "type": "object",
            "properties": {
                "child_id": {
                    "type": "integer",
                    "example": 4
                },
                "parent_id": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "rest.GenTreeError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Failed to deserialize. Check request."
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Genealogy Tree API",
	Description:      "API for handling person entity and its relationships in a genealogy tree.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
