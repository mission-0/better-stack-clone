"use client";

import { useState } from "react";
import { motion } from "framer-motion";
import { Button } from "@/components/ui/button";
import { AuthModal } from "@/components/auth/AuthModal";

export const Hero = () => {
  const [showAuthModal, setShowAuthModal] = useState(false);
  const [authMode, setAuthMode] = useState<"signin" | "signup">("signup");

  const handleStartMonitoring = () => {
    setAuthMode("signup");
    setShowAuthModal(true);
  };

  const handleSignIn = () => {
    setAuthMode("signin");
    setShowAuthModal(true);
  };

  return (
    <>
      <section className="relative min-h-screen flex items-center justify-center px-8 bg-gradient-hero overflow-hidden">
        <div className="absolute inset-0 bg-gradient-hero-overlay" />
        <div className="glitter-overlay" />
        <div className="container mx-auto max-w-4xl text-center relative z-10">
          <motion.div
            initial={{ opacity: 0, y: 30 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8, ease: "easeOut" }}
          >
            <h1 className="text-5xl md:text-7xl lg:text-8xl font-bold tracking-tight mb-8 leading-tight">
              Monitor your site
              <br />
              <span className="text-muted-foreground">from everywhere</span>
            </h1>

            <motion.p
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: 0.3, duration: 0.8 }}
              className="text-xl md:text-2xl text-muted-foreground mb-12 max-w-2xl mx-auto leading-relaxed"
            >
              Simple, reliable uptime monitoring from multiple global regions.
              Get instant alerts when your website goes down.
            </motion.p>

            <motion.div
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: 0.6, duration: 0.8 }}
              className="flex flex-col sm:flex-row gap-4 justify-center items-center"
            >
              <Button
                size="lg"
                className="text-lg px-12 py-6 h-auto font-semibold"
                onClick={handleStartMonitoring}
              >
                Start Monitoring for Free
              </Button>
              <Button
                variant="outline"
                size="lg"
                className="text-lg px-12 py-6 h-auto font-semibold"
                onClick={handleSignIn}
              >
                Sign In
              </Button>
            </motion.div>
          </motion.div>
        </div>
      </section>

      <AuthModal
        open={showAuthModal}
        onOpenChange={setShowAuthModal}
        defaultTab={authMode}
      />
    </>
  );
};
