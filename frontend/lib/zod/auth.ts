import * as z from "zod";

export const signInSchema = z.object({
  email: z.email("Please enter a valid email address."),
  password: z.string().min(8, "Password must be at least 6 characters"),
});

export const signUpSchema = z
  .object({
    fullname: z
      .string()
      .min(2, "Full name must be at least 2 characters")
      .max(18, "Name cannot exceed more than 18 characters"),
    email: z.email("Please enter a valid email address"),
    password: z
      .string()
      .min(8, "Please must be at least 8 characters")
      .max(15, "Password cannot exceed more than 15 characters"),
    confirmPassword: z.string(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Password do not match",
    path: ["confirmPassword"],
  });

export type SignInFormData = z.infer<typeof signInSchema>;
export type SignUpFormData = z.infer<typeof signUpSchema>;
