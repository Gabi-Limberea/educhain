import 'bootstrap-icons/font/bootstrap-icons.css';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import '../static/register-certification.css'
import { useState } from 'react';

function RegisterCertification() {
    const [certificationName, setCertificationName] = useState('');

    const handleChange = async(event: React.FormEvent) => {
        event.preventDefault()
    }

    const handleSubmit = async(event: React.FormEvent) => {
        event.preventDefault();
        if (certificationName.trim() !== '') {
            try {
                const response = await axios.post('http://localhost:8080/addCertification', {
                    certificationName,
                });
                console.log(response.data);
                setCertificationName(''); // Clear the input field
            } catch (error) {
                console.error('Error adding certification:', error);
            }
        }
    }
    
    return (
        <div className="container-content">
            {/* <div  className='container-content d-flex  justify-content-center py-4'> */}
                {/* form to upload students data as csv file */}
                <Form className='d-flex flex-column align-items-start inner-form' onSubmit={handleSubmit}>
                    <Form.Group className="mb-3">
                        <Form.Label className='h6'>Enter certification details </Form.Label>
                        <Form.Control type="text" placeholder="certification name" onChange={handleChange}/>
                    </Form.Group>

                    <Button variant="secondary aligns-self-start" type="submit">
                        Register
                    </Button>
                </Form>
            {/* </div> */}
        </div>
    )
} 

export default RegisterCertification;