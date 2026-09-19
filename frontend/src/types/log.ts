export interface LogItem {
  id: number;
  "@timestamp": string;

  tenant: string;
  source: string;

  vendor?: string;
  product?: string;

  event_type: string;
  event_id?: number;
  logon_type?: number;

  severity?: number;
  action?: string;

  src_ip?: string;
  src_port?: number;

  dst_ip?: string;
  dst_port?: number;

  protocol?: string;

  user?: string;
  host?: string;
  reason?: string;

  cloud_account_id?: string;
  cloud_region?: string;
  cloud_service?: string;

  raw?: string;

  created_at: string;
}

export interface LogResponse {
  count: number;
  logs: LogItem[];
}

export interface LogFilter {
  tenant?: string;
  source?: string;
  event_type?: string;
  user?: string;
  search?: string;
  from?: string;
  to?: string;
}