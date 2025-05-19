'use client';

import CustomCard from "@/components/CustomCard";
import CustomFileInput from "@/components/CustomFileInput";
import { useSessionContext } from "@/context/SessionContext";
import { useSelectedPatientContext } from "@/context/SelectedPatientContext";
import { usePatientImages } from "@/hooks/usePatientImages";
import { useState } from "react";
import CustomInput from "@/components/CustomInput";

export default function NewDiagnostic() {
    const { getToken } = useSessionContext();
    const { selectedPatient } = useSelectedPatientContext();
    const { processImage, createPatientImage } = usePatientImages();
    const [file, setFile] = useState<File | null>(null);
    const [description, setDescription] = useState<string>("");

    const handleUpload = async () => {
        try {
            const token = await getToken();
            if (selectedPatient?.id && token && file) {
                const resultImage = await processImage(file, token);
                if (resultImage) {
                    await createPatientImage(
                        selectedPatient.id,
                        resultImage.result_image,
                        'png',
                        description,
                        token
                    );
                } else {
                    throw new Error("Image processing failed");
                }
            }
        } catch (error) {
            console.error("Error uploading file:", error);
        }
    };

    return (
        <CustomCard
            title="Novo Diagnóstico"
            subtitle="Novo diagnóstico para o paciente"
            className="grid-cols-12 gap-4"
        >
            <div className="col-span-12">
                <CustomFileInput
                    label="Selecione uma imagem"
                    accept="image/*"
                    onChange={setFile}
                />
            </div>

            <div className="col-span-12">
                <CustomInput
                    label="Descrição"
                    value={description}
                    onChange={setDescription}
                    placeholder="Descrição da imagem"
                    className="w-full"
                />
            </div>

            <div className="col-span-12">
                <button
                    onClick={handleUpload}
                    disabled={!file || !description}
                    className="mt-4 px-4 py-2 bg-blue-600 text-white rounded-lg disabled:bg-gray-400 w-full"
                >
                    Enviar Imagem
                </button>
            </div>
        </CustomCard>
    );
}
