# Minecraft RCON Client REST API

[![Build Status](https://ci.olliejonas.com/job/mc-rest-rcon/job/master/badge/icon)](https://ci.olliejonas.com/job/mc-rest-rcon/job/master/)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Simple REST API that allows for the dispatching of commands remotely using the Minecraft RCON protocol. 

## Routes

### /health

GET status of service.

#### Example Response

```json
{
  "status": "OK"
}
```

### /dispatch

POST commands to dispatch to server.

#### Example Request

```json
{
  "address": "127.0.0.1",
  "rcon_port": "25575",
  "password": "mypassword",
  "commands": [
    "say Hello world!",
    "say I am being executed remotely!"
  ]
}
```

#### Example Response
```json
{
  "response": [

  ]
}
```

## License

[GNU General Public License v3.0](https://github.com/OllieJonas/mc-rcon-rest/blob/master/LICENSE)


