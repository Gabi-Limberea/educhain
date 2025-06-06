import 'bootstrap-icons/font/bootstrap-icons.css';
import Dropdown from 'react-bootstrap/Dropdown';
import 'bootstrap-icons/font/bootstrap-icons.css';
import '../static/account-dropdown.css'

function AccountDropdown() {

    return (
        <Dropdown align="end">
            <Dropdown.Toggle variant="secondary" id="dropdown-basic">
                <i className="bi bi-person-circle account-icon"></i>
            </Dropdown.Toggle>
    
            <Dropdown.Menu>
                <Dropdown.Item href="/account">View account</Dropdown.Item>
                <Dropdown.Item href="/logout">Logout</Dropdown.Item>
            </Dropdown.Menu>
        </Dropdown>

        // <button type="submit" className="btn btn-secondary">
        //     <i className="bi bi-person-circle account-icon"></i>
        // </button>
    )
} 

export default AccountDropdown;