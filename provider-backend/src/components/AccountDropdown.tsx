import 'bootstrap-icons/font/bootstrap-icons.css';
import Dropdown from 'react-bootstrap/Dropdown';
import { useState, useEffect } from 'react';
import api from '../services/api';
import '../static/account-dropdown.css'

function AccountDropdown() {
    const [providerName, setProviderName] = useState<string>('');

    useEffect(() => {
        const fetchProviderInfo = async () => {
            try {
                const { data } = await api.get<{ organizationInfo: { name: string } }>('/account');
                setProviderName(data.organizationInfo.name);
            } catch (error) {
                console.error('Error fetching provider info:', error);
            }
        };

        fetchProviderInfo();
    }, []);

    return (
        <Dropdown align="end">
            <Dropdown.Toggle variant="secondary" id="dropdown-basic">
                <i className="bi bi-person-circle account-icon"></i>
                {providerName && <span className="ms-2">{providerName}</span>}
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