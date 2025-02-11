/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * Approvals
 * Granted Approvals API
 * OpenAPI spec version: 1.0
 */
import type { ApproverConfig } from './approverConfig';
import type { CreateAccessRuleTarget } from './createAccessRuleTarget';
import type { TimeConstraints } from './timeConstraints';

export type CreateAccessRuleRequestBody = {
  /** The group IDs that the access rule applies to. */
  groups: string[];
  approval: ApproverConfig;
  name: string;
  description: string;
  target: CreateAccessRuleTarget;
  timeConstraints: TimeConstraints;
};
