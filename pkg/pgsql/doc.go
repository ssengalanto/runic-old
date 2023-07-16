// Package pgsql provides a concurrency-safe connection pool for PostgreSQL databases.
//
// The pgsql package offers a simple and concurrency-safe connection pool for connecting to PostgreSQL
// databases. It utilizes the sqlx library internally along with the pgx SQL driver.
//
// To establish a new connection, use the NewConnection function, which creates a new PostgreSQL database
// connection based on the provided parameters. The function takes the username, password, host address, port,
// database name, and optional query parameters. It returns a pointer to sqlx.DB, which represents the
// database instance, and an error if any.
//
// The package automatically sets the SQL driver to "pgx" for connecting to PostgreSQL databases.
package pgsql
