import React from "react";

interface SidebarFloatingButtonProps {
    isSmallScreen: boolean;
    sidebarOpen: boolean;
    toggleSidebar: () => void;
}

const SidebarFloatingButton: React.FC<SidebarFloatingButtonProps> = ({
    isSmallScreen,
    sidebarOpen,
    toggleSidebar,
}) => {
    if (!isSmallScreen || sidebarOpen) return null;

    return (
        <button
            onClick={toggleSidebar}
            aria-controls="default-sidebar"
            aria-expanded={sidebarOpen}
            type="button"
            className="fixed top-2 left-2 z-50 inline-flex items-center p-2 text-sm text-gray-500 rounded-lg bg-white shadow hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
        >
            <span className="sr-only">Open sidebar</span>
            <svg
                className="w-6 h-6"
                aria-hidden="true"
                fill="currentColor"
                viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg"
            >
                <path d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z" />
            </svg>
        </button>
    );
};

export default SidebarFloatingButton;