// components/CustomCard.tsx
import { ReactNode } from "react";
import clsx from "clsx";

interface CustomCardProps {
    children: ReactNode;
    className?: string;
    title?: string;
    subtitle?: string;
    action?: ReactNode;
}

export default function CustomCard({
    children,
    className = "",
    title,
    subtitle,
    action,
}: CustomCardProps) {
    return (
        <div className="w-full col-span-12">
            <div className="w-full max-w p-4 bg-white rounded-lg shadow dark:bg-gray-800 dark:border dark:border-gray-700">
                {(title || action || subtitle) && (
                    <div className="mb-4 flex flex-col sm:flex-row sm:items-center sm:justify-between gap-2">
                        <div>
                            {title && (
                                <h2 className="text-2xl font-semibold text-gray-900 dark:text-white">
                                    {title}
                                </h2>
                            )}
                            {subtitle && (
                                <p className="text-sm text-gray-600 dark:text-gray-400">
                                    {subtitle}
                                </p>
                            )}
                        </div>
                        {action && <div>{action}</div>}
                    </div>
                )}
                {/* 🛠️ Here's where we apply the grid classes */}
                <div className={clsx("grid", className)}>
                    {children}
                </div>
            </div>
        </div>
    );
}
