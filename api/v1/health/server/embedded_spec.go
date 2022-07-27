// Code generated by go-swagger; DO NOT EDIT.

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
          "$ref": "../openapi.yaml#/definitions/StatusResponse"
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
          "$ref": "#/definitions/statusResponse"
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
    "error": {
      "type": "string"
    },
    "statusResponse": {
      "description": "Health and status information of daemon",
      "type": "object",
      "properties": {
        "cilium": {
          "description": "Status of an individual component",
          "type": "object",
          "properties": {
            "msg": {
              "description": "Human readable status/error/warning message",
              "type": "string"
            },
            "state": {
              "description": "State the component is in",
              "type": "string",
              "enum": [
                "Ok",
                "Warning",
                "Failure",
                "Disabled"
              ]
            }
          }
        },
        "cluster": {
          "description": "Status of cluster",
          "properties": {
            "ciliumHealth": {
              "description": "Status of an individual component",
              "type": "object",
              "properties": {
                "msg": {
                  "description": "Human readable status/error/warning message",
                  "type": "string"
                },
                "state": {
                  "description": "State the component is in",
                  "type": "string",
                  "enum": [
                    "Ok",
                    "Warning",
                    "Failure",
                    "Disabled"
                  ]
                }
              }
            },
            "nodes": {
              "description": "List of known nodes",
              "type": "array",
              "items": {
                "description": "Known node in the cluster",
                "properties": {
                  "health-endpoint-address": {
                    "description": "Addressing information of a node for all address families",
                    "type": "object",
                    "properties": {
                      "ipv4": {
                        "description": "Addressing information",
                        "type": "object",
                        "properties": {
                          "address-type": {
                            "description": "Node address type, one of HostName, ExternalIP or InternalIP",
                            "type": "string"
                          },
                          "alloc-range": {
                            "description": "Address pool to be used for local endpoints",
                            "type": "string"
                          },
                          "enabled": {
                            "description": "True if address family is enabled",
                            "type": "boolean"
                          },
                          "ip": {
                            "description": "IP address of node",
                            "type": "string"
                          }
                        }
                      },
                      "ipv6": {
                        "description": "Addressing information",
                        "type": "object",
                        "properties": {
                          "address-type": {
                            "description": "Node address type, one of HostName, ExternalIP or InternalIP",
                            "type": "string"
                          },
                          "alloc-range": {
                            "description": "Address pool to be used for local endpoints",
                            "type": "string"
                          },
                          "enabled": {
                            "description": "True if address family is enabled",
                            "type": "boolean"
                          },
                          "ip": {
                            "description": "IP address of node",
                            "type": "string"
                          }
                        }
                      }
                    }
                  },
                  "name": {
                    "description": "Name of the node including the cluster association. This is typically\n\u003cclustername\u003e/\u003chostname\u003e.\n",
                    "type": "string"
                  },
                  "primary-address": {
                    "description": "Addressing information of a node for all address families",
                    "type": "object",
                    "properties": {
                      "ipv4": {
                        "description": "Addressing information",
                        "type": "object",
                        "properties": {
                          "address-type": {
                            "description": "Node address type, one of HostName, ExternalIP or InternalIP",
                            "type": "string"
                          },
                          "alloc-range": {
                            "description": "Address pool to be used for local endpoints",
                            "type": "string"
                          },
                          "enabled": {
                            "description": "True if address family is enabled",
                            "type": "boolean"
                          },
                          "ip": {
                            "description": "IP address of node",
                            "type": "string"
                          }
                        }
                      },
                      "ipv6": {
                        "description": "Addressing information",
                        "type": "object",
                        "properties": {
                          "address-type": {
                            "description": "Node address type, one of HostName, ExternalIP or InternalIP",
                            "type": "string"
                          },
                          "alloc-range": {
                            "description": "Address pool to be used for local endpoints",
                            "type": "string"
                          },
                          "enabled": {
                            "description": "True if address family is enabled",
                            "type": "boolean"
                          },
                          "ip": {
                            "description": "IP address of node",
                            "type": "string"
                          }
                        }
                      }
                    }
                  },
                  "secondary-addresses": {
                    "description": "Alternative addresses assigned to the node",
                    "type": "array",
                    "items": {
                      "description": "Addressing information",
                      "type": "object",
                      "properties": {
                        "address-type": {
                          "description": "Node address type, one of HostName, ExternalIP or InternalIP",
                          "type": "string"
                        },
                        "alloc-range": {
                          "description": "Address pool to be used for local endpoints",
                          "type": "string"
                        },
                        "enabled": {
                          "description": "True if address family is enabled",
                          "type": "boolean"
                        },
                        "ip": {
                          "description": "IP address of node",
                          "type": "string"
                        }
                      }
                    }
                  }
                }
              }
            },
            "self": {
              "description": "Name of local node (if available)",
              "type": "string"
            }
          }
        },
        "container-runtime": {
          "description": "Status of an individual component",
          "type": "object",
          "properties": {
            "msg": {
              "description": "Human readable status/error/warning message",
              "type": "string"
            },
            "state": {
              "description": "State the component is in",
              "type": "string",
              "enum": [
                "Ok",
                "Warning",
                "Failure",
                "Disabled"
              ]
            }
          }
        },
        "controllers": {
          "description": "Collection of controller statuses",
          "type": "array",
          "items": {
            "description": "Status of a controller",
            "type": "object",
            "properties": {
              "configuration": {
                "description": "Configuration of controller",
                "type": "object",
                "properties": {
                  "error-retry": {
                    "description": "Retry on error",
                    "type": "boolean"
                  },
                  "error-retry-base": {
                    "description": "Base error retry back-off time",
                    "type": "string",
                    "format": "duration"
                  },
                  "interval": {
                    "description": "Regular synchronization interval",
                    "type": "string",
                    "format": "duration"
                  }
                }
              },
              "name": {
                "description": "Name of controller",
                "type": "string"
              },
              "status": {
                "description": "Current status of controller",
                "type": "object",
                "properties": {
                  "consecutive-failure-count": {
                    "description": "Number of consecutive errors since last success",
                    "type": "integer"
                  },
                  "failure-count": {
                    "description": "Total number of failed runs",
                    "type": "integer"
                  },
                  "last-failure-msg": {
                    "description": "Error message of last failed run",
                    "type": "string"
                  },
                  "last-failure-timestamp": {
                    "description": "Timestamp of last error",
                    "type": "string",
                    "format": "date-time"
                  },
                  "last-success-timestamp": {
                    "description": "Timestamp of last success",
                    "type": "string",
                    "format": "date-time"
                  },
                  "success-count": {
                    "description": "Total number of successful runs",
                    "type": "integer"
                  }
                }
              },
              "uuid": {
                "description": "UUID of controller",
                "type": "string",
                "format": "uuid"
              }
            }
          }
        },
        "ipam": {
          "description": "Status of IP address management",
          "properties": {
            "allocations": {
              "description": "Map of allocated IPs\n",
              "type": "object",
              "additionalProperties": {
                "type": "string"
              }
            },
            "ipv4": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "ipv6": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "status": {
              "type": "string"
            }
          }
        },
        "kubernetes": {
          "description": "Status of Kubernetes integration",
          "type": "object",
          "properties": {
            "k8s-api-versions": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "msg": {
              "description": "Human readable status/error/warning message",
              "type": "string"
            },
            "state": {
              "description": "State the component is in",
              "type": "string",
              "enum": [
                "Ok",
                "Warning",
                "Failure",
                "Disabled"
              ]
            }
          }
        },
        "kvstore": {
          "description": "Status of an individual component",
          "type": "object",
          "properties": {
            "msg": {
              "description": "Human readable status/error/warning message",
              "type": "string"
            },
            "state": {
              "description": "State the component is in",
              "type": "string",
              "enum": [
                "Ok",
                "Warning",
                "Failure",
                "Disabled"
              ]
            }
          }
        },
        "nodeMonitor": {
          "description": "Status of the node monitor",
          "properties": {
            "cpus": {
              "description": "Number of CPUs to listen on for events.",
              "type": "integer"
            },
            "lost": {
              "description": "Number of samples lost by perf.",
              "type": "integer"
            },
            "npages": {
              "description": "Number of pages used for the perf ring buffer.",
              "type": "integer"
            },
            "pagesize": {
              "description": "Pages size used for the perf ring buffer.",
              "type": "integer"
            },
            "unknown": {
              "description": "Number of unknown samples.",
              "type": "integer"
            }
          }
        },
        "proxy": {
          "description": "Status of proxy",
          "type": "object",
          "properties": {
            "ip": {
              "description": "IP address that the proxy listens on",
              "type": "string"
            },
            "port-range": {
              "description": "Port range used for proxying",
              "type": "string"
            }
          }
        },
        "stale": {
          "description": "List of stale information in the status",
          "type": "object",
          "additionalProperties": {
            "description": "Timestamp when the probe was started",
            "type": "string",
            "format": "date-time"
          }
        }
      }
    }
  },
  "x-schemes": [
    "unix"
  ]
}`))
}
