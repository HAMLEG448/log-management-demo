import {
  Button,
  Card,
  DatePicker,
  Select,
  Space,
  Tag,
} from "antd";

import type { Dayjs } from "dayjs";

const { RangePicker } = DatePicker;

export interface DashboardFilterState {
  tenant?: string;
  source?: string;
  dateRange: [Dayjs | null, Dayjs | null] | null;
}

interface DashboardFiltersProps {
  filters: DashboardFilterState;
  isAdmin: boolean;
  viewerTenant?: string;

  onChange: (filters: DashboardFilterState) => void;
  onApply: () => void;
  onReset: () => void;
}

function DashboardFilters({
  filters,
  isAdmin,
  viewerTenant,
  onChange,
  onApply,
  onReset,
}: DashboardFiltersProps) {
  const updateFilter = (
    values: Partial<DashboardFilterState>,
  ) => {
    onChange({
      ...filters,
      ...values,
    });
  };

  return (
    <Card style={{ marginBottom: 20 }}>
      <Space wrap>
        {isAdmin ? (
          <Select
            allowClear
            placeholder="Tenant"
            style={{ width: 150 }}
            value={filters.tenant}
            onChange={(tenant) =>
              updateFilter({ tenant })
            }
            options={[
              {
                label: "demoA",
                value: "demoA",
              },
              {
                label: "demoB",
                value: "demoB",
              },
            ]}
          />
        ) : (
          <Tag color="blue">
            Tenant: {viewerTenant}
          </Tag>
        )}

        <Select
          allowClear
          placeholder="Source"
          style={{ width: 160 }}
          value={filters.source}
          onChange={(source) =>
            updateFilter({ source })
          }
          options={[
            {
              label: "API",
              value: "api",
            },
            {
              label: "Firewall",
              value: "firewall",
            },
            {
              label: "Active Directory",
              value: "ad",
            },
            {
              label: "AWS",
              value: "aws",
            },
          ]}
        />

        <RangePicker
          showTime
          value={filters.dateRange}
          onChange={(dateRange) =>
            updateFilter({
              dateRange: dateRange
                ? [dateRange[0], dateRange[1]]
                : null,
            })
          }
        />

        <Button
          type="primary"
          onClick={onApply}
        >
          Apply
        </Button>

        <Button onClick={onReset}>
          Clear
        </Button>
      </Space>
    </Card>
  );
}

export default DashboardFilters;