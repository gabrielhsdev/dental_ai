'use client';
import CustomButton from "@/components/CustomButton";
import CustomInput from "@/components/CustomInput";
import { useSessionContext } from "@/context/SessionContext";
import { useState } from "react";

export default function Home() {
  const { handleLogin } = useSessionContext();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  return (
    <div className="flex flex-col items-center justify-center w-full sm:max-w-md bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
      <a href="#" className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
        <img className="w-8 h-8 mr-2" src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/logo.svg" alt="logo" />
        Login
      </a>
      <div className="w-full space-y-4">
        <CustomInput type="email" label="Email" value={email} onChange={setEmail} placeholder="email" />
        <CustomInput type="password" label="Senha" value={password} onChange={setPassword} placeholder="senha" />
        <CustomButton text="Login" onClick={() => handleLogin(email, password)} />
        <p className="text-sm font-light text-gray-500 dark:text-gray-400">
          Don't have an account yet? <a href="#" className="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign up</a>
        </p>
      </div>
    </div>
  );
}
