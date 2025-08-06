import { SignUpFormData, signUpSchema } from "@/lib/zod/auth";
import { motion } from "motion/react";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

interface ISignUpForm {
  onSubmit: (data: SignUpFormData) => void;
}

export const SignUpForm = ({ onSubmit }: ISignUpForm) => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors, isSubmitting },
  } = useForm<SignUpFormData>({
    resolver: zodResolver(signUpSchema),
  });

  const password = watch("password");
  const confirmPassword = watch("confirmPassword");

  return (
    <motion.form
      noValidate
      initial={{ opacity: 0, x: 20 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: -20 }}
      transition={{ duration: 0.3 }}
      onSubmit={handleSubmit(onSubmit)}
      className="space-y-6"
    >
      <div className="space-y-2">
        <Label htmlFor="full-name" className="text-sm font-medium">
          Full Name
        </Label>
        <Input
          id="full-name"
          type="text"
          placeholder="Enter your full name"
          className="h-12"
          {...register("fullname")}
        />
        {errors.fullname && (
          <motion.p
            initial={{ opacity: 0, y: -10 }}
            animate={{ opacity: 1, y: 0 }}
            className="text-sm text-destructive"
          >
            {errors.fullname.message}
          </motion.p>
        )}
      </div>
      <div className="space-y-2">
        <Label htmlFor="signup-email" className="text-sm font-medium">
          Email
        </Label>
        <Input
          id="signup-email"
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
        <Label htmlFor="signup-password" className="text-sm font-medium">
          Password
        </Label>
        <Input
          id="signup-password"
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
      <div className="space-y-2">
        <Label htmlFor="confirm-password" className="text-sm font-medium">
          Confirm Password
        </Label>
        <Input
          id="confirm-password"
          type="password"
          placeholder="Confirm your password"
          className="h-12"
          {...register("confirmPassword")}
        />
        {errors.confirmPassword && (
          <motion.p
            initial={{ opacity: 0, y: -10 }}
            animate={{ opacity: 1, y: 0 }}
            className="text-sm text-destructive"
          >
            {errors.confirmPassword.message}
          </motion.p>
        )}
      </div>
      <Button
        type="submit"
        className="w-full h-12 font-medium"
        disabled={isSubmitting}
      >
        {isSubmitting ? "Creating Account" : "Create Account"}
      </Button>
    </motion.form>
  );
};
