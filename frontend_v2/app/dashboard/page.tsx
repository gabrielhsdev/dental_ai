'use client';

import CustomCard from "@/components/CustomCard";
import { usePatientImages } from "@/hooks/usePatientImages";
import { useEffect, useState } from "react";
import { useSessionContext } from "@/context/SessionContext";

export default function Dashboard() {
  const { session, getToken } = useSessionContext();
  const { fetchAllImagesByUser, images } = usePatientImages();
  const [dataForGraphs, setDataForGraphs] = useState<any[]>([]);
  const [teethForGraphs, setTeethForGraphs] = useState<any[]>([]);

  useEffect(() => {
    const userId = session.user?.id;
    const token = getToken();
    if (token && userId) {
      fetchAllImagesByUser(userId, token);
    }
  }, []);

  useEffect(() => {
    if (images.length > 0) {
      const rawGraphData = images.map((image) => {
        const parsedInference = typeof image.inferenceData === "string"
          ? JSON.parse(image.inferenceData)
          : image.inferenceData;

        const detections = parsedInference?.predictions?.detections ?? [];

        const cariesCount = detections.filter(
          (detection) => detection.class_name === "Caries"
        ).length;

        // Normalize date to 'YYYY-MM-DD'
        const date = new Date(image.uploadedAt).toISOString().split('T')[0];

        return { date, cariesCount };
      });

      // Aggregate by date
      const aggregatedGraphData = Object.values(
        rawGraphData.reduce((acc, { date, cariesCount }) => {
          if (!acc[date]) {
            acc[date] = { date, cariesCount };
          } else {
            acc[date].cariesCount += cariesCount;
          }
          return acc;
        }, {} as Record<string, { date: string; cariesCount: number }>)
      );

      setDataForGraphs(aggregatedGraphData);

      // Now for Teeth
      const teethData = images.map((image) => {
        const parsedInference = typeof image.inferenceData === "string"
          ? JSON.parse(image.inferenceData)
          : image.inferenceData;

        const detections = parsedInference?.predictions?.detections ?? [];

        const teethCount = detections.filter(
          (detection) => detection.class_name === "Tooth"
        ).length;
        const date = new Date(image.uploadedAt).toISOString().split('T')[0];
        return { date, teethCount };
      });

      // Aggregate teeth data by date
      const aggregatedTeethData = Object.values(
        teethData.reduce((acc, { date, teethCount }) => {
          if (!acc[date]) {
            acc[date] = { date, teethCount };
          } else {
            acc[date].teethCount += teethCount;
          }
          return acc;
        }, {} as Record<string, { date: string; teethCount: number }>)
      );
      setTeethForGraphs(aggregatedTeethData);
    } else {
      setDataForGraphs([]);
    }
  }, [images]);

  useEffect(() => {
    console.log("Graph data updated:", dataForGraphs);
    console.log("Teeth data updated:", teethForGraphs);
  }, [dataForGraphs, teethForGraphs]);

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

      {/* lets make a table to showcase the amount of caries per date */}
      <div className="mt-10 w-full max-w-4xl">
        <CustomCard
          title="Quantidade de Cáries por Data"
          subtitle="Visualize a evolução dos atendimentos"
        >
          {dataForGraphs.length > 0 ? (
            <table className="min-w-full bg-white dark:bg-gray-800 rounded-lg shadow-md">
              <thead>
                <tr className="border-b">
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Data</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Quantidade de Cáries</th>
                </tr>
              </thead>
              <tbody>
                {dataForGraphs.map((data, index) => (
                  <tr key={index} className="border-b hover:bg-gray-50 dark:hover:bg-gray-700">
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">{new Date(data.date).toLocaleDateString()}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">{data.cariesCount}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          ) : (
            <p className="text-gray-600 dark:text-gray-400">Nenhum dado disponível para exibir.</p>
          )}
        </CustomCard>
      </div>

      {/* Card para dentes */}
      <div className="mt-10 w-full max-w-4xl">
        <CustomCard
          title="Quantidade de Dentes por Data"
          subtitle="Visualize a evolução dos atendimentos"
        >
          {teethForGraphs.length > 0 ? (
            <table className="min-w-full bg-white dark:bg-gray-800 rounded-lg shadow-md">
              <thead>
                <tr className="border-b">
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Data</th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-400 uppercase tracking-wider">Quantidade de Dentes</th>
                </tr>
              </thead>
              <tbody>
                {teethForGraphs.map((data, index) => (
                  <tr key={index} className="border-b hover:bg-gray-50 dark:hover:bg-gray-700">
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">{new Date(data.date).toLocaleDateString()}</td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900 dark:text-gray-100">{data.teethCount}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          ) : (
            <p className="text-gray-600 dark:text-gray-400">Nenhum dado disponível para exibir.</p>
          )}
        </CustomCard>
      </div>
    </div>
  );
}
