import { StudentsInfo } from "../Types";
import axios from 'axios';

interface StudentsTableProps {
    jsonWithStudents: StudentsInfo[] | null;
};

function StudentsTable({jsonWithStudents}: StudentsTableProps) {
    // console.log("Data received by StudentsTable:", jsonWithStudents);
    if (!jsonWithStudents || jsonWithStudents.length === 0) {
        console.error('No data in studentsJSON to display');
        return (
            <div> No data available for students. Please check your .csv file and reload it. </div>
        )
    }
	return (
        <div className="table-responsive">
            <table className="table table-hover">
                <thead>
                    <tr>
                        <th scope="col" className="py-3">Student ID</th>
                        <th scope="col" className="py-3">Wallet</th>
                        <th scope="col" className="py-3">Diploma</th>
                    </tr>
                </thead>
                {/* map info received from JSON to table cells */}
                <tbody>
                    {jsonWithStudents.map((student, index) => (
                        <tr key={index}>
                            <td className="py-3">{student.studentId}</td>
                            <td className="py-3">{student.walletAddress}</td>
                            { student.studentId && (
                                <td className="py-3 "> {student.uploadedDiploma ? addGenerateButton(student.studentId) : "No added diploma"}
                                </td>
                            )}
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
	
};

const handleGenerateDiploma = async (studentId: string) => {
    try {
        // path to find student's diploma
        const filePath = `C:/Users/Miha/Desktop/POLI/licenta/uploadedFilesEduChain/${studentId}.pdf`;

        // create a File object from the file path
        const file = new File([filePath], `${studentId}.pdf`, { type: 'application/pdf' });

        const formData = new FormData();
        formData.append('file', file);

        // just for debug until POST works
        console.log(`Diploma generated for student ${studentId}:`);

        // POST request => !!! will not work until connected to the backend
        const response = await axios.post(`/students/${studentId}/diplomas`, formData, {
            headers: {
                'Content-Type': 'multipart/form-data',
            },
        });

        console.log(`Diploma generated for student ${studentId}:`, response.data);
        alert(`Diploma successfully generated for student ${studentId}`);
    } catch (error) {
        console.error(`Error generating diploma for student ${studentId}:`, error);
        alert(`Failed to generate diploma for student ${studentId}`);
    }
};


const addGenerateButton = (studentId: string) => {
    return (
        <div className="d-inline-flex flex-row w-100 ">
            <div>Added diploma</div>
            <button type="submit" 
            className="btn btn-secondary mx-3" 
            onClick={() => handleGenerateDiploma(studentId)}
            > Generate </button>        
        </div>
    )
};

export default StudentsTable;