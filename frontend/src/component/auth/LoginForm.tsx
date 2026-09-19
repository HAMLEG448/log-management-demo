import { Alert, Button, Form, Input } from "antd";

import type { LoginRequest } from "../../types/auth";

interface LoginFormProps {
  loading: boolean;
  error: string | null;
  onSubmit: (values: LoginRequest) => void;
}

function LoginForm({
  loading,
  error,
  onSubmit,
}: LoginFormProps) {
  return (
    <>
      {error && (
        <Alert
          type="error"
          message={error}
          showIcon
          style={{ marginBottom: 16 }}
        />
      )}

      <Form<LoginRequest>
        layout="vertical"
        onFinish={onSubmit}
      >
        <Form.Item
          label="Username"
          name="username"
          rules={[
            {
              required: true,
              message: "Please enter username",
            },
          ]}
        >
          <Input
            placeholder="Username"
            autoComplete="username"
          />
        </Form.Item>

        <Form.Item
          label="Password"
          name="password"
          rules={[
            {
              required: true,
              message: "Please enter password",
            },
          ]}
        >
          <Input.Password
            placeholder="Password"
            autoComplete="current-password"
          />
        </Form.Item>

        <Button
          type="primary"
          htmlType="submit"
          loading={loading}
          block
        >
          Login
        </Button>
      </Form>
    </>
  );
}

export default LoginForm;