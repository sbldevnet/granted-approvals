/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * Approvals
 * Granted Approvals API
 * OpenAPI spec version: 1.0
 */
import type { LogLevel } from './logLevel';

/**
 * A log entry.
 */
export interface Log {
  /** The log level. */
  level: LogLevel;
  /** The log message. */
  msg: string;
}
