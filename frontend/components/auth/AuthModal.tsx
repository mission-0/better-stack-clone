import { AnimatePresence, motion } from "motion/react";
import { Dialog, DialogContent, DialogHeader, DialogTitle } from "../ui/dialog";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../ui/tabs";
import { useEffect, useState } from "react";
import { SignUpForm } from "./SingUpForm";
import { SignInForm } from "./SignInForm";

interface AuthModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  defaultTab?: "signin" | "signup";
}

export const AuthModal = ({
  open,
  onOpenChange,
  defaultTab = "signin",
}: AuthModalProps) => {
  const [activeTab, setActiveTab] = useState(defaultTab);

  useEffect(() => {
    setActiveTab(defaultTab);
  }, [defaultTab, open]);

  const handleSignIn = () => {};
  const handleSignUp = () => {};

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="sm:max-w-md border-border/50 shadow-large bg-gradient-card backdrop-blur-xl">
        <motion.div
          initial={{ opacity: 0, scale: 0.95 }}
          animate={{ opacity: 1, scale: 1 }}
          exit={{ opacity: 0, scale: 0.95 }}
          transition={{ duration: 0.2 }}
        >
          <DialogHeader className="pb-6">
            <DialogTitle className="text-center text-2xl font-semibold">
              {activeTab == "signin" ? "Welcome back" : "Create Account"}
            </DialogTitle>
          </DialogHeader>

          <Tabs
            value={activeTab}
            onValueChange={(value) => {
              if (value === "signin" || value === "signup") setActiveTab(value);
            }}
            className="w-full"
          >
            <TabsList className="grid w-full grid-cols-2 mb-8">
              <TabsTrigger value="signin" className="font-medium">
                Sign In
              </TabsTrigger>
              <TabsTrigger value="signup" className="font-medium">
                Sign Up
              </TabsTrigger>
            </TabsList>

            <AnimatePresence mode="wait">
              <TabsContent value="signin" key="signin">
                <SignInForm onSubmit={handleSignIn} />
              </TabsContent>

              <TabsContent value="signup" key="signup">
                <SignUpForm onSubmit={handleSignUp} />
              </TabsContent>
            </AnimatePresence>
          </Tabs>
        </motion.div>
      </DialogContent>
    </Dialog>
  );
};
