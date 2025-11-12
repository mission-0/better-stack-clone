"use client";
import { motion, Variants } from "motion/react";
import { CheckCircle, XCircle } from "lucide-react";

export const StatusSection = () => {
  const regions = [
    { name: "US East", status: "up", responseTime: "45ms" },
    { name: "US West", status: "up", responseTime: "38ms" },
    { name: "Europe", status: "up", responseTime: "52ms" },
    { name: "Asia Pacific", status: "down", responseTime: "---" },
  ];

  const containerVariants: Variants = {
    hidden: { opacity: 0 },
    visible: {
      opacity: 1,
      transition: {
        staggerChildren: 0.1,
      },
    },
  };

  const itemVariants: Variants = {
    hidden: { opacity: 0, y: 20 },
    visible: {
      opacity: 1,
      y: 0,
      transition: { duration: 0.5 },
    },
  };

  return (
    <section className="py-32 px-8">
      <div className="container mx-auto max-w-4xl">
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 0.8 }}
          className="text-center mb-16"
        >
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Global monitoring
          </h2>
          <p className="text-xl text-muted-foreground">
            Real-time status from multiple regions worldwide
          </p>
        </motion.div>

        <motion.div
          variants={containerVariants}
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true }}
          className="bg-gradient-card border border-border/50 rounded-2xl p-8 shadow-card backdrop-blur-sm"
        >
          <div className="grid gap-4">
            {regions.map((region, index) => (
              <motion.div
                key={region.name}
                variants={itemVariants}
                className="flex items-center justify-between py-4 border-b border-border last:border-b-0"
              >
                <div className="flex items-center space-x-4">
                  {region.status === "up" ? (
                    <CheckCircle className="w-5 h-5 text-success" />
                  ) : (
                    <XCircle className="w-5 h-5 text-destructive" />
                  )}
                  <span className="font-medium text-lg">{region.name}</span>
                </div>
                <div className="flex items-center space-x-6">
                  <span
                    className={`text-sm font-medium ${
                      region.status === "up"
                        ? "text-success"
                        : "text-destructive"
                    }`}
                  >
                    {region.status === "up" ? "Operational" : "Down"}
                  </span>
                  <span className="text-sm text-muted-foreground font-mono">
                    {region.responseTime}
                  </span>
                </div>
              </motion.div>
            ))}
          </div>
        </motion.div>
      </div>
    </section>
  );
};
