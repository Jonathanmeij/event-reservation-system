import { cn } from "@/lib/utils";

export default function PageContainer({
  children,
  className
}: {
  children: React.ReactNode;
  className?: string;
}) {
  return (
    <div className="flex flex-col items-stretch flex-1 w-full px-6 bg-zinc-100">
      <div />
      <div
        className={cn(
          "mx-auto flex h-full w-full max-w-screen-2xl flex-1 flex-col",
          className
        )}
      >
        <div />
        {children}
      </div>
    </div>
  );
}
