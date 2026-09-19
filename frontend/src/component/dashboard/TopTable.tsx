import { Card, Table } from "antd";

import type {
  TableColumnsType,
} from "antd";

import type {
  DashboardCountItem,
} from "../../types/dashboard";

interface TopTableProps {
  title: string;
  data: DashboardCountItem[];
  emptyText: string;
}

const columns: TableColumnsType<DashboardCountItem> = [
  {
    title: "Name",
    dataIndex: "name",
    key: "name",
  },
  {
    title: "Count",
    dataIndex: "count",
    key: "count",
    width: 100,
  },
];

function TopTable({
  title,
  data,
  emptyText,
}: TopTableProps) {
  return (
    <Card title={title}>
      <Table
        rowKey="name"
        columns={columns}
        dataSource={data}
        pagination={false}
        locale={{ emptyText }}
      />
    </Card>
  );
}

export default TopTable;