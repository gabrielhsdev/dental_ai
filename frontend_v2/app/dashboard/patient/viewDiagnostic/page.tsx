'use client';

import CustomCard from "@/components/CustomCard";
import { useSessionContext } from "@/context/SessionContext";
import { useSelectedPatientContext } from "@/context/SelectedPatientContext";
import { usePatientImages } from "@/hooks/usePatientImages";

export default function ViewDiagnostic() {
    const { getToken } = useSessionContext();
    const { selectedPatient } = useSelectedPatientContext();
    const { fetchImagesByPatientId, images } = usePatientImages();

    return (
        <>
            <CustomCard
                title="Listar Pacientes"
                subtitle="Lista de pacientes cadastrados"
                className="grid-cols-12 gap-4"
            >
                Ver Diagnostico
            </CustomCard>
        </>
    );
}
