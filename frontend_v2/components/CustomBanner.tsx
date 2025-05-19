import React from "react";
import classNames from "classnames";

type BannerType = "success" | "warning" | "error" | "info";

interface CustomBannerProps extends React.HTMLAttributes<HTMLDivElement> {
    type?: BannerType;
    text: string;
    className?: string;
}

const typeStyles: Record<BannerType, string> = {
    success: "bg-green-100 border-green-500 text-green-800",
    warning: "bg-yellow-100 border-yellow-500 text-yellow-800",
    error: "bg-red-100 border-red-500 text-red-800",
    info: "bg-blue-100 border-blue-500 text-blue-800",
};

const iconMap: Record<BannerType, React.ReactNode> = {
    success: (
        <svg className="w-5 h-5 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
        </svg>
    ),
    warning: (
        <svg className="w-5 h-5 text-yellow-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12A9 9 0 113 12a9 9 0 0118 0z" />
        </svg>
    ),
    error: (
        <svg className="w-5 h-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
        </svg>
    ),
    info: (
        <svg className="w-5 h-5 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 16h-1v-4h-1m1-4h.01" />
        </svg>
    ),
};

const CustomBanner: React.FC<CustomBannerProps> = ({
    type = "info",
    text,
    className,
    ...props
}) => {
    if (!text) {
        return null;
    }
    if (!typeStyles[type]) {
        console.error(`Invalid type "${type}" provided to CustomBanner.`);
        return null;
    }
    return (
        <div
            className={classNames(
                "flex items-center gap-3 border-l-4 p-4 rounded-lg",
                typeStyles[type],
                className
            )}
            {...props}
        >
            <span>{iconMap[type]}</span>
            <span className="font-medium">{text}</span>
        </div>
    );
};

export default CustomBanner;