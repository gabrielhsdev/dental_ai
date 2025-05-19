import React from "react";

interface SidebarItemProps {
    id: string | number;
    href: string;
    label: string;
    isActive?: boolean;
    badge?: React.ReactNode;
}

const SidebarItem: React.FC<SidebarItemProps> = ({
    id,
    href,
    label,
    isActive = false,
    badge,
}) => (
    <li key={id}>
        <a
            href={href}
            className={`flex items-center p-2 rounded-lg group ${isActive
                    ? "bg-blue-500 text-white dark:bg-blue-600"
                    : "text-gray-900 hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700"
                }`}
            aria-current={isActive ? "page" : undefined}
        >
            <span className="flex-1 ms-3 whitespace-nowrap">{label}</span>
            {badge && (
                <span className="inline-flex items-center justify-center px-2 ms-3 text-sm font-medium text-gray-800 bg-gray-100 rounded-full dark:bg-gray-700 dark:text-gray-300">
                    {badge}
                </span>
            )}
        </a>
    </li>
);

export default SidebarItem;