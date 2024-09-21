import * as React from "react";
import { cn } from "@/lib/utils";
import { cva, VariantProps } from "class-variance-authority";
import { Search } from "lucide-react";

const inputVariants = cva(
  "flex w-full  transition-colors  text-sm focus-visible:outline-none focus-visible:ring-1 disabled:cursor-not-allowed disabled:opacity-50",
  {
    variants: {
      variant: {
        default:
          "border border-zinc-200 h-9 bg-white dark:border-zinc-800 rounded-lg",
        secondary: "border px-6 bg-zinc-200 h-8 rounded-full"
      },
      padding: {
        default: "px-3 py-1"
      }
    },
    defaultVariants: {
      variant: "default"
    }
  }
);

export interface InputProps
  extends React.InputHTMLAttributes<HTMLInputElement>,
    VariantProps<typeof inputVariants> {}

const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({ className, variant, type, padding, ...props }, ref) => {
    return (
      <input
        type={type}
        className={cn(inputVariants({ variant, padding }), className)}
        ref={ref}
        {...props}
      />
    );
  }
);
Input.displayName = "Input";

export interface IconInputProps
  extends React.InputHTMLAttributes<HTMLInputElement>,
    VariantProps<typeof inputVariants> {}

const SearchInput = React.forwardRef<HTMLInputElement, IconInputProps>(
  ({ variant, type, className, padding, ...props }, ref) => {
    //input ref
    const inputRef = React.useRef<HTMLInputElement | null>(null);

    return (
      <div
        className={cn(
          inputVariants({ variant, padding }),
          "flex items-center gap-2 ring-zinc-500 focus-within:ring-1",
          className
        )}
        onClick={() => inputRef.current?.focus()}
      >
        <Search className="size-4 text-zinc-500" />
        <input
          className="w-full p-0 bg-transparent focus:outline-none"
          type={type}
          ref={(el) => {
            inputRef.current = el;
            if (typeof ref === "function") {
              ref(el);
            } else if (ref) {
              (ref as React.MutableRefObject<HTMLInputElement>).current = el!;
            }
          }}
          {...props}
        />
      </div>
    );
  }
);

export { Input, SearchInput };
