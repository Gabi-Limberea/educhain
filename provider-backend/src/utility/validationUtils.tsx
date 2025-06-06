export const validateEmail = (value: string): string => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(value) ? '' : 'Invalid email address';
};

export const validatePassword = (value: string): string => {
    return value.length >= 6 ? '' : 'Password must be at least 6 characters long';
};

export const validateRepeatPassword = (password: string, repeatPassword: string): string => {
    return password === repeatPassword ? '' : 'Passwords do not match';
};

export const validateRequired = (value: string): string => {
    return value.trim() ? '' : 'This field is required';
};

export const validateName = (value: string): string => {
    return value.trim().length >= 2 ? '' : 'Name must be at least 2 characters long';
};

export const validateWalletAddress = (value: string): string => {
    // Basic wallet address validation (Ethereum format)
    const walletRegex = /^0x[a-fA-F0-9]{40}$/;
    return walletRegex.test(value) ? '' : 'Invalid wallet address format';
};

export const validatePhone = (value: string): string => {
    const phoneRegex = /^[\+]?[1-9][\d]{0,15}$/;
    return phoneRegex.test(value.replace(/[\s\-\(\)]/g, '')) ? '' : 'Invalid phone number';
};

export const validateWebsite = (value: string): string => {
    if (!value.trim()) return ''; // Optional field
    const urlRegex = /^https?:\/\/.+/;
    return urlRegex.test(value) ? '' : 'Invalid website URL (must start with http:// or https://)';
};