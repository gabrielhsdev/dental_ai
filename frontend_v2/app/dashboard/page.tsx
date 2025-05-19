'use client';

import CustomCard from "@/components/CustomCard";
import { useSessionContext } from "@/context/SessionContext";

export default function Dashboard() {
    const { session } = useSessionContext();
    return (
        <>
            <CustomCard
                title="Dashboard"
                subtitle="Click the button to debug your session"
            >
                <p className="text-gray-700 dark:text-gray-300">
                    This is your dashboard. You can use the buttons above to manage your workflow.
                </p>
                <button
                    onClick={() => { console.log(session); }}
                    className="px-4 py-2 text-white bg-blue-600 rounded hover:bg-blue-700"
                >
                    Print Session
                </button>
            </CustomCard>
            <CustomCard
                title="Bem vindo ao Dashboard"
                subtitle="Como usar nossa solução:"
            >
                <ul className="list-disc list-inside space-y-2">
                    <li>
                        1. Cadastre seus pacientes, utilize o botão <strong>"Novo Paciente"</strong> no menu lateral.
                    </li>
                    <li>
                        2. Acesse o seu paciente e clique no botão <strong>"Novo Atendimento"</strong> para criar um novo atendimento.
                    </li>
                    <li>
                        3. Faça o upload dos exames do paciente, clique no botão <strong>"Novo Exame"</strong> para fazer o upload de um novo exame.
                    </li>
                    <li>
                        4. Veja o resultado do exame, você também pode conferir resultados passados ao acessar um paciente.
                    </li>
                </ul>
            </CustomCard>
        </>
    );
}
