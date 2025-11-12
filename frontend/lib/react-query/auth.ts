"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query";
import {
  SigninRequest,
  SigninResponse,
  SignupRequest,
  SignupResponse,
} from "../types/auth";
import { authAPI, handleAPIError } from "@/services/api";
import { toast } from "sonner";

export const useSignInMutation = () => {
  const queryClient = useQueryClient();
  return useMutation<SigninResponse, Error, SigninRequest>({
    mutationFn: (data: SigninRequest) => authAPI.signIn(data),

    onSuccess: () => {
      toast.success("SignIn Successful!!");
      queryClient.invalidateQueries({ queryKey: ["user"] });
    },

    onError: (error) => {
      const errorMessage = handleAPIError(error);
      toast.error(errorMessage);
    },
  });
};

export const useSingUpMutation = () => {
  const queryClient = useQueryClient();
  return useMutation<SignupResponse, Error, SignupRequest>({
    mutationFn: (data: SignupRequest) => authAPI.signUp(data),

    onSuccess: () => {
      toast.success("Signup Successful!!");
      queryClient.invalidateQueries({ queryKey: ["user"] });
    },

    onError: (error) => {
      const errorMessage = handleAPIError(error);
      toast.error(errorMessage);
    },
  });
};
