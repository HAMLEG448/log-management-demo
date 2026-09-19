import { Card, Typography } from "antd";

import {
  CartesianGrid,
  Line,
  LineChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";

import type {
  DashboardTimelineItem,
} from "../../types/dashboard";

interface TimelineChartProps {
  data: DashboardTimelineItem[];
}

function TimelineChart({
  data,
}: TimelineChartProps) {
  const chartData = data.map((item) => ({
    ...item,
    displayTime: new Date(item.time).toLocaleString(),
  }));

  return (
    <Card title="Log Timeline">
      {chartData.length === 0 ? (
        <Typography.Text type="secondary">
          No timeline data
        </Typography.Text>
      ) : (
        <div
          style={{
            width: "100%",
            height: 320,
          }}
        >
          <ResponsiveContainer>
            <LineChart data={chartData}>
              <CartesianGrid strokeDasharray="3 3" />

              <XAxis
                dataKey="displayTime"
                minTickGap={30}
              />

              <YAxis allowDecimals={false} />

              <Tooltip />

              <Line
                type="monotone"
                dataKey="count"
              />
            </LineChart>
          </ResponsiveContainer>
        </div>
      )}
    </Card>
  );
}

export default TimelineChart;