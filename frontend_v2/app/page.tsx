'use client';

import CustomButton from "@/components/CustomButton";
import CustomInput from "@/components/CustomInput";
import { useSessionContext } from "@/context/SessionContext";
import { useState } from "react";
import { useRouter } from "next/navigation";

export default function LoginPage() {
  const { handleLogin } = useSessionContext();
  const router = useRouter();

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  return (
    <div className="flex items-center justify-center h-screen bg-gray-100 dark:bg-gray-900">
      <div className="flex space-x-8 items-center">

        {/* Imagem */}
        <div className="w-106 h-112 rounded-lg overflow-hidden shadow-lg">
          <img
            src="/dentist1.jpg"
            alt="Imagem decorativa"
            className="w-full h-full object-cover"
          />
        </div>

        {/* Formulário de Login */}
        <div className="w-80 bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md">
          <a href="#" className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
            <img className="w-25 h-30 mr-2" src="/logo.png" alt="logo" />
            Login
          </a>
          <div className="space-y-4">
            <CustomInput
              type="email"
              label="Email"
              value={email}
              onChange={setEmail}
              placeholder="Insira seu email"
            />
            <CustomInput
              type="password"
              label="Senha"
              value={password}
              onChange={setPassword}
              placeholder="Insira sua senha"
            />
            <CustomButton text="Fazer Login" onClick={() => handleLogin(email, password)} />
            <p className="text-sm font-light text-gray-500 dark:text-gray-400">
              Não tem uma conta ainda?{" "}
              <button
                onClick={() => router.push("/registerlogin")}
                className="font-medium text-primary-600 hover:underline dark:text-primary-500"
              >
                Cadastre-se
              </button>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
