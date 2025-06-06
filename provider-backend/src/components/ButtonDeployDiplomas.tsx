import axios from 'axios';

function ButtonDeployDiplomas() {
    const handleBatchDiplomasDeploy = async () => {
        try {
            // generate the zip file on the server
            await axios.get('http://localhost:8080/generate-zip');

            // download the saved zip file
            const response = await axios.get('http://localhost:8080/download-zip', {
                responseType: 'blob', // the response should a binary file
            });

            // check if the response is a Blob
            if (!(response.data instanceof Blob)) {
                throw new Error('Response is not a valid Blob');
            }
            // create a download link for the zip file
            const url = window.URL.createObjectURL(response.data);
            const link = document.createElement('a');
            link.href = url;
            link.setAttribute('download', 'diplomas.zip'); // Set the file name
            document.body.appendChild(link);
            link.click();
            link.remove();

            alert('Diplomas successfully zipped and downloaded');
        } catch (error) {
            console.error('Error generating diplomas zip:', error);
            alert('Failed to generate diplomas zip');
        }
    };

	return (
        <button type="submit" className="btn btn-secondary" 
        onClick={() => handleBatchDiplomasDeploy()}
        > Generate diplomas for all </button>
    );
}

export default ButtonDeployDiplomas;