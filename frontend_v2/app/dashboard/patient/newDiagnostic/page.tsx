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
    const [loading, setLoading] = useState<boolean>(false);

    const handleUpload = async () => {
        try {
            setLoading(true);
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
        } finally {
            setLoading(false);
            setDescription("");
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
                    disabled={!file || !description || loading}
                    className="mt-4 px-4 py-2 bg-blue-600 text-white rounded-lg disabled:bg-gray-400 w-full flex items-center justify-center"
                >
                    {loading ? (
                        <span className="flex items-center">
                            <svg className="animate-spin h-5 w-5 mr-2 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                                <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"></path>
                            </svg>
                            Enviando...
                        </span>
                    ) : (
                        "Enviar Imagem"
                    )}
                </button>
            </div>
        </CustomCard>
    );
}
