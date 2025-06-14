import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import NavDropdown from 'react-bootstrap/NavDropdown';
import { useLocation } from 'react-router-dom';
import 'bootstrap-icons/font/bootstrap-icons.css';
import AccountDropdown from './AccountDropdown';


function NavbarComponent() {
    const currentLocation = useLocation();

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
                <Nav.Link href="/registerStudents">Register Students</Nav.Link>
                <Nav.Link href="/registerCertification">Register Certification</Nav.Link>
                
            </Nav>

			{/* LOGIN, SIGNUP  and LOGOUT*/}
			<Nav className="ms-auto">
				{['/home', '/login'].includes(currentLocation.pathname)  && (
					<>
						<Nav.Link href="/signup" className="d-inline-flex">Signup</Nav.Link>
					</>
				)}
				{['/home', '/signup'].includes(currentLocation.pathname) && (
					<>
					<Nav.Link href="/login" className="d-inline-flex">Login</Nav.Link>
					</>
				)}
				{currentLocation.pathname === '/registerStudents' || currentLocation.pathname === '/registerCertification' && (
                    <AccountDropdown></AccountDropdown>
                    // <i className="bi bi-person-circle account-icon"></i>
                    // <Nav.Link href="/logout">Logout</Nav.Link>
				)}
                {currentLocation.pathname === '/account' && (
                    // <i className="bi bi-person-circle account-icon"></i>
                    <Nav.Link href="/logout">Logout</Nav.Link>
				)}

            </Nav>

            </Navbar.Collapse>
        </Container>
        </Navbar>
    );
}

export default NavbarComponent;