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
        <CustomCard
            title="Listar Pacientes"
            subtitle="Lista de pacientes cadastrados"
            className="grid-cols-12 gap-4"
        >
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
    );
}
