## REST API Boilerplate + gRPC
[![CI](https://github.com/Lynicis/go-rest-api-boilerplate/actions/workflows/master-ci.yml/badge.svg?branch=master&event=push)](https://github.com/Lynicis/go-rest-api-boilerplate/actions/workflows/master-ci.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Lynicis_go-rest-api-boilerplate&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Lynicis_go-rest-api-boilerplate)

[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=Lynicis_go-rest-api-boilerplate&metric=coverage)](https://sonarcloud.io/summary/new_code?id=Lynicis_go-rest-api-boilerplate)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=Lynicis_go-rest-api-boilerplate&metric=bugs)](https://sonarcloud.io/summary/new_code?id=Lynicis_go-rest-api-boilerplate)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=Lynicis_go-rest-api-boilerplate&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=Lynicis_go-rest-api-boilerplate)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=Lynicis_go-rest-api-boilerplate&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=Lynicis_go-rest-api-boilerplate)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=Lynicis_go-rest-api-boilerplate&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=Lynicis_go-rest-api-boilerplate)

#### Directory Structure

```
repo-name/
├── .github //pipeline configs for github
│    ├── actions
│    └── workflows
├── .k8s //kubernetes configs
├── cmd //main program will be here
├── config
├── internal //endpoints
│   └── health //health check endpoint
└── pkg
    ├── config
    ├── http_server
    ├── logger
    ├── pact
    ├── project_path
    └── rpc_server
```