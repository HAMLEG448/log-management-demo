import {
  Button,
  Card,
  DatePicker,
  Input,
  Select,
  Space,
  Tag,
} from "antd";

import type { Dayjs } from "dayjs";

const { RangePicker } = DatePicker;

export interface LogFilterState {
  search: string;
  tenant?: string;
  source?: string;
  eventType: string;
  dateRange: [Dayjs | null, Dayjs | null] | null;
}

interface LogsFiltersProps {
  filters: LogFilterState;
  isAdmin: boolean;
  viewerTenant?: string;

  onChange: (filters: LogFilterState) => void;
  onApply: () => void;
  onReset: () => void;
}

function LogsFilters({
  filters,
  isAdmin,
  viewerTenant,
  onChange,
  onApply,
  onReset,
}: LogsFiltersProps) {
  const updateFilter = (
    values: Partial<LogFilterState>,
  ) => {
    onChange({
      ...filters,
      ...values,
    });
  };

  return (
    <Card style={{ marginBottom: 20 }}>
      <Space wrap>
        <Input
          placeholder="Search logs"
          style={{ width: 200 }}
          value={filters.search}
          onChange={(event) =>
            updateFilter({
              search: event.target.value,
            })
          }
          onPressEnter={onApply}
        />

        {isAdmin ? (
          <Select
            allowClear
            placeholder="Tenant"
            style={{ width: 140 }}
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

        <Input
          placeholder="Event Type"
          style={{ width: 180 }}
          value={filters.eventType}
          onChange={(event) =>
            updateFilter({
              eventType: event.target.value,
            })
          }
          onPressEnter={onApply}
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

export default LogsFilters;