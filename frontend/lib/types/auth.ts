export interface SignupRequest {
  email: string;
  password: string;
  fullname: string;
}

export interface SignupResponse {
  message: string;
  user: {
    id: string;
    email: string;
    fullname: string;
    password: string;
  };
}

export interface ApiError {
  message: string;
  error?: string;
  errors?: string[];
}

export interface SigninRequest {
  id: string;
  password: string;
}

export interface SigninResponse {
  message: string;
  JWT: string;
}
