import 'bootstrap-icons/font/bootstrap-icons.css';
import { useState } from 'react';
import Form from 'react-bootstrap/Form';
import axios from 'axios';
import { StudentsInfo } from "../Types";


function ButtonAddBatchDiplomas({ jsonWithStudents, setJsonWithStudents }: 
    { jsonWithStudents: StudentsInfo[]; 
    setJsonWithStudents: React.Dispatch<React.SetStateAction<StudentsInfo[]>> }) {
    const [selectedFiles, setSelectedFiles] = useState<File[] | null>(null);

    function handleChange(event: React.ChangeEvent<HTMLInputElement>) {
        setSelectedFiles(event.target.files ? Array.from(event.target.files) : null);
    }

    const handleSubmit = async(event: React.FormEvent) => {
        event.preventDefault();

        if (!selectedFiles || selectedFiles.length === 0) {
            console.error("no files selected, try again");
            alert("No files selected. Please select files to upload.");
            return;
        } try {
            const url = 'http://localhost:8080/uploadFiles';
            const formData = new FormData();
            selectedFiles.forEach((file) => {
                formData.append('files', file);
                formData.append('fileName', file.name);
            });
            // console.log("File Name: ", selectedFiles.name);
            // formData.append('fileName', selectedFiles.name);
            const config = {
                headers: {
                    'content-type': 'multipart/form-data',
                },
            };

            // make sure that the post request was done
            // console.log('works until here')
            const response = await axios.post(url, formData, config);
            console.log('Upload response:', response.data);

            // map the diplomas to the coresponding students
            const updatedStudents = jsonWithStudents.map((student) => {
                const matchingFile = selectedFiles.find((file) => file.name.startsWith(student.studentId));
                if (matchingFile) {
                    return { ...student, uploadedDiploma: true }; // mark as having a diploma
                }
                return student;
            });
            setJsonWithStudents(updatedStudents);

            // zip archive to send to Gabi's API

        } catch (error) {
            console.error('Error during upload or parsing:', error);
        }
    }

    return (
        <Form className='student-upload-form d-flex flex-column align-items-start' onSubmit={handleSubmit}>
            <Form.Group className="mb-3">
            <Form.Label className='h5'>Select diplomas</Form.Label>
            <Form.Control type="file" multiple accept=".pdf,.jpeg,.png" onChange={handleChange}/>

            <Form.Text className="text-muted">
                Add diplomas for students in .pdf, .jpeg or .png format.
            </Form.Text>
            </Form.Group>
            <button type="submit" className="btn btn-secondary"> Upload diplomas <i className="bi bi-upload"></i> </button>
        </Form>
    )
} 

export default ButtonAddBatchDiplomas;