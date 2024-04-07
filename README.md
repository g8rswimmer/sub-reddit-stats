# sub-reddit-stats
This application is a couple of services to retrieve subnreddit data and return it to the user.

## Services
The applicaiton uses multiple services to retrieve the data and to serve it a user.  The use of multiple services allows for seperation of responsiblities and for some scaling if desired.

### Setup
Before running the application, the database file (used for `sqlite`) needs to be created.  Running the make command, `init_db`, will create the director and db file that will be used.  One can clean up the db file but running the make command `clean_db`.

After initializing the database, running the make command `migrate` will set up the database.

#### Configuration
There is a configuration file that is used for migration, example `config-tmpl.json`.  

Those settings are under `database` object:

| Name          | Type     | Required | Description                                       |
|---------------|----------|----------|---------------------------------------------------|
| `database.data_source` | `string` | Y        | The data source used with opening up the database |

##### Example
```json
{
    "database": {
        "data_source": "./db/sqlite-database.db"
    }
}
```

### Deamon
The daemon is used to retrieve subreddit data from reddit's APIs.  The data will be stored in a datastore for the service to query when needed.

You can run the daemon by using the make command `daemon`.

#### Configuration
There is a configuration file that is used for daemon, example `config-tmpl.json`.

Those settings are under `database` and `reddit` objects:

| Name                   | Type     | Required | Description                                                           |
|------------------------|----------|----------|-----------------------------------------------------------------------|
| `database.data_source` | `string` | Y        | The data source used with opening up the database                     |
| `reddit.client_id`     | `string` | Y        | The reddit client ID used to obtain the access token for the APIs     |
| `reddit.client_secret` | `string` | Y        | The reddit client secret used to obtain the access token for the APIs |
| `reddit.oauth_url`     | `string` | N        | The reddit oauth URL used for obtaining an access token               |
| `reddit.base_url`      | `string` | Y        | The reddit base URL used for access to the API                        |
| `reddit.subreddit`     | `string` | Y        | The subreddit name for the daemon.  This should not contain the `/r`  |

#### Example
```json
{
    "database": {
        "data_source": "./db/sqlite-database.db"
    },
    "reddit": {
        "client_id": "your client ID",
        "client_secret": "your client secret",
        "oauth_url": "https://www.reddit.com",
        "base_url": "https://oauth.reddit.com",
        "subreddit": "funny"
    }
}
```

### Server
The server is used to retrieve stats based on the subreddit listings.  This server supports both gRPC and HTTP requests.

You can run the server by using the make command `server`.

#### Documenation
The swagger documentation for the service can be found at:
```
http://[host]/docs
```

#### Configuration
There is a configuration file that is used for server, example `config-tmpl.json`.

Those settings are under `database` and `server` objects:

| Name                   | Type     | Required | Description                                       |
|------------------------|----------|----------|---------------------------------------------------|
| `database.data_source` | `string` | Y        | The data source used with opening up the database |
| `server.grpc_port`     | `int`    | Y        | The gRPC server port for request to be handled    |
| `server.http_port`     | `int`    | Y        | The HTTP server port for request to be handled    |

#### Example

```json
{
    "database": {
        "data_source": "./db/sqlite-database.db"
    },
    "server": {
        "grpc_port": 5050,
        "http_port": 8080
    }
}
```