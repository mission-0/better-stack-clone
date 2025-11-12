"use client";

import { useState } from "react";
import { motion } from "framer-motion";

import { AuthModal } from "@/components/auth/AuthModal";
import { Button } from "./ui/button";

export const Header = () => {
  const [showAuthModal, setShowAuthModal] = useState(false);

  return (
    <>
      <motion.header
        initial={{ y: -100 }}
        animate={{ y: 0 }}
        transition={{ duration: 0.6, ease: "easeOut" }}
        className="sticky top-0 z-50 w-full border-b border-border/50 bg-gradient-secondary/90 backdrop-blur-xl shadow-subtle"
      >
        <div className="container mx-auto flex h-20 items-center justify-between px-8">
          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.2, duration: 0.5 }}
            className="flex items-center"
          >
            <span className="font-semibold text-2xl tracking-tight">
              UptimeGuard
            </span>
          </motion.div>

          <nav className="hidden md:flex items-center space-x-12">
            <motion.a
              href="#features"
              className="text-muted-foreground hover:text-foreground transition-colors duration-200 font-medium"
              whileHover={{ y: -2 }}
              transition={{ duration: 0.2 }}
            >
              Features
            </motion.a>
          </nav>

          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.4, duration: 0.5 }}
          >
            <Button
              variant="ghost"
              onClick={() => setShowAuthModal(true)}
              className="font-medium hover:bg-secondary/50"
            >
              Sign In
            </Button>
          </motion.div>
        </div>
      </motion.header>

      <AuthModal open={showAuthModal} onOpenChange={setShowAuthModal} />
    </>
  );
};
