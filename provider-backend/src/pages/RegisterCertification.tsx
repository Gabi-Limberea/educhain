import React, { useState } from 'react';
import { Container, Row, Col, Card, Button, Form } from 'react-bootstrap';
import 'bootstrap-icons/font/bootstrap-icons.css';
import axios from 'axios';

function RegisterCertification() {
    const [certificationName, setCertificationName] = useState('');

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setCertificationName(event.target.value);
    };

    const handleSubmit = async (event: React.FormEvent) => {
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
    };
    
    return (
        <div className="bg-light min-vh-100 py-5">
            <Container>
                <Row className="justify-content-center">
                    <Col md={6} lg={5}>
                        <Card className="shadow-sm border-0">
                            <Card.Header className="bg-white border-0 text-center py-4">
                                <div className="mb-3">
                                </div>
                                <h3 className="fw-bold text-dark mb-2">Register Certification</h3>
                                <p className="text-muted mb-0">Create a new certificate template</p>
                            </Card.Header>
                            <Card.Body className="p-4">
                                {/* FORM */}
                                <Form onSubmit={handleSubmit}>
                                    <Form.Group className="mb-4">
                                        <Form.Label className="fw-medium text-dark mb-2">
                                            <i className="bi bi-pencil me-2"></i>
                                            Certification Name
                                        </Form.Label>
                                        <Form.Control 
                                            type="text" 
                                            placeholder="Enter certification name (e.g., Web Development Bootcamp)" 
                                            value={certificationName}
                                            onChange={handleChange}
                                            className="py-3"
                                            style={{ fontSize: '1rem' }}
                                        />
                                        <Form.Text className="text-muted">
                                            <i className="bi bi-info-circle me-1"></i>
                                            This will be the title of your certificate template
                                        </Form.Text>
                                    </Form.Group>

                                    <div className="d-grid">
                                        <Button 
                                            variant="secondary" 
                                            type="submit" 
                                            size="lg"
                                            className="py-3"
                                        >
                                            <i className="bi bi-plus-circle me-2"></i>
                                            Register Certification
                                        </Button>
                                    </div>
                                </Form>
                            </Card.Body>
                        </Card>
                        
                        {/* Additional Info Card */}
                        <Card className="mt-4 border-0 bg-transparent">
                            <Card.Body className="p-0">
                                <div className="text-center">
                                    <h6 className="text-muted mb-3">What happens next?</h6>
                                    <Row className="g-3">
                                        <Col md={4}>
                                            <div className="text-center">
                                                <div className="bg-light rounded-circle d-inline-flex align-items-center justify-content-center mb-2" 
                                                     style={{ width: '40px', height: '40px' }}>
                                                    <i className="bi bi-1-circle text-secondary"></i>
                                                </div>
                                                <small className="text-muted d-block">Certificate template created</small>
                                            </div>
                                        </Col>
                                        <Col md={4}>
                                            <div className="text-center">
                                                <div className="bg-light rounded-circle d-inline-flex align-items-center justify-content-center mb-2" 
                                                     style={{ width: '40px', height: '40px' }}>
                                                    <i className="bi bi-2-circle text-secondary"></i>
                                                </div>
                                                <small className="text-muted d-block">Ready for student registration</small>
                                            </div>
                                        </Col>
                                        <Col md={4}>
                                            <div className="text-center">
                                                <div className="bg-light rounded-circle d-inline-flex align-items-center justify-content-center mb-2" 
                                                     style={{ width: '40px', height: '40px' }}>
                                                    <i className="bi bi-3-circle text-secondary"></i>
                                                </div>
                                                <small className="text-muted d-block">Issue verified certificates</small>
                                            </div>
                                        </Col>
                                    </Row>
                                </div>
                            </Card.Body>
                        </Card>
                    </Col>
                </Row>
            </Container>
        </div>
    );
} 

export default RegisterCertification;