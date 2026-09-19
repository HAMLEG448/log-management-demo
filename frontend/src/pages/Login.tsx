import { useState } from "react";

import {
  Card,
  Typography,
} from "antd";

import { login } from "../api/auth";
import LoginForm from "../component/auth/LoginForm";

import { useAuth } from "../contexts/useAuth";

import type {
  LoginRequest,
} from "../types/auth";

const { Title, Text } = Typography;

function Login() {
  const { loginUser } = useAuth();

  const [loading, setLoading] =
    useState(false);

  const [error, setError] =
    useState<string | null>(null);

  const handleLogin = async (
    values: LoginRequest,
  ) => {
    setLoading(true);
    setError(null);

    try {
      const response =
        await login(values);

      loginUser(
        response.token,
        response.user,
      );
    } catch (err) {
      console.error(
        "Login failed:",
        err,
      );

      setError(
        "Invalid username or password",
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <div
      style={{
        minHeight: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        background: "#f5f5f5",
        padding: 20,
      }}
    >
      <Card
        style={{
          width: "100%",
          maxWidth: 420,
        }}
      >
        <div
          style={{
            textAlign: "center",
            marginBottom: 24,
          }}
        >
          <Title
            level={2}
            style={{
              marginBottom: 4,
            }}
          >
            Log Management
          </Title>

          <Text type="secondary">
            Sign in to continue
          </Text>
        </div>

        <LoginForm
          loading={loading}
          error={error}
          onSubmit={handleLogin}
        />
      </Card>
    </div>
  );
}

export default Login;