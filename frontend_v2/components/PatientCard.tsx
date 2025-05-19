import { PatientInterface } from "@/common/commonInterfaces";
import CustomButton from "./CustomButton";

interface Props {
    patient: PatientInterface;
    onView: (patient: PatientInterface) => void;
}

export default function PatientCard({ patient, onView }: Props) {
    return (
        <div className="w-full p-4 bg-white border border-gray-200 rounded-lg shadow-sm dark:bg-gray-800 dark:border-gray-700 flex flex-col justify-between gap-4">
            <div>
                <p className="text-lg font-semibold text-gray-900 dark:text-white">
                    {patient.firstName} {patient.lastName}
                </p>
                <p className="text-sm text-gray-500 dark:text-gray-400">
                    Email: {patient.email}
                </p>
                <p className="text-sm text-gray-500 dark:text-gray-400">
                    Telefone: {patient.phoneNumber}
                </p>
            </div>

            <div className="flex justify-end">
                <CustomButton
                    text="Ver Paciente"
                    className="text-sm px-4 py-2"
                    onClick={() => onView(patient)}
                />
            </div>
        </div>
    );
}
