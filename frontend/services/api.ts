import {
  ApiError,
  SigninRequest,
  SigninResponse,
  SignupRequest,
  SignupResponse,
} from "@/lib/types/auth";
import axios, { AxiosResponse } from "axios";
import { toast } from "sonner";

export const api = axios.create({
  baseURL: process.env.NEXT_API_URL_DEV,
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 10000,
});

api.interceptors.request.use(
  (config) => {
    if (process.env.NODE_ENV === "development") {
      console.log(`API Request: ${config.method?.toUpperCase()} ${config.url}`);
    }
    return config;
  },
  (error) => {
    console.error("API Request Error:", error);
    return Promise.reject(error);
  }
);

api.interceptors.response.use(
  (response: AxiosResponse) => {
    if (process.env.NODE_ENV === "development") {
      console.log(`API Response: ${response.status} ${response.config.url}`);
    }
    return response;
  },
  (error) => {
    if (!error.response) {
      toast.error("Network. Please check your connection.");
      error.message = "Network error. Please check your connection.";
    }

    if (process.env.NODE_ENV === "development") {
      console.error("API Error:", {
        status: error.response.status,
        message: error.response.data.message || error.message,
        url: error.config.url,
      });
    }
    return Promise.reject(error);
  }
);

export const authAPI = {
  signUp: async (data: SignupRequest): Promise<SignupResponse> => {
    try {
      const response = await api.post<SignupResponse>("/signup", data);
      return response.data;
    } catch (error: any) {
      if (error.response) {
        const apiError: ApiError = error.response.data;

        switch (error.response.status) {
          case 400:
            throw new Error(apiError.message || "Validation failed");
          case 406:
            throw new Error(apiError.message || "Invalid JSON format");
          case 409:
            throw new Error(apiError.message || "Email already exists");
          case 500:
            throw new Error("Internal server error");
          default:
            throw new Error(apiError.message || "Signup failed");
        }
      }

      throw new Error(error.message || "Network error occurred");
    }
  },

  signIn: async (data: SigninRequest): Promise<SigninResponse> => {
    try {
      const response = await api.post<SigninResponse>("/signin", data);
      return response.data;
    } catch (error: any) {
      if (error.response) {
        const apiError: ApiError = error.response.data;

        switch (error.response.status) {
          case 400:
            throw new Error(apiError.message || "User not found");
          case 406:
            throw new Error(apiError.message || "Incorrect password");
          default:
            throw new Error(apiError.message || "Sign-in failed");
        }
      }

      throw new Error(error.message || "Check your connection");
    }
  },
};

export const handleAPIError = (error: any): string => {
  if (error.response?.data?.message) {
    return error.response.data.message;
  }

  if (error.message) {
    return error.message;
  }
  return "An unexpected error occurred";
};
