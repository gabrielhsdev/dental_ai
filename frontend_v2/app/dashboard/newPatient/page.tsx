'use client';

import CustomBanner from "@/components/CustomBanner";
import CustomButton from "@/components/CustomButton";
import CustomCard from "@/components/CustomCard";
import CustomInput from "@/components/CustomInput";
import { useSessionContext } from "@/context/SessionContext";
import { usePatients } from "@/hooks/usePatients";
import { toRFC3339 } from "@/utils/dateUtils";
import { useState } from "react";
import Image from "next/image";

export default function NewPatient() {
    const { session, getToken } = useSessionContext();
    const { isLoading, error, createPatient } = usePatients();
    const [formState, setFormState] = useState({
        firstName: '',
        lastName: '',
        dateOfBirth: '',
        gender: '',
        phoneNumber: '',
        email: '',
        notes: '',
    });

    const handleCreatePatient = async () => {
        if (isLoading) return;

        const token = await getToken();
        if (!token) return;

        console.log('Valores enviados:', {
            ...formState,
            dateOfBirth: toRFC3339(formState.dateOfBirth),
            token
        });

        try {
            await createPatient(
                formState.firstName,
                formState.lastName,
                toRFC3339(formState.dateOfBirth),
                formState.gender,
                formState.phoneNumber,
                formState.email,
                formState.notes,
                token
            );
        } catch (error) {
            console.error('Erro ao criar paciente:', error);
        }
    };

    return (
        <>
            {/* Cabeçalho com título e logo */}
            <div className="flex justify-between items-center mb-6">
                <div>
                    <h2 className="text-2xl font-bold">Adicionar Novo Paciente</h2>
                    <p className="text-gray-500">Preencha os dados do paciente abaixo</p>
                </div>
                <Image
                    src="/logo.png"
                    alt="Logo"
                    width={100}
                    height={110}
                />
            </div>

            {/* Formulário */}
            <CustomCard className="grid-cols-12 gap-4">
                <CustomBanner
                    type="error"
                    text={error ? error : ''}
                    className="col-span-12"
                />
                <CustomInput
                    label="Nome"
                    type="text"
                    placeholder="Nome"
                    value={formState.firstName}
                    className="col-span-12"
                    onChange={(value: string) => setFormState({ ...formState, firstName: value })}
                />
                <CustomInput
                    label="Sobrenome"
                    type="text"
                    placeholder="Sobrenome"
                    value={formState.lastName}
                    className="col-span-6"
                    onChange={(value: string) => setFormState({ ...formState, lastName: value })}
                />
                <CustomInput
                    label="Data de Nascimento"
                    type="date"
                    placeholder="Data de Nascimento"
                    value={formState.dateOfBirth}
                    className="col-span-6"
                    onChange={(value: string) => setFormState({ ...formState, dateOfBirth: value })}
                />
                <CustomInput
                    label="Gênero"
                    type="text"
                    placeholder="Gênero"
                    value={formState.gender}
                    className="col-span-6"
                    onChange={(value: string) => setFormState({ ...formState, gender: value })}
                />
                <CustomInput
                    label="Número de Telefone"
                    type="text"
                    placeholder="Número de Telefone"
                    value={formState.phoneNumber}
                    className="col-span-6"
                    onChange={(value: string) => setFormState({ ...formState, phoneNumber: value })}
                />
                <CustomInput
                    label="Email"
                    type="email"
                    placeholder="Email"
                    value={formState.email}
                    className="col-span-6"
                    onChange={(value: string) => setFormState({ ...formState, email: value })}
                />
                <CustomInput
                    label="Notas"
                    type="text"
                    placeholder="Notas"
                    value={formState.notes}
                    className="col-span-6"
                    onChange={(value: string) => setFormState({ ...formState, notes: value })}
                />
                <CustomButton
                    text="Adicionar Paciente"
                    className="col-span-12"
                    onClick={handleCreatePatient}
                    disabled={isLoading}
                />
            </CustomCard>
        </>
    );
}