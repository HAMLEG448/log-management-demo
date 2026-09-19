import { createContext } from "react";

import type { User } from "../types/auth";

export interface AuthContextType {
  user: User | null;
  token: string | null;

  loginUser: (
    token: string,
    user: User
  ) => void;

  logout: () => void;
}

export const AuthContext =
  createContext<AuthContextType | undefined>(
    undefined
  );