import { createContext, ReactNode, useContext, useState } from "react";

const authContext = createContext<{
  authCode: string;
  login: (authCode: string) => void;
  logout: () => void;
} | null>(null);

export default function AuthProvider({ children }: { children: ReactNode }) {
  const [authCode, setAuthCode] = useState("");

  const login = (authCode: string) => setAuthCode(authCode);
  const logout = () => setAuthCode("");

  return (
    <authContext.Provider value={{ authCode, login, logout }}>
      {" "}
      {children}{" "}
    </authContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(authContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
