import '../static/footer.css';

function FooterComponent() {
    return (
        <footer className="bg-body-tertiary text-black py-2">
        <div className="container mx-auto text-center copyright-text">
            <p>&copy; {new Date().getFullYear()} EduChain. All rights reserved.</p>
            <p>
                <a href="/privacy-policy" className="text-gray-400 hover:text-white">Privacy Policy</a> | 
                <a href="/terms-of-service" className="text-gray-400 hover:text-white">Terms of Service</a>
            </p>
        </div>
        </footer>
    );
}

export default FooterComponent;