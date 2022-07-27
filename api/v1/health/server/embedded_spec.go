// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Cilium Health Checker",
    "title": "Cilium-Health API",
    "version": "v1beta"
  },
  "basePath": "/v1beta",
  "paths": {
    "/healthz": {
      "get": {
        "description": "Returns health and status information of the local node including\nload and uptime, as well as the status of related components including\nthe Cilium daemon.\n",
        "summary": "Get health of Cilium node",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/HealthResponse"
            }
          },
          "500": {
            "description": "Failed to contact local Cilium daemon",
            "schema": {
              "$ref": "../openapi.yaml#/definitions/Error"
            },
            "x-go-name": "Failed"
          }
        }
      }
    },
    "/status": {
      "get": {
        "description": "Returns the connectivity status to all other cilium-health instances\nusing interval-based probing.\n",
        "tags": [
          "connectivity"
        ],
        "summary": "Get connectivity status of the Cilium cluster",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/HealthStatusResponse"
            }
          }
        }
      }
    },
    "/status/probe": {
      "put": {
        "description": "Runs a synchronous probe to all other cilium-health instances and\nreturns the connectivity status.\n",
        "tags": [
          "connectivity"
        ],
        "summary": "Run synchronous connectivity probe to determine status of the Cilium cluster",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/HealthStatusResponse"
            }
          },
          "500": {
            "description": "Internal error occurred while conducting connectivity probe",
            "schema": {
              "$ref": "../openapi.yaml#/definitions/Error"
            },
            "x-go-name": "Failed"
          }
        }
      }
    }
  },
  "definitions": {
    "ConnectivityStatus": {
      "description": "Connectivity status of a path",
      "type": "object",
      "properties": {
        "latency": {
          "description": "Round trip time to node in nanoseconds",
          "type": "integer"
        },
        "status": {
          "description": "Human readable status/error/warning message",
          "type": "string"
        }
      }
    },
    "HealthResponse": {
      "description": "Health and status information of local node",
      "type": "object",
      "properties": {
        "cilium": {
          "description": "Status of Cilium daemon",
          "$ref": "#/definitions/StatusResponse"
        },
        "system-load": {
          "description": "System load on node",
          "$ref": "#/definitions/LoadResponse"
        },
        "uptime": {
          "description": "Uptime of cilium-health instance",
          "type": "string"
        }
      }
    },
    "HealthStatusResponse": {
      "description": "Connectivity status to other daemons",
      "type": "object",
      "properties": {
        "local": {
          "description": "Description of the local node",
          "$ref": "#/definitions/SelfStatus"
        },
        "nodes": {
          "description": "Connectivity status to each other node",
          "type": "array",
          "items": {
            "$ref": "#/definitions/NodeStatus"
          }
        },
        "timestamp": {
          "type": "string"
        }
      }
    },
    "HostStatus": {
      "description": "Connectivity status to host cilium-health instance via different paths,\nprobing via all known IP addresses\n",
      "properties": {
        "primary-address": {
          "$ref": "#/definitions/PathStatus"
        },
        "secondary-addresses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/PathStatus"
          }
        }
      }
    },
    "LoadResponse": {
      "description": "System load on node",
      "type": "object",
      "properties": {
        "last15min": {
          "description": "Load average over the past 15 minutes",
          "type": "string"
        },
        "last1min": {
          "description": "Load average over the past minute",
          "type": "string"
        },
        "last5min": {
          "description": "Load average over the past 5 minutes",
          "type": "string"
        }
      }
    },
    "NodeStatus": {
      "description": "Connectivity status of a remote cilium-health instance",
      "type": "object",
      "properties": {
        "endpoint": {
          "description": "Connectivity status to simulated endpoint on node IP",
          "$ref": "#/definitions/PathStatus"
        },
        "host": {
          "description": "Connectivity status to cilium-health instance on node IP",
          "$ref": "#/definitions/HostStatus"
        },
        "name": {
          "description": "Identifying name for the node",
          "type": "string"
        }
      }
    },
    "PathStatus": {
      "description": "Connectivity status via different paths, for example using different\npolicies or service redirection\n",
      "type": "object",
      "properties": {
        "http": {
          "description": "Connectivity status without policy applied",
          "$ref": "#/definitions/ConnectivityStatus"
        },
        "icmp": {
          "description": "Basic ping connectivity status to node IP",
          "$ref": "#/definitions/ConnectivityStatus"
        },
        "ip": {
          "description": "IP address queried for the connectivity status",
          "type": "string"
        }
      }
    },
    "SelfStatus": {
      "description": "Description of the cilium-health node",
      "type": "object",
      "properties": {
        "name": {
          "description": "Name associated with this node",
          "type": "string"
        }
      }
    },
    "StatusResponse": {
      "description": "Status of Cilium daemon",
      "type": "object",
      "x-go-type": {
        "hint": {
          "kind": "object",
          "nullable": true
        },
        "import": {
          "alias": "ciliumModels",
          "package": "github.com/cilium/cilium/api/v1/models"
        },
        "type": "StatusResponse"
      }
    }
  },
  "x-schemes": [
    "unix"
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Cilium Health Checker",
    "title": "Cilium-Health API",
    "version": "v1beta"
  },
  "basePath": "/v1beta",
  "paths": {
    "/healthz": {
      "get": {
        "description": "Returns health and status information of the local node including\nload and uptime, as well as the status of related components including\nthe Cilium daemon.\n",
        "summary": "Get health of Cilium node",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/HealthResponse"
            }
          },
          "500": {
            "description": "Failed to contact local Cilium daemon",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "x-go-name": "Failed"
          }
        }
      }
    },
    "/status": {
      "get": {
        "description": "Returns the connectivity status to all other cilium-health instances\nusing interval-based probing.\n",
        "tags": [
          "connectivity"
        ],
        "summary": "Get connectivity status of the Cilium cluster",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/HealthStatusResponse"
            }
          }
        }
      }
    },
    "/status/probe": {
      "put": {
        "description": "Runs a synchronous probe to all other cilium-health instances and\nreturns the connectivity status.\n",
        "tags": [
          "connectivity"
        ],
        "summary": "Run synchronous connectivity probe to determine status of the Cilium cluster",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/HealthStatusResponse"
            }
          },
          "500": {
            "description": "Internal error occurred while conducting connectivity probe",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "x-go-name": "Failed"
          }
        }
      }
    }
  },
  "definitions": {
    "ConnectivityStatus": {
      "description": "Connectivity status of a path",
      "type": "object",
      "properties": {
        "latency": {
          "description": "Round trip time to node in nanoseconds",
          "type": "integer"
        },
        "status": {
          "description": "Human readable status/error/warning message",
          "type": "string"
        }
      }
    },
    "HealthResponse": {
      "description": "Health and status information of local node",
      "type": "object",
      "properties": {
        "cilium": {
          "description": "Status of Cilium daemon",
          "$ref": "#/definitions/StatusResponse"
        },
        "system-load": {
          "description": "System load on node",
          "$ref": "#/definitions/LoadResponse"
        },
        "uptime": {
          "description": "Uptime of cilium-health instance",
          "type": "string"
        }
      }
    },
    "HealthStatusResponse": {
      "description": "Connectivity status to other daemons",
      "type": "object",
      "properties": {
        "local": {
          "description": "Description of the local node",
          "$ref": "#/definitions/SelfStatus"
        },
        "nodes": {
          "description": "Connectivity status to each other node",
          "type": "array",
          "items": {
            "$ref": "#/definitions/NodeStatus"
          }
        },
        "timestamp": {
          "type": "string"
        }
      }
    },
    "HostStatus": {
      "description": "Connectivity status to host cilium-health instance via different paths,\nprobing via all known IP addresses\n",
      "properties": {
        "primary-address": {
          "$ref": "#/definitions/PathStatus"
        },
        "secondary-addresses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/PathStatus"
          }
        }
      }
    },
    "LoadResponse": {
      "description": "System load on node",
      "type": "object",
      "properties": {
        "last15min": {
          "description": "Load average over the past 15 minutes",
          "type": "string"
        },
        "last1min": {
          "description": "Load average over the past minute",
          "type": "string"
        },
        "last5min": {
          "description": "Load average over the past 5 minutes",
          "type": "string"
        }
      }
    },
    "NodeStatus": {
      "description": "Connectivity status of a remote cilium-health instance",
      "type": "object",
      "properties": {
        "endpoint": {
          "description": "Connectivity status to simulated endpoint on node IP",
          "$ref": "#/definitions/PathStatus"
        },
        "host": {
          "description": "Connectivity status to cilium-health instance on node IP",
          "$ref": "#/definitions/HostStatus"
        },
        "name": {
          "description": "Identifying name for the node",
          "type": "string"
        }
      }
    },
    "PathStatus": {
      "description": "Connectivity status via different paths, for example using different\npolicies or service redirection\n",
      "type": "object",
      "properties": {
        "http": {
          "description": "Connectivity status without policy applied",
          "$ref": "#/definitions/ConnectivityStatus"
        },
        "icmp": {
          "description": "Basic ping connectivity status to node IP",
          "$ref": "#/definitions/ConnectivityStatus"
        },
        "ip": {
          "description": "IP address queried for the connectivity status",
          "type": "string"
        }
      }
    },
    "SelfStatus": {
      "description": "Description of the cilium-health node",
      "type": "object",
      "properties": {
        "name": {
          "description": "Name associated with this node",
          "type": "string"
        }
      }
    },
    "StatusResponse": {
      "description": "Status of Cilium daemon",
      "type": "object",
      "x-go-type": {
        "hint": {
          "kind": "object",
          "nullable": true
        },
        "import": {
          "alias": "ciliumModels",
          "package": "github.com/cilium/cilium/api/v1/models"
        },
        "type": "StatusResponse"
      }
    },
    "error": {
      "type": "string"
    }
  },
  "x-schemes": [
    "unix"
  ]
}`))
}
