'use client';

import CustomCard from "@/components/CustomCard";

export default function Dashboard() {
  return (
    <div className="flex flex-col items-center w-full min-h-screen px-6 py-10 bg-gray-50 dark:bg-gray-900">
      
      {/* Card de instruções */}
      <div className="w-full max-w-4xl">
        <CustomCard
          title="Bem-vindo ao IACare!"
          subtitle="Como usar nossa solução:"
        >
          <ul className="list-disc list-inside space-y-2 text-gray-800 dark:text-gray-200">
            <li>
              1. Cadastre seus pacientes usando o botão <strong>"Novo Paciente"</strong> no menu lateral.
            </li>
            <li>
              2. Acesse o paciente e clique em <strong>"Novo Atendimento"</strong> para criar um novo atendimento.
            </li>
            <li>
              3. Faça upload dos exames do paciente clicando em <strong>"Novo Exame"</strong>.
            </li>
            <li>
              4. Veja o resultado do exame ou acesse exames anteriores no histórico do paciente.
            </li>
          </ul>
        </CustomCard>
      </div>

      {/* Imagem ilustrativa abaixo */}
      <div className="mt-10 w-full max-w-6xl">
        <img
          src="/dentist4.jpg"
          alt="Ilustração do dashboard"
          className="w-full h-64 object-cover rounded-lg shadow-md"
        />
      </div>
    </div>
  );
}
