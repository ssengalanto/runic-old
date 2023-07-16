// Package config provides a configuration management solution for the application.
//
// The config package offers functionalities for managing the application's configuration. It provides a Config struct,
// which represents the application configuration structure, as well as specific configuration sub-structs such as App,
// HTTP, and PGSQL that define different aspects of the application's configuration.
//
// The New function in this package creates a new Config instance with default values loaded from a YAML file and
// environment variables. It initializes a new Config struct, reads configuration values from a YAML file named "config.yml"
// in the current working directory using the cleanenv package, and then reads additional configuration values from
// environment variables using the same package.
//
// The New function returns a pointer to a new Config instance and an error if any occurred during initialization.
// The returned Config struct contains the default configuration values from the YAML file and any overrides from
// environment variables.
package config
