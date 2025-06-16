import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import React, { useState, useEffect } from "react";
import api from '../services/api';
import Papa from 'papaparse';
import { StudentsInfo } from '../Types';
import StudentsTable from '../components/StudentsTable';
import ButtonAddBatchDiplomas from '../components/ButtonAddBatchDiplomas';
import ButtonDeployDiplomas from '../components/ButtonDeployDiplomas';

interface StudentPayload {
    id: string;
    walletAddress: string;
}

function RegisterStudents() {
    const [selectedFile, setSelectedFile] = useState<File | null>(null);
    const [jsonWithStudents, setJsonWithStudents] = useState<StudentsInfo[]>([]);
    const [submittedFile, setSubmittedFile] = useState(false);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    // Check if a student has a diploma
    const checkStudentDiploma = async (studentId: string): Promise<boolean> => {
        try {
            const response = await api.get(`/students/${studentId}/diplomas`);
            // Check if the response has any diplomas
            return Array.isArray(response.data) && response.data.length > 0;
        } catch (error) {
            console.log(`No diploma found for student ${studentId}:`, error);
            return false;
        }
    };

    // Load students from database when component mounts
    useEffect(() => {
        const fetchStudents = async () => {
            try {
                const response = await api.get<StudentPayload[]>('/students');
                
                // Get issued diplomas from localStorage
                const issuedDiplomas = JSON.parse(localStorage.getItem('issuedDiplomas') || '[]') as string[];
                
                // Check diploma status for each student
                const studentsWithDiplomas = await Promise.all(
                    response.data.map(async (student) => {
                        const hasDiploma = await checkStudentDiploma(student.id);
                        return {
                            studentId: student.id,
                            walletAddress: student.walletAddress,
                            uploadedDiploma: hasDiploma,
                            diplomaIssued: issuedDiplomas.includes(student.id)
                        };
                    })
                );

                setJsonWithStudents(studentsWithDiplomas);
                setLoading(false);
            } catch (err) {
                console.error('Error fetching students:', err);
                setError('Failed to load students');
                setLoading(false);
            }
        };

        fetchStudents();
    }, []);

    // used to read only the first line of the csv to find out the delimiter
    const readFirstLine = (file: File): Promise<string> => {
        return new Promise((resolve, reject) => {
            const CHUNK_SIZE = 2048; // 2KB should be enough to read the header of a .csv file 
            const blob = file.slice(0, CHUNK_SIZE);
            const reader = new FileReader();
      
            reader.onload = () => {
                const result = reader.result as string;
                const firstLine = result.split(/\r\n|\n/)[0];
                resolve(firstLine);
            };
      
            reader.onerror = reject;
      
            reader.readAsText(blob);
        });
    };

    function handleChange(event: React.ChangeEvent<HTMLInputElement>) {
        setSelectedFile(event.target.files ? event.target.files[0] : null);
    }

    // parse csv file with Papaparse
    const parsedCSVfile = (file: File): Promise<StudentPayload[]> => {
        return new Promise((resolve, reject) => {
            // Read the first line to determine the delimiter
            readFirstLine(file)
                .then((stringToRead) => {
                    const parsedMeta = Papa.parse(stringToRead, { preview: 1 });
                    const foundDelimiter = parsedMeta.meta.delimiter;
                    console.log("delimiter is:", foundDelimiter);
    
                    // parse the file to transform into json
                    Papa.parse<{ studentId: string; walletAddress: string }>(file, {
                        delimiter: foundDelimiter,
                        dynamicTyping: true,
                        header: true,
                        skipEmptyLines: true,
                        complete: (results) => {
                            console.log("Parsed data: ", results.data);
                            // Transform the data to match the server's expected format
                            const transformedData = results.data.map(student => ({
                                id: student.studentId,
                                walletAddress: student.walletAddress
                            }));
                            resolve(transformedData);
                        },
                        error: (error) => {
                            console.error("Error parsing the file:", error);
                            reject(error);
                        },
                    });
                })
                .catch((error) => {
                    console.error("Error reading the first line:", error);
                    reject(error);
                });
        });
    };

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();

        if (!selectedFile) {
            console.error("no file selected, try again");
            return;
        }

        try {
            // Parse the CSV file
            const parsedData = await parsedCSVfile(selectedFile);
            console.log('Parsed CSV:', parsedData);

            // Send the data to the server
            await api.post('/students', parsedData);
            console.log('Students uploaded successfully');

            // Fetch the updated list of students
            const response = await api.get<StudentPayload[]>('/students');
            console.log('Fetched students:', response.data);

            // Check diploma status for each student
            const studentsWithDiplomas = await Promise.all(
                response.data.map(async (student) => {
                    const hasDiploma = await checkStudentDiploma(student.id);
                    return {
                        studentId: student.id,
                        walletAddress: student.walletAddress,
                        uploadedDiploma: hasDiploma,
                        diplomaIssued: false // We'll set this to true only after successful deployment
                    };
                })
            );

            // Update the local state with the new data
            setJsonWithStudents(studentsWithDiplomas);
            setSubmittedFile(true);
            setError(null); // Clear any previous errors

        } catch (error) {
            console.error('Error during upload or parsing:', error);
            if (error instanceof Error) {
                setError(`Failed to upload students: ${error.message}`);
            } else {
                setError('Failed to upload students');
            }
        }
    };
    
    if (loading) {
        return <div className="text-center py-5">Loading...</div>;
    }

    if (error) {
        return <div className="text-center py-5 text-danger">{error}</div>;
    }

    return (
        <div className='d-flex flex-column'>
            <div className='d-flex justify-content-center py-4'>
                <Form className='student-upload-form d-flex flex-column align-items-start' onSubmit={handleSubmit}>
                    <Form.Group className="mb-3">
                        <Form.Label className='h5'>Select csv file</Form.Label>
                        <Form.Control type="file" accept=".csv" onChange={handleChange}/>
                        <Form.Text className="text-muted">
                            Add the .csv file with students' IDs and wallet addresses.
                        </Form.Text>
                    </Form.Group>

                    <Button variant="secondary aligns-self-start" type="submit">
                        Upload
                    </Button>
                </Form>
            </div>

            <div className='d-flex justify-content-center py-4'>
                {jsonWithStudents.length > 0 && (
                    <div>
                        <StudentsTable 
                            jsonWithStudents={jsonWithStudents}
                            setJsonWithStudents={setJsonWithStudents}
                        />
                        <div className='d-inline-flex flex-row w-100'>
                            <div className=''>
                                <ButtonAddBatchDiplomas 
                                    jsonWithStudents={jsonWithStudents} 
                                    setJsonWithStudents={setJsonWithStudents}
                                />
                            </div>
                            <div className='d-flex align-self-center ms-auto'>
                                <ButtonDeployDiplomas
                                    jsonWithStudents={jsonWithStudents}
                                    setJsonWithStudents={setJsonWithStudents}
                                />
                            </div>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
}

export default RegisterStudents;

// bg-warning <- doodles