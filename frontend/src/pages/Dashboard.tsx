import { useEffect, useState } from "react";
import { Alert, Card, Col, Row, Spin, Typography } from "antd";

import { getDashboard } from "../api/dashboard";
import DashboardFilters from "../component/dashboard/DashboardFilters";
import TimelineChart from "../component/dashboard/TimelineChart";
import TopTable from "../component/dashboard/TopTable";
import { useAuth } from "../contexts/useAuth";

import type { DashboardFilterState } from "../component/dashboard/DashboardFilters";
import type {
  DashboardFilter,
  DashboardSummary,
} from "../types/dashboard";

const { Title, Text } = Typography;

const createEmptyFilters = (): DashboardFilterState => ({
  tenant: undefined,
  source: undefined,
  dateRange: null,
});

function Dashboard() {
  const { user } = useAuth();

  const isAdmin = user?.role === "admin";

  const [filters, setFilters] = useState<DashboardFilterState>(
    createEmptyFilters,
  );

  const [dashboard, setDashboard] = useState<DashboardSummary | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const buildQuery = (
    currentFilters: DashboardFilterState,
  ): DashboardFilter => ({
    tenant: isAdmin ? currentFilters.tenant : undefined,
    source: currentFilters.source,
    from: currentFilters.dateRange?.[0]?.toISOString(),
    to: currentFilters.dateRange?.[1]?.toISOString(),
  });

  const loadDashboard = async (
    currentFilters: DashboardFilterState,
  ) => {
    setLoading(true);
    setError(null);

    try {
      const query = buildQuery(currentFilters);
      const data = await getDashboard(query);

      setDashboard(data);
    } catch (err) {
      console.error("Failed to load dashboard:", err);
      setError("Failed to load dashboard data");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    let cancelled = false;

    getDashboard({})
      .then((data) => {
        if (!cancelled) {
          setDashboard(data);
        }
      })
      .catch((err) => {
        console.error("Failed to load dashboard:", err);

        if (!cancelled) {
          setError("Failed to load dashboard data");
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
    void loadDashboard(filters);
  };

  const handleReset = () => {
    const emptyFilters = createEmptyFilters();

    setFilters(emptyFilters);
    void loadDashboard(emptyFilters);
  };

  return (
    <div>
      <Title level={2}>Log Management Dashboard</Title>

      <DashboardFilters
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

      {loading ? (
        <div
          style={{
            display: "flex",
            justifyContent: "center",
            padding: 50,
          }}
        >
          <Spin size="large" />
        </div>
      ) : dashboard ? (
        <>
          <Row gutter={[16, 16]} style={{ marginBottom: 20 }}>
            <Col xs={24} sm={12} md={8} lg={6}>
              <Card>
                <Text type="secondary">Total Logs</Text>

                <Title level={2} style={{ margin: 0 }}>
                  {dashboard.total_logs ?? 0}
                </Title>
              </Card>
            </Col>
          </Row>

          <TimelineChart data={dashboard.timeline ?? []} />

          <Row gutter={[16, 16]} style={{ marginTop: 20 }}>
            <Col xs={24} lg={8}>
              <TopTable
                title="Top IP"
                data={dashboard.top_ips ?? []}
                emptyText="No IP data"
              />
            </Col>

            <Col xs={24} lg={8}>
              <TopTable
                title="Top Users"
                data={dashboard.top_users ?? []}
                emptyText="No user data"
              />
            </Col>

            <Col xs={24} lg={8}>
              <TopTable
                title="Top Event Types"
                data={dashboard.top_event_types ?? []}
                emptyText="No event type data"
              />
            </Col>
          </Row>
        </>
      ) : (
        !error && <Alert type="info" message="No dashboard data" />
      )}
    </div>
  );
}

export default Dashboard;