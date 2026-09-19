import {
  Button,
  Table,
  Tag,
} from "antd";

import type {
  TableColumnsType,
} from "antd";

import type {
  LogItem,
} from "../../types/log";

interface LogsTableProps {
  logs: LogItem[];
  loading: boolean;

  onView: (log: LogItem) => void;
}

function LogsTable({
  logs,
  loading,
  onView,
}: LogsTableProps) {
  const columns: TableColumnsType<LogItem> = [
    {
      title: "Time",
      dataIndex: "@timestamp",
      key: "timestamp",
      width: 200,

      render: (value: string) =>
        value
          ? new Date(value).toLocaleString()
          : "-",
    },
    {
      title: "Tenant",
      dataIndex: "tenant",
      key: "tenant",
      width: 100,

      render: (value: string) => (
        <Tag color="blue">
          {value}
        </Tag>
      ),
    },
    {
      title: "Source",
      dataIndex: "source",
      key: "source",
      width: 120,

      render: (value: string) => (
        <Tag>{value}</Tag>
      ),
    },
    {
      title: "Event Type",
      dataIndex: "event_type",
      key: "event_type",
    },
    {
      title: "User",
      dataIndex: "user",
      key: "user",

      render: (value?: string) =>
        value || "-",
    },
    {
      title: "Source IP",
      dataIndex: "src_ip",
      key: "src_ip",

      render: (value?: string) =>
        value || "-",
    },
    {
      title: "Action",
      dataIndex: "action",
      key: "action",

      render: (value?: string) =>
        value || "-",
    },
    {
      title: "",
      key: "view",
      width: 90,

      render: (_, record) => (
        <Button
          type="link"
          onClick={() => onView(record)}
        >
          View
        </Button>
      ),
    },
  ];

  return (
    <Table
      rowKey="id"
      loading={loading}
      columns={columns}
      dataSource={logs}
      scroll={{ x: 1100 }}
    />
  );
}

export default LogsTable;