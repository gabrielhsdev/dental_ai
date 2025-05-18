'use client';
import React, { ReactNode, useState, useEffect } from 'react';

export default function RootLayout({ children }: { children: ReactNode }) {
    const [sidebarOpen, setSidebarOpen] = useState(false);
    const [isSmallScreen, setIsSmallScreen] = useState(false);

    useEffect(() => {
        const mediaQuery = window.matchMedia('(max-width: 640px)');
        setIsSmallScreen(mediaQuery.matches);

        const handler = (e: MediaQueryListEvent) => setIsSmallScreen(e.matches);
        mediaQuery.addEventListener('change', handler);

        // Close sidebar if small screen
        if (mediaQuery.matches) setSidebarOpen(false);

        return () => mediaQuery.removeEventListener('change', handler);
    }, []);

    useEffect(() => {
        if (!isSmallScreen) {
            setSidebarOpen(true);
        }
    }, [isSmallScreen]);

    const toggleSidebar = () => setSidebarOpen((open) => !open);

    return (
        <div className="flex">
            <aside
                id="default-sidebar"
                className={`fixed top-0 left-0 z-40 w-64 h-screen transition-transform bg-gray-50 dark:bg-gray-800
          ${sidebarOpen ? 'translate-x-0' : '-translate-x-full'}
          sm:translate-x-0
        `}
                aria-label="Sidebar"
                aria-hidden={!sidebarOpen && isSmallScreen}
            >
                <div className="h-full px-3 py-4 overflow-y-auto flex flex-col justify-between">
                    {/* Sidebar content */}
                    <ul className="space-y-2 font-medium">
                        <li>
                            <a
                                href="#"
                                className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
                            >
                                <span className="ms-3">Dashboard</span>
                            </a>
                        </li>
                        <li>
                            <a
                                href="#"
                                className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
                            >
                                <span className="flex-1 ms-3 whitespace-nowrap">Kanban</span>
                                <span className="inline-flex items-center justify-center px-2 ms-3 text-sm font-medium text-gray-800 bg-gray-100 rounded-full dark:bg-gray-700 dark:text-gray-300">
                                    Pro
                                </span>
                            </a>
                        </li>
                    </ul>

                    {/* Toggle button inside sidebar, at bottom */}
                    {isSmallScreen && (
                        <button
                            onClick={toggleSidebar}
                            aria-controls="default-sidebar"
                            aria-expanded={sidebarOpen}
                            type="button"
                            className="w-full flex items-center justify-center p-2 mt-4 mb-2 text-sm text-gray-500 rounded-lg hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
                        >
                            <span className="sr-only">{sidebarOpen ? 'Close sidebar' : 'Open sidebar'}</span>
                            <svg
                                className="w-6 h-6"
                                aria-hidden="true"
                                fill="currentColor"
                                viewBox="0 0 20 20"
                                xmlns="http://www.w3.org/2000/svg"
                            >
                                {/* Close icon */}
                                {sidebarOpen ? (
                                    <path
                                        fillRule="evenodd"
                                        clipRule="evenodd"
                                        d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                                    />
                                ) : (
                                    // Hamburger icon
                                    <path d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z" />
                                )}
                            </svg>
                        </button>
                    )}
                </div>
            </aside>

            {isSmallScreen && !sidebarOpen && (
                <button
                    onClick={toggleSidebar}
                    aria-controls="default-sidebar"
                    aria-expanded={sidebarOpen}
                    type="button"
                    className="fixed top-4 left-4 z-50 inline-flex items-center p-2 text-sm text-gray-500 rounded-lg bg-white shadow hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
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
            )}

            {/* Main content */}
            <main className={`flex-1 p-4 transition-margin duration-300 ${sidebarOpen && !isSmallScreen ? 'sm:ml-64' : ''}`}>
                {children}
            </main>
        </div>
    );
}
