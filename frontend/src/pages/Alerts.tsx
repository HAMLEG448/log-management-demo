import { useEffect, useState } from "react";
import {
  Alert as AntAlert,
  Button,
  Card,
  Modal,
  Space,
  Table,
  Tag,
  Typography,
} from "antd";

import type { ColumnsType } from "antd/es/table";

import { getAlerts } from "../api/alerts";
import type { AlertItem } from "../types/alert";

const { Title, Text } = Typography;

function Alerts() {
  const [alerts, setAlerts] = useState<AlertItem[]>([]);
  const [loading, setLoading] = useState(true);

  const [error, setError] = useState<string | null>(null);

  const [selectedAlert, setSelectedAlert] =
    useState<AlertItem | null>(null);

  useEffect(() => {
    const loadAlerts = async () => {
      try {
        setLoading(true);
        setError(null);

        const data = await getAlerts();

        setAlerts(data.alerts);
      } catch (err) {
        console.error("Failed to load alerts:", err);

        setError("Failed to load alerts");
      } finally {
        setLoading(false);
      }
    };

    loadAlerts();
  }, []);

  const getSeverityColor = (severity: number) => {
    if (severity >= 8) {
      return "red";
    }

    if (severity >= 5) {
      return "orange";
    }

    return "blue";
  };

  const columns: ColumnsType<AlertItem> = [
    {
      title: "Time",
      dataIndex: "timestamp",
      key: "timestamp",

      render: (value: string) =>
        new Date(value).toLocaleString(),
    },

    {
      title: "Tenant",
      dataIndex: "tenant",
      key: "tenant",
    },

    {
      title: "Severity",
      dataIndex: "severity",
      key: "severity",

      render: (severity: number) => (
        <Tag color={getSeverityColor(severity)}>
          {severity}
        </Tag>
      ),
    },

    {
      title: "Rule",
      dataIndex: "rule_name",
      key: "rule_name",
    },

    {
      title: "Source IP",
      dataIndex: "src_ip",
      key: "src_ip",
    },

    {
      title: "Event Type",
      dataIndex: "event_type",
      key: "event_type",
    },

    {
      title: "Detail",
      key: "detail",

      render: (_, record) => (
        <Button onClick={() => setSelectedAlert(record)}>
          View
        </Button>
      ),
    },
  ];

  return (
    <div style={{ padding: 24 }}>
      <Title level={2}>Alerts</Title>

      {error && (
        <AntAlert
          message="Error"
          description={error}
          type="error"
          showIcon
          style={{ marginBottom: 16 }}
        />
      )}

      <Card>
        <Table<AlertItem>
          rowKey="id"
          columns={columns}
          dataSource={alerts}
          loading={loading}
          scroll={{
            x: 900,
          }}
        />
      </Card>

      <Modal
        title="Alert Detail"
        open={selectedAlert !== null}
        onCancel={() => setSelectedAlert(null)}
        footer={null}
        width={700}
      >
        {selectedAlert && (
          <Space
            direction="vertical"
            size="middle"
            style={{ width: "100%" }}
          >
            <Text>
              <strong>Rule:</strong>{" "}
              {selectedAlert.rule_name}
            </Text>

            <Text>
              <strong>Severity:</strong>{" "}
              <Tag
                color={getSeverityColor(
                  selectedAlert.severity
                )}
              >
                {selectedAlert.severity}
              </Tag>
            </Text>

            <Text>
              <strong>Tenant:</strong>{" "}
              {selectedAlert.tenant}
            </Text>

            <Text>
              <strong>Source IP:</strong>{" "}
              {selectedAlert.src_ip}
            </Text>

            <Text>
              <strong>Event Type:</strong>{" "}
              {selectedAlert.event_type}
            </Text>

            <Text>
              <strong>Timestamp:</strong>{" "}
              {new Date(
                selectedAlert.timestamp
              ).toLocaleString()}
            </Text>

            <Text strong>Message</Text>

            <div
              style={{
                padding: 12,
                background: "#f5f5f5",
                borderRadius: 6,
              }}
            >
              {selectedAlert.message}
            </div>
          </Space>
        )}
      </Modal>
    </div>
  );
}

export default Alerts;