import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import { useLocation } from 'react-router-dom';
import 'bootstrap-icons/font/bootstrap-icons.css';
import AccountDropdown from './AccountDropdown';
import { useState, useEffect } from 'react';

function NavbarComponent() {
    const currentLocation = useLocation();
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [userType, setUserType] = useState<'student' | 'provider' | null>(null);

    useEffect(() => {
        const token = localStorage.getItem('authToken');
        const type = localStorage.getItem('userType') as 'student' | 'provider' | null;
        setIsLoggedIn(!!token || !!type);
        setUserType(type);
    }, []);

    return (
        <Navbar expand="lg" className="bg-body-tertiary">
            <Container fluid>
                <Navbar.Brand href="#">EduChain</Navbar.Brand>
                <Navbar.Toggle aria-controls="navbarScroll" />
                <Navbar.Collapse id="navbarScroll">
                    <Nav
                        className="me-auto my-2 my-lg-0"
                        style={{ maxHeight: '100px' }}
                        navbarScroll
                    >
                        <Nav.Link href="/home">Home</Nav.Link>  
                        {isLoggedIn && userType === 'provider' && (
                            <>
                                <Nav.Link href="/registerStudents">Register Students</Nav.Link>
                                <Nav.Link href="/registerCertification">Register Certification</Nav.Link>
                            </>
                        )}
                    </Nav>

                    <Nav className="ms-auto">
                        {!isLoggedIn && ['/home', '/login'].includes(currentLocation.pathname) && (
                            <>
                                <Nav.Link href="/signup" className="d-inline-flex">Signup</Nav.Link>
                            </>
                        )}
                        {!isLoggedIn && ['/home', '/signup'].includes(currentLocation.pathname) && (
                            <>
                                <Nav.Link href="/login" className="d-inline-flex">Login</Nav.Link>
                            </>
                        )}
                        {isLoggedIn && (
                            <AccountDropdown />
                        )}
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default NavbarComponent;