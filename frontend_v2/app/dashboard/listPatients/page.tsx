'use client';

import CustomBanner from "@/components/CustomBanner";
import CustomCard from "@/components/CustomCard";
import { useSessionContext } from "@/context/SessionContext";
import { usePatients } from "@/hooks/usePatients";
import { useEffect } from "react";
import { ClipLoader } from "react-spinners";
import PatientCard from "@/components/PatientCard";
import { useSelectedPatientContext } from "@/context/SelectedPatientContext";
import { PatientInterface } from "@/common/commonInterfaces";
import Image from "next/image";

export default function ListPatients() {
    const { getToken } = useSessionContext();
    const { selectPatient } = useSelectedPatientContext();
    const { isLoading, error, fetchPatientsByUserId, patients } = usePatients();

    const handleFetchPatients = async () => {
        try {
            if (isLoading) return;
            const token = await getToken();
            if (!token) return;
            await fetchPatientsByUserId(token);
        } catch (error) {
            console.error('Error fetching patients:', error);
        }
    };

    const handleSelectPatient = (patient: PatientInterface) => {
        selectPatient(patient);
    }

    useEffect(() => {
        handleFetchPatients();
    }, []);

    return (
        <>
            {/* Cabeçalho com título e logo */}
            <div className="flex justify-between items-center mb-6">
                <div>
                    <h2 className="text-2xl font-bold">Listar Pacientes</h2>
                    <p className="text-gray-500">Lista de pacientes cadastrados</p>
                </div>
                <Image
                    src="/logo.png"
                    alt="Logo"
                    width={100}
                    height={110}
                />
            </div>

            {/* Conteúdo principal */}
            <CustomCard className="grid-cols-12 gap-4">
                {isLoading ? (
                    <div className="col-span-12 flex justify-center py-8">
                        <ClipLoader size={40} color="#3B82F6" />
                    </div>
                ) : (
                    <>
                        <CustomBanner
                            type="error"
                            text={error || ''}
                            className="col-span-12"
                        />

                        {patients && patients.length > 0 ? (
                            <div className="col-span-12 grid grid-cols-1 gap-4">
                                {patients.map((patient) => (
                                    <PatientCard
                                        key={patient.id}
                                        patient={patient}
                                        onView={() => handleSelectPatient(patient)}
                                    />
                                ))}
                            </div>
                        ) : (
                            <div className="col-span-12 py-8 text-center">
                                <p className="text-gray-500">Nenhum paciente encontrado.</p>
                            </div>
                        )}
                    </>
                )}
            </CustomCard>
        </>
    );
}