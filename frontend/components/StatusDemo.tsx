import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { CheckCircle, Clock, AlertTriangle } from "lucide-react";

const regions = [
  {
    name: "US East",
    location: "Virginia",
    status: "operational",
    latency: "45ms",
  },
  {
    name: "US West",
    location: "California",
    status: "operational",
    latency: "52ms",
  },
  {
    name: "Europe",
    location: "Frankfurt",
    status: "operational",
    latency: "78ms",
  },
  {
    name: "Asia Pacific",
    location: "Singapore",
    status: "degraded",
    latency: "156ms",
  },
  {
    name: "Australia",
    location: "Sydney",
    status: "operational",
    latency: "89ms",
  },
  {
    name: "South America",
    location: "São Paulo",
    status: "operational",
    latency: "112ms",
  },
];

const getStatusIcon = (status: string) => {
  switch (status) {
    case "operational":
      return <CheckCircle className="w-4 h-4 text-success" />;
    case "degraded":
      return <AlertTriangle className="w-4 h-4 text-warning" />;
    default:
      return <Clock className="w-4 h-4 text-muted-foreground" />;
  }
};

const getStatusBadge = (status: string) => {
  switch (status) {
    case "operational":
      return (
        <Badge
          variant="secondary"
          className="bg-success/10 text-success border-success/20"
        >
          Operational
        </Badge>
      );
    case "degraded":
      return (
        <Badge
          variant="secondary"
          className="bg-warning/10 text-warning border-warning/20"
        >
          Degraded
        </Badge>
      );
    default:
      return <Badge variant="secondary">Unknown</Badge>;
  }
};

export const StatusDemo = () => {
  return (
    <section className="py-16 bg-gradient-feature">
      <div className="container">
        <div className="max-w-4xl mx-auto">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold mb-4">Live Status Dashboard</h2>
            <p className="text-muted-foreground">
              See real-time monitoring data from our global network
            </p>
          </div>

          <Card className="shadow-card">
            <CardHeader>
              <CardTitle className="flex items-center justify-between">
                <span>Global Monitor Status</span>
                <Badge
                  variant="secondary"
                  className="bg-success/10 text-success border-success/20"
                >
                  <CheckCircle className="w-4 h-4 mr-1" />
                  All Systems Operational
                </Badge>
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="grid gap-4">
                {regions.map((region) => (
                  <div
                    key={region.name}
                    className="flex items-center justify-between p-4 rounded-lg border bg-card/50 hover:bg-card transition-colors"
                  >
                    <div className="flex items-center space-x-3">
                      {getStatusIcon(region.status)}
                      <div>
                        <div className="font-medium">{region.name}</div>
                        <div className="text-sm text-muted-foreground">
                          {region.location}
                        </div>
                      </div>
                    </div>

                    <div className="flex items-center space-x-4">
                      <div className="text-right">
                        <div className="text-sm font-medium">
                          {region.latency}
                        </div>
                        <div className="text-xs text-muted-foreground">
                          response time
                        </div>
                      </div>
                      {getStatusBadge(region.status)}
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </section>
  );
};
