import apiClient from "./client";

import type {
  DashboardFilter,
  DashboardSummary,
} from "../types/dashboard";

export const getDashboard = async (
  filter?: DashboardFilter
): Promise<DashboardSummary> => {
  const response = await apiClient.get<DashboardSummary>(
    "/dashboard",
    {
      params: filter,
    }
  );

  return response.data;
};