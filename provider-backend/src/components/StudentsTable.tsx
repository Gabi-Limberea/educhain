import React from 'react';
import { StudentsInfo } from "../Types";
import api from '../services/api';
import axios from 'axios';

interface StudentsTableProps {
    jsonWithStudents: StudentsInfo[];
    setJsonWithStudents: React.Dispatch<React.SetStateAction<StudentsInfo[]>>;
}

const StudentsTable: React.FC<StudentsTableProps> = ({ jsonWithStudents, setJsonWithStudents }) => {
    const handleGenerateDiploma = async (studentId: string) => {
        try {
            // First, get the diploma file from our local server
            const fileResponse = await axios.get(`http://localhost:8080/diploma-files/${studentId}.pdf`, {
                responseType: 'blob'
            });

            // Create a File object from the blob
            const file = new File([fileResponse.data as Blob], `${studentId}.pdf`, { type: 'application/pdf' });

            // Create FormData and append the file with the key 'file'
            const formData = new FormData();
            formData.append('file', file);

            // Send the file to the API
            const response = await api.post(`/students/${studentId}/diplomas`, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            });

            // Update localStorage with issued diplomas
            const issuedDiplomas = JSON.parse(localStorage.getItem('issuedDiplomas') || '[]');
            issuedDiplomas.push(studentId);
            localStorage.setItem('issuedDiplomas', JSON.stringify(issuedDiplomas));

            // Update the students list to reflect the issuance
            setJsonWithStudents(prevStudents => 
                prevStudents.map(student => 
                    student.studentId === studentId 
                        ? { ...student, diplomaIssued: true }
                        : student
                )
            );

            // Show success message
            alert('Diploma generated successfully!');
        } catch (error) {
            console.error('Error generating diploma:', error);
            alert('Failed to generate diploma. Please try again.');
        }
    };

    return (
        <div className="table-responsive">
            <table className="table table-striped">
                <thead>
                    <tr>
                        <th>Student ID</th>
                        <th>Wallet Address</th>
                        <th>Status</th>
                        <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                    {jsonWithStudents.map((student) => (
                        <tr key={student.studentId}>
                            <td>{student.studentId}</td>
                            <td>{student.walletAddress}</td>
                            <td>
                                {student.diplomaIssued ? (
                                    <span className="text-success">Diploma issued</span>
                                ) : student.uploadedDiploma ? (
                                    <span className="text-warning">Diploma uploaded</span>
                                ) : (
                                    <span className="text-danger">No diploma</span>
                                )}
                            </td>
                            <td>
                                {student.uploadedDiploma && !student.diplomaIssued && (
                                    <button
                                        className="btn btn-primary btn-sm"
                                        onClick={() => handleGenerateDiploma(student.studentId)}
                                    >
                                        Generate
                                    </button>
                                )}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}

export default StudentsTable;