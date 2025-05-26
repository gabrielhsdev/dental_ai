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
    const [expandedImages, setExpandedImages] = useState<Record<string, boolean>>({});

    const fetchPatientImages = async () => {
        const token = await getToken();
        if (selectedPatient?.id && token) {
            await fetchImagesByPatientId(selectedPatient.id, token);
        }
    };

    const handleVisualizarClick = async (imagePath: string) => {
        const token = await getToken();
        if (!imageBlobUrls[imagePath] && token) {
            const image = await getImageByPath(imagePath, token);
            const objectUrl = URL.createObjectURL(image);
            setImageBlobUrls((prev) => ({
                ...prev,
                [imagePath]: objectUrl,
            }));
        }
        setExpandedImages((prev) => ({
            ...prev,
            [imagePath]: !prev[imagePath],
        }));
    };

    useEffect(() => {
        fetchPatientImages();
    }, []);

    return (
        <>
            <CustomCard title="Histórico de Exames" subtitle="Lista de exames cadastrados" className="grid-cols-12 gap-4">
                {images && images.length > 0 && images.map((image) => (
                    <div key={image.id} className="col-span-12 border p-4 rounded-xl mb-4 shadow-sm">
                        <p className="font-medium">{image.description}</p>
                        <p className="text-sm text-gray-500">{image.imageData}</p>

                        <CustomButton
                            text={expandedImages[image.imageData] ? "Ocultar Exame" : "Visualizar Exame"}
                            onClick={() => handleVisualizarClick(image.imageData)}
                            className="mt-2"
                        />

                        {expandedImages[image.imageData] && imageBlobUrls[image.imageData] && (
                            <div className="mt-4">
                                <img
                                    src={imageBlobUrls[image.imageData]}
                                    alt={image.description}
                                    className="w-full h-auto max-w-full"
                                />
                            </div>
                        )}
                    </div>
                ))}
            </CustomCard>
        </>
    );
}
