/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * Approvals
 * Granted Approvals API
 * OpenAPI spec version: 1.0
 */
import type { WithOption } from './withOption';

export interface RequestArgument {
  title: string;
  options: WithOption[];
  description?: string;
  /** This will be true if a selection is require when creating a request */
  requiresSelection: boolean;
}
