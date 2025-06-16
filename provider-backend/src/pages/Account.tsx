import { useEffect, useState } from 'react';
import 'bootstrap-icons/font/bootstrap-icons.css';
import '../static/account.css';
import avatarPlaceholder from '../resources/avatar-placeholder.png';
import api from '../services/api';
import { useCertifications } from '../context/CertificationContext';

interface ProviderData {
    email: string;
    organizationInfo: {
        name: string;
        address: string;
        phoneNumber: string;
        website?: string;
    };
}

function Account() {
    const [providerData, setProviderData] = useState<ProviderData | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const { certifications } = useCertifications();

    useEffect(() => {
        const fetchProviderData = async () => {
            try {
                const response = await api.get<ProviderData>('/account');
                setProviderData(response.data);
                setLoading(false);
            } catch (err) {
                console.error('Error fetching provider data:', err);
                setError('Failed to load provider data');
                setLoading(false);
            }
        };

        fetchProviderData();
    }, []);

    if (loading) {
        return <div className="text-center py-5">Loading...</div>;
    }

    if (error) {
        return <div className="text-center py-5 text-danger">{error}</div>;
    }

    return (
        <div className="account-grid">
            <div className="top-container">
                <img className='profile-picture' src={avatarPlaceholder} alt="profile picture placeholder"></img>
                <div className='profile-header-text'>{providerData?.organizationInfo.name || 'Provider'}</div>
            </div>

            <div className="bottom-row">
                <div className="left-container container-border">
                    <div className='user-info-top'> 
                        <h4 className='user-info-text'>User details</h4>
                        <button className='btn btn-secondary user-info-button'> Edit profile </button>
                    </div>
                    <ul className='user-info-list'>
                        <li> 
                            <dt>Email address</dt>
                            <dl>{providerData?.email || 'Not available'}</dl> 
                        </li>
                        <li>  
                            <dt>Address</dt>
                            <dl>{providerData?.organizationInfo.address || 'Not available'}</dl>
                        </li>
                        <li>
                            <dt>Contact phone</dt>
                            <dl>{providerData?.organizationInfo.phoneNumber || 'Not available'}</dl>
                        </li>
                        {providerData?.organizationInfo.website && (
                            <li>
                                <dt>Website</dt>
                                <dl>
                                    <a href={providerData.organizationInfo.website} target="_blank" rel="noopener noreferrer">
                                        {providerData.organizationInfo.website}
                                    </a>
                                </dl>
                            </li>
                        )}
                    </ul>
                </div>

                <div className="right-container">
                    <div className="right-top container-border">
                        <div className='user-info-top'> 
                            <h4 className='user-info-text'>Certifications</h4>
                        </div>
                        <ul className='user-info-list' id='certification-list'>
                            {certifications.length === 0 ? (
                                <li className="text-muted">No certifications registered yet</li>
                            ) : (
                                certifications.map(cert => (
                                    <li key={cert.id}>
                                        <dt>{cert.name}</dt>
                                        <dl>
                                            <div>{cert.description}</div>
                                            <small className="text-muted">Registered on: {cert.date}</small>
                                        </dl>
                                    </li>
                                ))
                            )}
                        </ul>
                    </div>

                    <div className="right-bottom container-border">
                        <div className='user-info-top'> 
                            <h4 className='user-info-text'>Other info</h4>
                        </div>
                        {/* Additional information can be added here */}
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Account;