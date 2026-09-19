import { useEffect, useState } from "react";
import { Alert, Typography } from "antd";

import { getLogs } from "../api/logs";
import LogDetailModal from "../component/logs/LogDetailModal";
import LogsFilters from "../component/logs/LogsFilters";
import LogsTable from "../component/logs/LogsTable";
import { useAuth } from "../contexts/useAuth";

import type { LogFilterState } from "../component/logs/LogsFilters";
import type { LogFilter, LogItem } from "../types/log";

const { Title } = Typography;

const createEmptyFilters = (): LogFilterState => ({
  search: "",
  tenant: undefined,
  source: undefined,
  eventType: "",
  dateRange: null,
});

function Logs() {
  const { user } = useAuth();

  const isAdmin = user?.role === "admin";

  const [filters, setFilters] = useState<LogFilterState>(
    createEmptyFilters,
  );

  const [logs, setLogs] = useState<LogItem[]>([]);
  const [selectedLog, setSelectedLog] = useState<LogItem | null>(null);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const buildQuery = (
    currentFilters: LogFilterState,
  ): LogFilter => ({
    search: currentFilters.search || undefined,
    tenant: isAdmin ? currentFilters.tenant : undefined,
    source: currentFilters.source,
    event_type: currentFilters.eventType || undefined,
    from: currentFilters.dateRange?.[0]?.toISOString(),
    to: currentFilters.dateRange?.[1]?.toISOString(),
  });

  const loadLogs = async (
    currentFilters: LogFilterState,
  ) => {
    setLoading(true);
    setError(null);

    try {
      const query = buildQuery(currentFilters);
      const response = await getLogs(query);

      setLogs(response.logs ?? []);
    } catch (err) {
      console.error("Failed to load logs:", err);

      setLogs([]);
      setError("Failed to load logs");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    let cancelled = false;

    getLogs({})
      .then((response) => {
        if (!cancelled) {
          setLogs(response.logs ?? []);
        }
      })
      .catch((err) => {
        console.error("Failed to load logs:", err);

        if (!cancelled) {
          setLogs([]);
          setError("Failed to load logs");
        }
      })
      .finally(() => {
        if (!cancelled) {
          setLoading(false);
        }
      });

    return () => {
      cancelled = true;
    };
  }, []);

  const handleApply = () => {
    void loadLogs(filters);
  };

  const handleReset = () => {
    const emptyFilters = createEmptyFilters();

    setFilters(emptyFilters);
    void loadLogs(emptyFilters);
  };

  return (
    <div>
      <Title level={2}>Logs</Title>

      <LogsFilters
        filters={filters}
        isAdmin={isAdmin}
        viewerTenant={user?.tenant}
        onChange={setFilters}
        onApply={handleApply}
        onReset={handleReset}
      />

      {error && (
        <Alert
          type="error"
          message={error}
          showIcon
          style={{ marginBottom: 20 }}
        />
      )}

      <LogsTable
        logs={logs}
        loading={loading}
        onView={setSelectedLog}
      />

      <LogDetailModal
        log={selectedLog}
        onClose={() => setSelectedLog(null)}
      />
    </div>
  );
}

export default Logs;