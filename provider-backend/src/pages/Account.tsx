import { useEffect, useState } from 'react';
import 'bootstrap-icons/font/bootstrap-icons.css';
import '../static/account.css';
import avatarPlaceholder from '../resources/avatar-placeholder.png';
import api from '../services/api';
import { useCertifications } from '../context/CertificationContext';
import Web3 from 'web3';

interface ProviderData {
    email: string;
    organizationInfo: {
        name: string;
        address: string;
        phoneNumber: string;
        website?: string;
    };
}

interface StudentData {
    name: string;
    walletId: string;
}

function Account() {
    const [providerData, setProviderData] = useState<ProviderData | null>(null);
    const [studentData, setStudentData] = useState<StudentData | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const { certifications } = useCertifications();
    const userType = localStorage.getItem('userType');

    const connectWallet = async () => {
        if (typeof window.ethereum !== 'undefined') {
            try {
                const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
                if (userType === 'student') {
                    localStorage.setItem('studentWalletId', accounts[0]);
                    setStudentData(prev => prev ? { ...prev, walletId: accounts[0] } : null);
                }
            } catch (error) {
                console.error('Error connecting to MetaMask:', error);
                alert('Failed to connect to MetaMask');
            }
        } else {
            alert('Please install MetaMask to use this feature');
        }
    };

    useEffect(() => {
        const fetchData = async () => {
            try {
                if (userType === 'provider') {
                    const response = await api.get<ProviderData>('/account');
                    setProviderData(response.data);
                } else if (userType === 'student') {
                    const name = localStorage.getItem('studentName');
                    const walletId = localStorage.getItem('studentWalletId');
                    if (name && walletId) {
                        setStudentData({ name, walletId });
                    }
                }
                setLoading(false);
            } catch (err) {
                console.error('Error fetching data:', err);
                setError('Failed to load data');
                setLoading(false);
            }
        };

        fetchData();
    }, [userType]);

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
                <div className='profile-header-text'>
                    {userType === 'provider' 
                        ? providerData?.organizationInfo.name || 'Provider'
                        : studentData?.name || 'Student'
                    }
                </div>
            </div>

            <div className="bottom-row">
                <div className="left-container container-border">
                    <div className='user-info-top'> 
                        <h4 className='user-info-text'>User details</h4>
                        {userType === 'provider' && (
                            <button className='btn btn-secondary user-info-button'> Edit profile </button>
                        )}
                    </div>
                    <ul className='user-info-list'>
                        {userType === 'provider' ? (
                            <>
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
                            </>
                        ) : (
                            <>
                                <li>
                                    <dt>Name</dt>
                                    <dl>{studentData?.name || 'Not available'}</dl>
                                </li>
                                <li>
                                    <dt>Wallet Address</dt>
                                    <dl>{studentData?.walletId || 'Not available'}</dl>
                                </li>
                            </>
                        )}
                    </ul>
                </div>

                <div className="right-container">
                    {userType === 'provider' ? (
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
                    ) : (
                        <div className="right-top container-border">
                            <div className='user-info-top'> 
                                <h4 className='user-info-text'>Wallet Information</h4>
                            </div>
                            <div className="text-center mb-4">
                                <img 
                                    src="https://upload.wikimedia.org/wikipedia/commons/3/36/MetaMask_Fox.svg" 
                                    alt="MetaMask Logo" 
                                    style={{ width: '64px', height: '64px' }}
                                />
                            </div>
                            <div className="text-center mb-4">
                                <button 
                                    className="btn btn-primary"
                                    onClick={connectWallet}
                                >
                                    Connect MetaMask
                                </button>
                            </div>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
}

export default Account;