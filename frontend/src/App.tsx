import {
  Button,
  Space,
  Tabs,
  Tag,
  Typography,
} from "antd";

import Dashboard from "./pages/Dashboard";
import Logs from "./pages/Logs";
import Alerts from "./pages/Alerts";
import Login from "./pages/Login";

import { useAuth } from "./contexts/useAuth";

const { Text } = Typography;

function App() {
  const {
    user,
    token,
    logout,
  } = useAuth();

  if (!token || !user) {
    return <Login />;
  }

  return (
    <div>
      <div
        style={{
          display: "flex",
          justifyContent:
            "space-between",
          alignItems: "center",
          padding: "16px 24px 0 24px",
        }}
      >
        <Space>
          <Text strong>
            {user.username}
          </Text>

          <Tag>
            {user.role.toUpperCase()}
          </Tag>

          {user.tenant && (
            <Tag color="blue">
              {user.tenant}
            </Tag>
          )}
        </Space>

        <Button
          danger
          onClick={logout}
        >
          Logout
        </Button>
      </div>

      <Tabs
        defaultActiveKey="dashboard"
        style={{
          padding: "0 24px",
        }}
        items={[
          {
            key: "dashboard",
            label: "Dashboard",
            children: <Dashboard />,
          },
          {
            key: "logs",
            label: "Logs",
            children: <Logs />,
          },
          {
            key: "alerts",
            label: "Alerts",
            children: <Alerts />,
          },
        ]}
      />
    </div>
  );
}

export default App;