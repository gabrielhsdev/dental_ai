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
    const [imageBlobUrls, setImageBlobUrls] = useState<Record<string, string>>({});


    const fetchPatientImages = async () => {
        const token = await getToken();
        if (selectedPatient?.id && token) {
            await fetchImagesByPatientId(selectedPatient.id, token);
            console.log("Fetched images:", images);
        }
    };

    const debugOnImageClick = async (imagePath: string) => {
        const token = await getToken();
        if (token) {
            const image = await getImageByPath(imagePath, token);

            console.log("typeof image:", typeof image);
            console.log("instanceof Blob:", image instanceof Blob);
            console.log("Raw image:", image);

            try {
                const objectUrl = URL.createObjectURL(image);
                console.log("Image Blob URL:", objectUrl);

                // Save the blob URL in state
                setImageBlobUrls((prev) => ({
                    ...prev,
                    [imagePath]: objectUrl,
                }));
            } catch (err) {
                console.error("Failed to create object URL:", err);
            }
        }
    };

    useEffect(() => {
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
                            src={imageBlobUrls[image.imageData]} // ✅ dynamically set if available
                            onClick={() => debugOnImageClick(image.imageData)} // clicking fetches it
                            alt={image.description}
                            className="w-32 h-32 object-cover border cursor-pointer"
                        />

                    </div>
                ))}
            </CustomCard>
        </>
    );
}
