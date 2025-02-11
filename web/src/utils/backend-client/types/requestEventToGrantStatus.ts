/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * Approvals
 * Granted Approvals API
 * OpenAPI spec version: 1.0
 */

/**
 * The current state of the grant.
 */
export type RequestEventToGrantStatus = typeof RequestEventToGrantStatus[keyof typeof RequestEventToGrantStatus];


// eslint-disable-next-line @typescript-eslint/no-redeclare
export const RequestEventToGrantStatus = {
  PENDING: 'PENDING',
  ACTIVE: 'ACTIVE',
  ERROR: 'ERROR',
  REVOKED: 'REVOKED',
  EXPIRED: 'EXPIRED',
} as const;
