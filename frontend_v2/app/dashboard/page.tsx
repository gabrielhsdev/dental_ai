'use client'

import { useSessionContext } from "@/context/SessionContext";

export default function Dashboard() {
    const { session } = useSessionContext();

    function printsession() {
        console.log(session);
    }

    return (
        <section className="bg-gray-50 dark:bg-gray-900">
            <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
                <a href="#" className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
                    <img className="w-8 h-8 mr-2" src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/logo.svg" alt="logo" />
                    Login
                </a>
                <div className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
                    <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                        <div className="space-y-4 md:space-y-6">
                            Click on the button to print the session data:
                            <button
                                className="px-4 py-2 text-white bg-blue-500 rounded hover:bg-blue-600"
                                onClick={printsession}
                            >
                                Print Session
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    );
}
