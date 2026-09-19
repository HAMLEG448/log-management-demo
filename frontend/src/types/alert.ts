export interface AlertItem {
  id: number;
  timestamp: string;

  tenant: string;
  rule_name: string;

  severity: number;

  src_ip: string;
  event_type: string;

  message: string;

  created_at: string;
}

export interface AlertResponse {
  count: number;
  alerts: AlertItem[];
}