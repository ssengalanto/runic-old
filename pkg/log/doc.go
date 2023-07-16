// Package log provides a log implementation that uses zap.Logger internally.
//
// The log package offers a Logger type that serves as a logging abstraction with various log levels,
// such as Info, Error, Debug, Warn, Fatal, and Panic. It uses the popular zap.Logger library internally
// for efficient and customizable logging.
//
// To create a new Logger instance, use the New function, providing the desired environment as a parameter.
// The environment parameter helps determine the appropriate configuration for the underlying zap.Logger.
// If the initialization fails, an error will be returned along with a nil Logger instance.
//
// The Logger type satisfies the interfaces.Logger interface, which is defined in the "interfaces" package.
// This allows the Logger to be easily swapped or used with other components that depend on the Logger interface.
//
// Additional fields can be provided when logging messages by passing a map of fields to the respective log method.
// These fields will be included in the logged message to provide more contextual information.
package log
