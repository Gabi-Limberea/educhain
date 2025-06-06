import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import React, { useState, useEffect } from "react";
import axios from 'axios';
import Papa from 'papaparse';
import { StudentsInfo } from '../Types';
import StudentsTable from '../components/StudentsTable';
import ButtonAddBatchDiplomas from '../components/ButtonAddBatchDiplomas';
import ButtonDeployDiplomas from '../components/ButtonDeployDiplomas';

function RegisterStudents() {
    const [selectedFile, setSelectedFile] = useState<File | null>(null);
    const [jsonWithStudents, setJsonWithStudents] = useState<StudentsInfo[]>(() => {
            const saved = sessionStorage.getItem('selectedItems');
            // console.log("Retrieved from sessionStorage:", saved);
            return saved ? JSON.parse(saved) : [];
    });
    useEffect(() => {
        if (jsonWithStudents) {
            // console.log("Saving to sessionStorage:", jsonWithStudents);
            sessionStorage.setItem('submittedFile', 'true');
            sessionStorage.setItem('uploadedDiploma', 'true')
            sessionStorage.setItem('selectedItems', JSON.stringify(jsonWithStudents));
        }
    }, [jsonWithStudents]);

    const [submittedFile, setSubmittedFile] = useState(() => {
        const saved = sessionStorage.getItem('submittedFile');
        return saved === 'true';
    });

    const [uploadedDiploma, setUploadedDipoloma] = useState(() => {
        const saved = sessionStorage.getItem('uploadedDiploma');
        return saved === 'true';
    });

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
    const parsedCSVfile = (file: File): Promise<StudentsInfo[]> => {
        return new Promise((resolve, reject) => {
            // Read the first line to determine the delimiter
            readFirstLine(file)
                .then((stringToRead) => {
                    const parsedMeta = Papa.parse(stringToRead, { preview: 1 });
                    const foundDelimiter = parsedMeta.meta.delimiter;
                    console.log("delimitatorul este:", foundDelimiter);
    
                    // parse the file to transfrom into json
                    Papa.parse<StudentsInfo>(file, {
                        delimiter: foundDelimiter, // Use the detected delimiter
                        dynamicTyping: true, // Automatically convert numeric fields
                        header: true, // Treat the first row as headers
                        skipEmptyLines: true, // Ignore empty lines in the CSV
                        complete: (results) => {
                            console.log("Parsed data: ", results.data);
                            setJsonWithStudents(results.data);
                            resolve(results.data); // Resolve with the parsed data
                        },
                        error: (error) => {
                            console.error("Error parsing the file:", error);
                            reject(error); // Reject with the error
                        },
                    });
                })
                .catch((error) => {
                    console.error("Error reading the first line:", error);
                    reject(error); // Reject if reading the first line fails
                });
        });
    }

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault()

        if (!selectedFile) {
            console.error("no file selected, try again");
            return;
        }
        try {
            const url = 'http://localhost:8080/uploadFile';
            const formData = new FormData();
            formData.append('file', selectedFile);
            // console.log("File Name: ", selectedFile.name);
            formData.append('fileName', selectedFile.name);
            const config = {
            headers: {
                'content-type': 'multipart/form-data',
            },
            };

            // make sure that the post request was done so that we can read from the file
            const response = await axios.post(url, formData, config);
            console.log('Upload response:', response.data);
      
            // parse file only after upload
            const parsedData = await parsedCSVfile(selectedFile);
            console.log('Parsed CSV after upload:', parsedData);

            setSubmittedFile(true);

        } catch (error) {
            console.error('Error during upload or parsing:', error);
        } 
    }
    
    return (
        <div className='d-flex flex-column'>
            <div  className='d-flex  justify-content-center py-4'>
                {/* form to upload students data as csv file */}
                <Form className='student-upload-form d-flex flex-column align-items-start' onSubmit={handleSubmit}>
                    <Form.Group className="mb-3">
                        <Form.Label className='h5'>Select csv file</Form.Label>
                        {/* extension .txt is only for testing purpose, !!! TO DELETE LATER */}
                        <Form.Control type="file" accept=".csv" onChange={handleChange}/>
                        <Form.Text className="text-muted">
                            Add the .csv file with students' IDs and wallet adresses.
                        </Form.Text>
                    </Form.Group>

                    <Button variant="secondary aligns-self-start" type="submit">
                        Upload
                    </Button>
                </Form>
            </div>

            <div className='d-flex justify-content-center py-4'>
                    {/* after we loaded the .csv file with the students data and parsed it to JSON ... */}
                { submittedFile && jsonWithStudents && (
                    <div>
                        <StudentsTable jsonWithStudents={jsonWithStudents}/>
                        <div className='d-inline-flex flex-row w-100'>
                            <div className=''>
                            <ButtonAddBatchDiplomas jsonWithStudents={jsonWithStudents} 
                                                    setJsonWithStudents={setJsonWithStudents}/>
                            </div>
                            <div className='d-flex align-self-center ms-auto'> <ButtonDeployDiplomas/> </div>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
}

export default RegisterStudents;

// bg-warning <- doodles