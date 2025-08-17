"use client";

import { motion } from "motion/react";
import { Globe, Activity } from "lucide-react";

export const Features = () => {
  const features = [
    {
      icon: Globe,
      title: "Global Monitoring",
      description:
        "Monitor from multiple regions worldwide to ensure reliable uptime detection.",
    },
    {
      icon: Activity,
      title: "Instant Alerts",
      description:
        "Get notified immediately when your website goes down or comes back online.",
    },
  ];

  const containerVariants = {
    hidden: { opacity: 0 },
    visible: {
      opacity: 1,
      transition: {
        staggerChildren: 0.2,
      },
    },
  };

  const itemVariants = {
    hidden: { opacity: 0, y: 30 },
    visible: {
      opacity: 1,
      y: 0,
      transition: { duration: 0.6 },
    },
  };

  return (
    <section id="features" className="py-32 px-8">
      <div className="container mx-auto max-w-6xl">
        <motion.div
          initial={{ opacity: 0, y: 30 }}
          whileInView={{ opacity: 1, y: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 0.8 }}
          className="text-center mb-20"
        >
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            Simple monitoring
          </h2>
          <p className="text-xl text-muted-foreground max-w-2xl mx-auto">
            Everything you need to keep your websites running smoothly.
          </p>
        </motion.div>

        <motion.div
          variants={containerVariants}
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true }}
          className="grid grid-cols-1 md:grid-cols-2 gap-16 max-w-4xl mx-auto"
        >
          {features.map((feature, index) => (
            <motion.div
              key={index}
              variants={itemVariants}
              className="text-center group"
            >
              <motion.div
                className="w-16 h-16 mx-auto mb-6 rounded-2xl bg-gradient-card border border-border/50 flex items-center justify-center group-hover:bg-gradient-primary group-hover:text-primary-foreground transition-all duration-300 shadow-subtle group-hover:shadow-glow"
                whileHover={{ scale: 1.05 }}
                transition={{ duration: 0.2 }}
              >
                <feature.icon className="w-8 h-8" />
              </motion.div>
              <h3 className="text-2xl font-semibold mb-4">{feature.title}</h3>
              <p className="text-lg text-muted-foreground leading-relaxed">
                {feature.description}
              </p>
            </motion.div>
          ))}
        </motion.div>
      </div>
    </section>
  );
};
