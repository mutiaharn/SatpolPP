import sharedConstants from '$contracts/constants.json';
import type { ComplaintStatus, Regency, UserRole } from '$lib/types';

export const SHARED_CONSTANTS = sharedConstants;

export const STATUS_LABELS = sharedConstants.complaint_statuses as Record<ComplaintStatus, string>;

export const ROLE_LABELS = sharedConstants.roles as Record<UserRole, string>;

export const MASTER_REGENCIES: Regency[] = sharedConstants.regencies;

export const UPLOAD_LIMITS = sharedConstants.upload_limits;
