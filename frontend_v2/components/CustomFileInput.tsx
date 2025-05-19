'use client';

import React from 'react';

interface CustomFileInputProps {
    label?: string;
    onChange: (file: File | null) => void;
    accept?: string;
    className?: string;
}

const CustomFileInput = ({
    label,
    onChange = () => console.warn('onChange not implemented'),
    accept = '',
    className = '',
}: CustomFileInputProps) => {
    return (
        <div className={`flex flex-col gap-2 ${className}`}>
            {label && (
                <label className="text-sm font-medium text-gray-900 dark:text-white">
                    {label}
                </label>
            )}
            <input
                type="file"
                accept={accept}
                onChange={(e) => onChange(e.target.files?.[0] || null)}
                className="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400"
            />
        </div>
    );
};

export default CustomFileInput;
