export const Footer = () => {
  return (
    <footer className="border-t border-border bg-background">
      <div className="container mx-auto px-8 py-16">
        <div className="flex flex-col md:flex-row justify-between items-start md:items-center space-y-8 md:space-y-0">
          <div>
            <span className="font-semibold text-xl">UptimeGuard</span>
            <p className="text-muted-foreground mt-2 max-w-md">
              Simple, reliable website monitoring from multiple global regions.
            </p>
          </div>

          <div className="flex flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-12">
            <div className="space-y-2">
              <a
                href="#"
                className="text-muted-foreground hover:text-foreground transition-colors block"
              >
                Privacy Policy
              </a>
              <a
                href="#"
                className="text-muted-foreground hover:text-foreground transition-colors block"
              >
                Terms of Service
              </a>
            </div>
            <div className="space-y-2">
              <a
                href="#"
                className="text-muted-foreground hover:text-foreground transition-colors block"
              >
                Support
              </a>
              <a
                href="#"
                className="text-muted-foreground hover:text-foreground transition-colors block"
              >
                Contact
              </a>
            </div>
          </div>
        </div>

        <div className="border-t border-border mt-16 pt-8">
          <p className="text-sm text-muted-foreground text-center">
            © 2024 UptimeGuard. All rights reserved.
          </p>
        </div>
      </div>
    </footer>
  );
};
