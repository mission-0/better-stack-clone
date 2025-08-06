import { Card, CardContent } from "@/components/ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

const testimonials = [
  {
    name: "Sarah Chen",
    role: "CTO at TechFlow",
    avatar: "/api/placeholder/40/40",
    content:
      "UptimeGuard has been crucial for our SaaS platform. The instant alerts and global monitoring give us confidence that we'll know about issues before our customers do.",
  },
  {
    name: "Marcus Rodriguez",
    role: "DevOps Lead at DataSync",
    avatar: "/api/placeholder/40/40",
    content:
      "The detailed analytics and SSL monitoring features are exactly what we needed. Setup was incredibly simple and the team loves the mobile notifications.",
  },
  {
    name: "Emily Watson",
    role: "Founder at CloudMart",
    avatar: "/api/placeholder/40/40",
    content:
      "Since switching to UptimeGuard, our incident response time has decreased by 75%. The multi-channel alerts ensure our team never misses a critical issue.",
  },
];

const clientLogos = [
  "TechFlow",
  "DataSync",
  "CloudMart",
  "WebScale",
  "DevOps Pro",
  "SiteGuard",
];

export const Testimonials = () => {
  return (
    <section className="py-20 bg-gradient-feature">
      <div className="container">
        <div className="max-w-4xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-3xl font-bold mb-4">
              Trusted by 1000+ Companies
            </h2>
            <p className="text-muted-foreground">
              Join thousands of developers and businesses who rely on
              UptimeGuard
            </p>
          </div>

          <div className="flex flex-wrap justify-center items-center gap-8 mb-16 opacity-60">
            {clientLogos.map((logo) => (
              <div
                key={logo}
                className="text-lg font-semibold text-muted-foreground hover:text-foreground transition-colors"
              >
                {logo}
              </div>
            ))}
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {testimonials.map((testimonial) => (
              <Card key={testimonial.name} className="shadow-card border-0">
                <CardContent className="p-6">
                  <p className="text-muted-foreground mb-6 italic">
                    "{testimonial.content}"
                  </p>
                  <div className="flex items-center space-x-3">
                    <Avatar>
                      <AvatarImage
                        src={testimonial.avatar}
                        alt={testimonial.name}
                      />
                      <AvatarFallback>
                        {testimonial.name
                          .split(" ")
                          .map((n) => n[0])
                          .join("")}
                      </AvatarFallback>
                    </Avatar>
                    <div>
                      <div className="font-semibold">{testimonial.name}</div>
                      <div className="text-sm text-muted-foreground">
                        {testimonial.role}
                      </div>
                    </div>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      </div>
    </section>
  );
};
