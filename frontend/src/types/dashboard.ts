export interface DashboardCountItem {
  name: string;
  count: number;
}

export interface DashboardTimelineItem {
  time: string;
  count: number;
}

export interface DashboardSummary {
  total_logs: number;
  top_ips: DashboardCountItem[];
  top_users: DashboardCountItem[];
  top_event_types: DashboardCountItem[];
  timeline: DashboardTimelineItem[];
}

export interface DashboardFilter {
  tenant?: string;
  source?: string;
  from?: string;
  to?: string;
}