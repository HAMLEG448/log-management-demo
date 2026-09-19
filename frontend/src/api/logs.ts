import apiClient from "./client";

import type {
  LogFilter,
  LogResponse,
} from "../types/log";

export const getLogs = async (
  filter?: LogFilter
): Promise<LogResponse> => {
  const response = await apiClient.get<LogResponse>(
    "/logs",
    {
      params: filter,
    }
  );

  return response.data;
};