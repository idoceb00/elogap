import { Card, CardContent } from "../components/ui/card";

export function ComingSoon({
  title,
  description,
}: {
  title: string;
  description?: string;
}) {
  return (
    <div className="space-y-6">
      <div className="mb-8">
        <h1 className="text-3xl font-semibold mb-2">{title}</h1>
        <p className="text-muted-foreground">
          {description || "This page will be available soon."}
        </p>
      </div>

      <Card className="border-primary/20 bg-primary/5">
        <CardContent className="p-6">
          <p className="font-medium">Coming soon</p>
          <p className="text-sm text-muted-foreground mt-1">
            This section is waiting for backend support.
          </p>
        </CardContent>
      </Card>
    </div>
  );
}