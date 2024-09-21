import * as React from "react";

import { cn } from "@/lib/utils";

const Separator = React.forwardRef<
  HTMLDivElement,
  React.ComponentPropsWithoutRef<"div">
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    className={cn("bg-zinc-100 dark:bg-zinc-800", "h-[1px] w-full", className)}
    {...props}
  />
));
Separator.displayName = "Separator";

export { Separator };
