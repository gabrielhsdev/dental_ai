'use client';

import { useSessionContext } from "@/context/SessionContext";

export default function Dashboard() {
    const { session } = useSessionContext();

    function printSession() {
        console.log(session);
    }

    return (
        <div className="grid w-full max-w-md gap-6 p-6 bg-white rounded-lg shadow dark:bg-gray-800 dark:border dark:border-gray-700">
            {/* Logo / Header */}
            <a href="#" className="flex items-center text-2xl font-semibold text-gray-900 dark:text-white">
                <img
                    className="w-8 h-8 mr-2"
                    src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/logo.svg"
                    alt="logo"
                />
                Dashboard
            </a>

            {/* Body */}
            <div className="space-y-4">
                <p className="text-gray-700 dark:text-gray-300">
                    Click the button below to print the session data:
                </p>

                <button
                    onClick={printSession}
                    className="px-4 py-2 text-white bg-blue-600 rounded hover:bg-blue-700"
                >
                    Print Session
                </button>
            </div>
        </div>
    );
}
