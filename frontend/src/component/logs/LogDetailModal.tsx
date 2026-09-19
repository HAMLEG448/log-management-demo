import {
  Button,
  Descriptions,
  Modal,
  Typography,
} from "antd";

import type {
  LogItem,
} from "../../types/log";

const { Title } = Typography;

interface LogDetailModalProps {
  log: LogItem | null;
  onClose: () => void;
}

function LogDetailModal({
  log,
  onClose,
}: LogDetailModalProps) {
  if (!log) {
    return null;
  }

  return (
    <Modal
      title="Log Detail"
      open
      width={850}
      onCancel={onClose}
      footer={
        <Button onClick={onClose}>
          Close
        </Button>
      }
    >
      <Descriptions
        bordered
        column={2}
        size="small"
      >
        <Descriptions.Item label="ID">
          {log.id}
        </Descriptions.Item>

        <Descriptions.Item label="Time">
          {log["@timestamp"]
            ? new Date(
                log["@timestamp"],
              ).toLocaleString()
            : "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Tenant">
          {log.tenant || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Source">
          {log.source || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Vendor">
          {log.vendor || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Product">
          {log.product || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Event Type">
          {log.event_type || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Event ID">
          {log.event_id ?? "-"}
        </Descriptions.Item>

        <Descriptions.Item label="User">
          {log.user || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Host">
          {log.host || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Source IP">
          {log.src_ip || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Source Port">
          {log.src_port ?? "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Destination IP">
          {log.dst_ip || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Destination Port">
          {log.dst_port ?? "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Protocol">
          {log.protocol || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Action">
          {log.action || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Reason">
          {log.reason || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Severity">
          {log.severity ?? "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Cloud Account">
          {log.cloud_account_id || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Cloud Region">
          {log.cloud_region || "-"}
        </Descriptions.Item>

        <Descriptions.Item label="Cloud Service">
          {log.cloud_service || "-"}
        </Descriptions.Item>
      </Descriptions>

      <Title
        level={5}
        style={{ marginTop: 20 }}
      >
        Raw Log
      </Title>

      <pre
        style={{
          background: "#f5f5f5",
          padding: 12,
          borderRadius: 6,
          overflowX: "auto",
          whiteSpace: "pre-wrap",
          wordBreak: "break-word",
        }}
      >
        {log.raw || "-"}
      </pre>
    </Modal>
  );
}

export default LogDetailModal;