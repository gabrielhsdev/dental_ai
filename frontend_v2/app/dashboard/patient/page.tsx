'use client';

import CustomCard from "@/components/CustomCard";
import { useSessionContext } from "@/context/SessionContext";
import { useSelectedPatientContext } from "@/context/SelectedPatientContext";
import CustomButton from "@/components/CustomButton";
import { usePatientImages } from "@/hooks/usePatientImages";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

export default function ListPatients() {
    const router = useRouter();
    const { getToken } = useSessionContext();
    const { selectedPatient } = useSelectedPatientContext();
    const { fetchImagesByPatientId, images, getImageByPath } = usePatientImages();

    const [imageURLs, setImageURLs] = useState<Record<string, string>>({});

    useEffect(() => {
        const fetchPatientImages = async () => {
            const token = await getToken();
            if (selectedPatient?.id && token) {
                await fetchImagesByPatientId(selectedPatient.id, token);

                // Get actual image blobs for previews
                const urls: Record<string, string> = {};
                for (const image of images) {
                    const blob = await getImageByPath(image.imageData, token);
                    const objectURL = URL.createObjectURL(blob);
                    urls[image.id] = objectURL;
                }
                setImageURLs(urls);
            }
        };
        fetchPatientImages();
    }, []);

    return (
        <>
            <CustomCard title="Listar Pacientes" subtitle="Lista de pacientes cadastrados" className="grid-cols-12 gap-4">
                <CustomButton
                    text="Debugar Paciente"
                    onClick={() => console.log("Selected Patient:", selectedPatient)}
                    className="col-span-12"
                />
            </CustomCard>

            <CustomCard title="Novo Diagnóstico" subtitle="Novo diagnóstico para o paciente" className="grid-cols-12 gap-4">
                <CustomButton
                    text="Novo Diagnóstico"
                    onClick={() => router.push(`patient/newDiagnostic`)}
                    className="col-span-12"
                />
            </CustomCard>

            <CustomCard title="Histórico de Exames" subtitle="Lista de exames cadastrados" className="grid-cols-12 gap-4">
                {images.map((image) => (
                    <div key={image.id} className="col-span-12">
                        <p>{image.description}</p>
                        <p>{image.imageData}</p>
                        <img
                            src={imageURLs[image.id]}
                            alt={image.description}
                            className="w-32 h-32 object-cover"
                        />
                    </div>
                ))}
            </CustomCard>
        </>
    );
}
