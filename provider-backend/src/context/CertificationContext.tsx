import React, { createContext, useContext, useState, useEffect } from 'react';

interface Certification {
    id: string;
    name: string;
    description: string;
    date: string;
}

interface CertificationContextType {
    certifications: Certification[];
    addCertification: (certification: Omit<Certification, 'id'>) => void;
    removeCertification: (id: string) => void;
}

const CertificationContext = createContext<CertificationContextType | undefined>(undefined);

export function CertificationProvider({ children }: { children: React.ReactNode }) {
    const [certifications, setCertifications] = useState<Certification[]>(() => {
        const saved = localStorage.getItem('certifications');
        return saved ? JSON.parse(saved) : [];
    });

    useEffect(() => {
        localStorage.setItem('certifications', JSON.stringify(certifications));
    }, [certifications]);

    const addCertification = (certification: Omit<Certification, 'id'>) => {
        const newCertification = {
            ...certification,
            id: Date.now().toString(), // Generate a unique ID
        };
        setCertifications(prev => [...prev, newCertification]);
    };

    const removeCertification = (id: string) => {
        setCertifications(prev => prev.filter(cert => cert.id !== id));
    };

    return (
        <CertificationContext.Provider value={{ certifications, addCertification, removeCertification }}>
            {children}
        </CertificationContext.Provider>
    );
}

export function useCertifications() {
    const context = useContext(CertificationContext);
    if (context === undefined) {
        throw new Error('useCertifications must be used within a CertificationProvider');
    }
    return context;
} 