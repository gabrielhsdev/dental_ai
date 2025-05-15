'use client'
import CustomButton from "@/components/CustomButton";
import CustomInput from "@/components/CustomInput";
import { useSession } from "@/hooks/useSession";
import { useState } from "react";

export default function Home() {
  const { handleLogin } = useSession();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

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
              <div>
                <CustomInput type="email" label="Email" value={email} onChange={(value) => setEmail(value)} placeholder="email" />
              </div>
              <div>
                <CustomInput type="password" label="Senha" value={password} onChange={(value) => setPassword(value)} placeholder="senha" />
              </div>
              <CustomButton text="Login" onClick={() => handleLogin(email, password)} />
              <p className="text-sm font-light text-gray-500 dark:text-gray-400">
                Don't have an account yet? <a href="#" className="font-medium text-primary-600 hover:underline dark:text-primary-500">Sign up</a>
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
