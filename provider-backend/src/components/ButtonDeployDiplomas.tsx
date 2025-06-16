import axios from 'axios';
import { useState } from 'react';
import api from '../services/api';
import { StudentsInfo } from '../Types';

interface ButtonDeployDiplomasProps {
    jsonWithStudents: StudentsInfo[];
    setJsonWithStudents: React.Dispatch<React.SetStateAction<StudentsInfo[]>>;
}

interface DiplomaResponse {
    tokenID: number;
    student: string;
}

function ButtonDeployDiplomas({ jsonWithStudents, setJsonWithStudents }: ButtonDeployDiplomasProps) {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const handleBatchDiplomasDeploy = async () => {
        setLoading(true);
        setError(null);

        try {
            // Generate the zip file on the server
            await axios.get('http://localhost:8080/generate-zip');

            // Download the saved zip file
            const response = await axios.get('http://localhost:8080/download-zip', {
                responseType: 'blob', // the response should be a binary file
            });

            // Check if the response is a Blob
            if (!(response.data instanceof Blob)) {
                throw new Error('Response is not a valid Blob');
            }

            // Create a File object from the Blob
            const zipFile = new File([response.data], 'diplomas.zip', { type: 'application/zip' });
            console.log('Zip file details:', {
                name: zipFile.name,
                type: zipFile.type,
                size: zipFile.size
            });

            // Send the zip file to the API
            const formData = new FormData();
            formData.append('file', zipFile);

            const apiResponse = await api.post<DiplomaResponse[]>('/diplomas', formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                }
            });
            console.log('API response:', apiResponse.data);

            // Get the list of students whose diplomas were issued
            const issuedStudentIds = apiResponse.data.map(diploma => diploma.student);

            // Save issued diplomas to localStorage
            const existingIssuedDiplomas = JSON.parse(localStorage.getItem('issuedDiplomas') || '[]') as string[];
            const updatedIssuedDiplomas = Array.from(new Set([...existingIssuedDiplomas, ...issuedStudentIds]));
            localStorage.setItem('issuedDiplomas', JSON.stringify(updatedIssuedDiplomas));

            // Update the students list to mark only the issued diplomas
            const updatedStudents = jsonWithStudents.map(student => ({
                ...student,
                diplomaIssued: issuedStudentIds.includes(student.studentId)
            }));
            setJsonWithStudents(updatedStudents);

            alert('Diplomas successfully deployed');
        } catch (error) {
            console.error('Error deploying diplomas:', error);
            if (error instanceof Error) {
                setError(`Failed to deploy diplomas: ${error.message}`);
            } else {
                setError('Failed to deploy diplomas');
            }
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="d-flex flex-column align-items-end">
            <button 
                type="button" 
                className="btn btn-secondary" 
                onClick={handleBatchDiplomasDeploy}
                disabled={loading}
            >
                {loading ? 'Deploying...' : 'Deploy diplomas'}
            </button>
            {error && <div className="text-danger mt-2">{error}</div>}
        </div>
    );
}

export default ButtonDeployDiplomas;