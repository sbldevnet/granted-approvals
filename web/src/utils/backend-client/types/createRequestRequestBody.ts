/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * Approvals
 * Granted Approvals API
 * OpenAPI spec version: 1.0
 */
import type { RequestTiming } from './requestTiming';
import type { CreateRequestWith } from './createRequestWith';

export type CreateRequestRequestBody = {
  accessRuleId: string;
  reason?: string;
  timing: RequestTiming;
  with?: CreateRequestWith;
};
