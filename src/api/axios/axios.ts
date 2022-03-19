import axios from "axios";

const instance = axios.create({
  baseURL: import.meta.env.BASE_URL || "http://localhost:8080",
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
});

export const req = instance;
