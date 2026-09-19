import apiClient from "./client";

import type { AlertResponse } from "../types/alert";

export const getAlerts = async (): Promise<AlertResponse> => {
  const response = await apiClient.get<AlertResponse>(
    "/alerts"
  );

  return response.data;
};