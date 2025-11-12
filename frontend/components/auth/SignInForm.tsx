import { SignInFormData, signInSchema } from "@/lib/zod/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { motion } from "motion/react";
import Link from "next/link";
import { useForm } from "react-hook-form";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Label } from "../ui/label";

interface ISignInForm {
  onSubmit: (data: SignInFormData) => void;
}

export const SignInForm = ({ onSubmit }: ISignInForm) => {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<SignInFormData>({
    resolver: zodResolver(signInSchema),
  });

  return (
    <motion.form
      noValidate
      initial={{ opacity: 0, x: -20 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: 20 }}
      transition={{ duration: 0.3 }}
      onSubmit={handleSubmit(onSubmit)}
      className="space-y-6"
    >
      <div className="space-y-2">
        <Label htmlFor="signin-email" className="text-sm font-medium">
          Email
        </Label>
        <Input
          id="signin-email"
          type="email"
          placeholder="Enter your email"
          className="h-12"
          {...register("email")}
        />
        {errors.email && (
          <motion.p
            initial={{ opacity: 0, y: -10 }}
            animate={{ opacity: 1, y: 0 }}
            className="text-sm text-destructive"
          >
            {errors.email.message}
          </motion.p>
        )}
      </div>
      <div className="space-y-2">
        <Label htmlFor="signin-password" className="text-sm font-medium">
          Password
        </Label>
        <Input
          id="signin-password"
          type="password"
          placeholder="Enter your password"
          className="h-12"
          {...register("password")}
        />
        {errors.password && (
          <motion.p
            initial={{ opacity: 0, y: -10 }}
            animate={{ opacity: 1, y: 0 }}
            className="text-sm text-destructive"
          >
            {errors.password.message}
          </motion.p>
        )}
      </div>
      <Button
        type="submit"
        className="w-full h-12 font-medium"
        disabled={isSubmitting}
      >
        {isSubmitting ? "Signing In..." : "Sign In"}
      </Button>
      <div className="text-center">
        <Link
          href="#"
          className="text-sm text-muted-foreground hover:text-foreground transition-colors"
        >
          Forgot Password?
        </Link>
      </div>
    </motion.form>
  );
};
