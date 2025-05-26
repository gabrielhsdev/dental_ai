'use client';

import CustomButton from "@/components/CustomButton";
import CustomInput from "@/components/CustomInput";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { useSessionContext } from "@/context/SessionContext";

export default function RegisterPage() {
  const { handleRegister } = useSessionContext();
  const router = useRouter();
  const [form, setForm] = useState({
    firstName: '',
    lastName: '',
    email: '',
    password: '',
    userName: '',
  });

  const handleChange = (key: string, value: string) => {
    setForm(prev => ({ ...prev, [key]: value }));
  };

  const handleRegisterAux = async () => {
    if (!form.userName || !form.email || !form.password || !form.firstName || !form.lastName) {
      return alert("Por favor, preencha todos os campos.");
    }
    await handleRegister(
      form.userName,
      form.email,
      form.password,
      form.firstName,
      form.lastName
    );
  };

  return (
    <div className="flex items-center justify-center h-screen bg-gray-100 dark:bg-gray-900">
      <div className="flex space-x-8 items-center">

        {/* Imagem */}
        <div className="w-105 h-150 rounded-lg overflow-hidden shadow-lg">
          <img
            src="/dentist3.jpg"
            alt="Imagem decorativa"
            className="w-full h-full object-cover"
          />
        </div>

        {/* Formulário de Cadastro */}
        <div className="w-125 h-150 bg-white dark:bg-gray-800 pt-1 px-6 pb-6 py-2 rounded-lg shadow-md flex flex-col">
          <a
            href="#"
            className="flex items-center mb-0.5 text-2xl font-semibold text-gray-900 dark:text-white"
          >
            <img className="w-30 h-30 mr-2" src="/logo.png" alt="logo" />
            Cadastro
          </a>

          <div className="flex flex-col space-y-4">
            <CustomInput
              label="Usuário"
              type="text"
              value={form.userName}
              onChange={(val) => handleChange('userName', val)}
              placeholder="Digite seu usuário"
            />
            <CustomInput
              label="Nome"
              type="text"
              value={form.firstName}
              onChange={(val) => handleChange('firstName', val)}
              placeholder="Digite seu nome"
            />
            <CustomInput
              label="Sobrenome"
              type="text"
              value={form.lastName}
              onChange={(val) => handleChange('lastName', val)}
              placeholder="Digite seu sobrenome"
            />
            <CustomInput
              label="Email"
              type="email"
              value={form.email}
              onChange={(val) => handleChange('email', val)}
              placeholder="Digite seu email"
            />
            <CustomInput
              label="Senha"
              type="password"
              value={form.password}
              onChange={(val) => handleChange('password', val)}
              placeholder="Digite sua senha"
            />
            <CustomButton
              text="Cadastrar"
              onClick={handleRegisterAux}
              className="mt-4"
            />
            <p className="text-sm font-light text-gray-500 dark:text-gray-400">
              Já tem uma conta?{" "}
              <button
                onClick={() => router.push("/")}
                className="font-medium text-primary-600 hover:underline dark:text-primary-500"
              >
                Fazer Login
              </button>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
