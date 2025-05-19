import React from "react";

interface SidebarCloseButtonProps {
    sidebarOpen: boolean;
    toggleSidebar: () => void;
}

const SidebarCloseButton: React.FC<SidebarCloseButtonProps> = ({
    sidebarOpen,
    toggleSidebar,
}) => (
    <button
        onClick={toggleSidebar}
        aria-controls="default-sidebar"
        aria-expanded={sidebarOpen}
        type="button"
        className="w-full flex items-center justify-center p-2 mt-4 mb-2 text-sm text-gray-500 rounded-lg hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
    >
        <span className="sr-only">
            {sidebarOpen ? "Close sidebar" : "Open sidebar"}
        </span>
        <svg
            className="w-6 h-6"
            aria-hidden="true"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
        >
            {sidebarOpen ? (
                <path
                    fillRule="evenodd"
                    clipRule="evenodd"
                    d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                />
            ) : (
                <path d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z" />
            )}
        </svg>
    </button>
);

export default SidebarCloseButton;