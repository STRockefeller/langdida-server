# Usage Documents for Config Definitions

## Configuration Options

### UseGinServer

- Type: boolean
- Description: This flag determines whether to use Gin server for delivery or not. If set to `true`, Gin server will be used; if set to `false`, other delivery methods will be used.
- Usage: Set this flag to `true` if you want to use Gin server for handling incoming requests, and `false` otherwise.

### UseSqlite

- Type: boolean
- Description: This flag determines whether to use SQLite for storage or not. If set to `true`, SQLite will be used as the storage method; if set to `false`, other storage methods will be used.
- Usage: Set this flag to `true` if you want to use SQLite as the storage method, and `false` otherwise.

### UsePostgres

- Type: boolean
- Description: This flag determines whether to use PostgreSQL for storage or not. If set to `true`, PostgreSQL will be used as the storage method; if set to `false`, other storage methods will be used.
- Usage: Set this flag to `true` if you want to use PostgreSQL as the storage method, and `false` otherwise.

### GinServerSettings

- Type: struct
- Description: This struct contains configuration options for the Gin server.
- Usage: Use this struct to configure the settings related to the Gin server, such as the port number on which the server should listen.

#### Port

- Type: string
- Description: This option specifies the port number on which the Gin server should listen for incoming requests.
- Usage: Set this option to the desired port number on which you want the Gin server to listen for incoming requests. The port number should be a valid string representation of an integer.

### SqliteSettings

- Type: struct
- Description: This struct contains configuration options for SQLite storage.
- Usage: Use this struct to configure the settings related to SQLite storage, such as the path to the SQLite database file.

#### DBPath

- Type: string
- Description: This option specifies the path to the SQLite database file.
- Usage: Set this option to the absolute or relative path of the SQLite database file that you want to use for storing data.

### PostgresSettings

- Type: struct
- Description: This struct contains configuration options for PostgreSQL storage.
- Usage: Use this struct to configure the settings related to PostgreSQL storage, such as the connection string.

#### ConnectionString

- Type: string
- Description: This option specifies the connection string for connecting to the PostgreSQL database.
- Usage: Set this option to the connection string for connecting to your PostgreSQL database. The connection string should include the necessary information such as host, port, username, password, and database name.

## Example Usage

```yaml
useGinServer: true
useSqlite: true
usePostgres: false
ginServerSettings:
  port: "8080"
sqliteSettings:
  dbPath: "/path/to/sqlite.db"
postgresSettings:
  connectionString: "postgres://username:password@localhost/dbname?sslmode=disable"
```
