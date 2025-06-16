import 'bootstrap-icons/font/bootstrap-icons.css';
import { useState } from 'react';
import Form from 'react-bootstrap/Form';
import axios from 'axios';
import { StudentsInfo } from "../Types";
import api from '../services/api';

function ButtonAddBatchDiplomas({ jsonWithStudents, setJsonWithStudents }: 
    { jsonWithStudents: StudentsInfo[]; 
    setJsonWithStudents: React.Dispatch<React.SetStateAction<StudentsInfo[]>> }) {
    const [selectedFiles, setSelectedFiles] = useState<File[] | null>(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    function handleChange(event: React.ChangeEvent<HTMLInputElement>) {
        setSelectedFiles(event.target.files ? Array.from(event.target.files) : null);
        setError(null); // Clear any previous errors
    }

    const handleSubmit = async(event: React.FormEvent) => {
        event.preventDefault();
        setLoading(true);
        setError(null);

        if (!selectedFiles || selectedFiles.length === 0) {
            console.error("no files selected, try again");
            setError("No files selected. Please select files to upload.");
            setLoading(false);
            return;
        }

        try {
            // Upload files to the local server
            const url = 'http://localhost:8080/uploadFiles';
            const formData = new FormData();
            selectedFiles.forEach((file) => {
                formData.append('files', file);
                formData.append('fileName', file.name);
            });

            const config = {
                headers: {
                    'content-type': 'multipart/form-data',
                },
            };

            await axios.post(url, formData, config);
            console.log('Files uploaded successfully');

            // Update the students list to mark diplomas as uploaded
            const updatedStudents = jsonWithStudents.map((student) => {
                const matchingFile = selectedFiles.find((file) => file.name.startsWith(student.studentId));
                if (matchingFile) {
                    return { ...student, uploadedDiploma: true };
                }
                return student;
            });

            // Update the local state
            setJsonWithStudents(updatedStudents);

            // Clear the selected files
            setSelectedFiles(null);
            setLoading(false);

        } catch (error) {
            console.error('Error during upload:', error);
            if (error instanceof Error) {
                setError(`Failed to upload diplomas: ${error.message}`);
            } else {
                setError('Failed to upload diplomas');
            }
            setLoading(false);
        }
    }

    return (
        <Form className='student-upload-form d-flex flex-column align-items-start' onSubmit={handleSubmit}>
            <Form.Group className="mb-3">
                <Form.Label className='h5'>Select diplomas</Form.Label>
                <Form.Control 
                    type="file" 
                    multiple 
                    accept=".pdf,.jpeg,.png" 
                    onChange={handleChange}
                    disabled={loading}
                />
                <Form.Text className="text-muted">
                    Add diplomas for students in .pdf, .jpeg or .png format.
                </Form.Text>
            </Form.Group>
            <button 
                type="submit" 
                className="btn btn-secondary" 
                disabled={loading || !selectedFiles}
            >
                {loading ? 'Uploading...' : 'Upload diplomas'} 
                <i className="bi bi-upload"></i>
            </button>
            {error && <div className="text-danger mt-2">{error}</div>}
        </Form>
    )
} 

export default ButtonAddBatchDiplomas;