// Log heap & RSS every 30s:
// setInterval(() => {
//     const m = process.memoryUsage();
//     console.log(
//       `heapUsed ${(m.heapUsed / 1024 / 1024).toFixed(1)}MB  ` +
//       `heapTotal ${(m.heapTotal / 1024 / 1024).toFixed(1)}MB  ` +
//       `rss ${(m.rss / 1024 / 1024).toFixed(1)}MB`
//     );
// }, 30_000);

const express = require('express');
const cors = require('cors');
const multer = require('multer');
const path = require('path');
const fs = require('fs');
const archiver = require('archiver');

const app = express();
const port = 8080;

// Enable CORS
app.use(cors());
app.use(express.json());

// Log all incoming requests
app.use((req, res, next) => {
    console.log(`${new Date().toISOString()} - ${req.method} ${req.url}`);
    next();
});

// Configure multer for file uploads
const storage = multer.diskStorage({
    destination: function (req, file, cb) {
        // where is the folder located locally
        cb(null, 'C:/Users/Miha/Desktop/POLI/licenta/uploadedFilesEduChain')
    },
    filename: function (req, file, cb) {
        // rename file
        const uniqueSuffix = Date.now();
        const extractedName = file.originalname.substring(0, file.originalname.lastIndexOf('.'));
        const fileExtension = path.extname(file.originalname); // to keep the original file extesion
        cb(null, `${extractedName}-${uniqueSuffix}${fileExtension}`);
    }
});

const storageForMultipleFiles = multer.diskStorage({
    destination: function (req, file, cb) {
        // where is the folder located locally
        cb(null, 'C:/Users/Miha/Desktop/POLI/licenta/uploadedFilesEduChain')
    },
    filename: function (req, file, cb) {
        // keep original filename
        cb(null, file.originalname);
    }
});

const upload = multer({ storage: storage });
const uploadMultiple = multer({ storage: storageForMultipleFiles });

// POST request to handle file uploads from providers for students' data
app.post('/uploadFile', upload.single('file'), (req, res) =>{
    try {
        // try to access uploaded file
        const file = req.file;

        if (!file) {
            return res.status(400).json({message: 'Bad request'});
        }

        console.log(`Archive uploaded: ${file.filename}`);
        res.status(200).json({message: 'File uploaded successfully', file});
    } catch (error) {
        console.error('Error uploading file:', error);
        res.status(500).json({message: 'Internal server error'});
    }
});

// Endpoint to serve files from uploadedFilesEduChain folder
app.get('/local-diplomas/:filename', (req, res) => {
    console.log('Received request for file:', req.params.filename);
    
    const filePath = path.join('C:/Users/Miha/Desktop/POLI/licenta/uploadedFilesEduChain', req.params.filename);
    console.log('Looking for file at:', filePath);

    if (!fs.existsSync(filePath)) {
        console.log('File not found at path:', filePath);
        return res.status(404).json({ error: 'File not found' });
    }

    console.log('File found, sending...');
    // Send the file
    res.sendFile(filePath, (err) => {
        if (err) {
            console.error('Error sending file:', err);
            res.status(500).json({ error: 'Error sending file' });
        } else {
            console.log('File sent successfully');
        }
    });
});

// limit for uploading in bulk is 100 for now, to modify as you need
app.post('/uploadFiles', uploadMultiple.array('files', 100), (req, res) => {
    try {
        if (!req.files || req.files.length === 0) {
            return res.status(400).json({ error: 'No files uploaded' });
        }

        const fileNames = req.files.map(file => file.originalname);
        res.json({ message: 'Files uploaded successfully', files: fileNames });
    } catch (error) {
        console.error('Error uploading files:', error);
        res.status(500).json({ error: 'Error uploading files' });
    }
});

// generate a zip file for .pdf diplomas
app.get('/generate-zip', (req, res) => {
    const folderPath = 'C:/Users/Miha/Desktop/POLI/licenta/uploadedFilesEduChain';
    const zipFilePath = path.join(__dirname, 'diplomas.zip');

    // create a zip archive
    const archive = archiver('zip', { zlib: { level: 9 } });

    // errors
    archive.on('error', (err) => {
        console.error('Error creating zip:', err);
        res.status(500).send('Error creating zip file');
    });

    // write the archive to a file on the server
    const output = fs.createWriteStream(zipFilePath);
    archive.pipe(output);

    // read the folder and add .pdf files to the archive
    fs.readdir(folderPath, (err, files) => {
        if (err) {
            console.error('Error reading folder:', err);
            res.status(500).send('Error reading folder');
            return;
        }
        console.log("ajunge sa citeasca din folder");

        files.forEach((file) => {
            if (path.extname(file) === '.pdf') {
                console.log("numele fisierului este: ", file);
                const filePath = path.join(folderPath, file);
                archive.file(filePath, { name: file });
            }
        });

        archive.finalize();
    });

    // Send a response when the zip file is successfully created
    output.on('close', () => {
        console.log(`Zip file created: ${zipFilePath}`);
        res.status(200).send('Zip file successfully created and saved on the server');
    });
});

// Endpoint to download the saved zip file
app.get('/download-zip', (req, res) => {
    const zipFilePath = path.join(__dirname, 'diplomas.zip');
    res.download(zipFilePath, 'diplomas.zip', (err) => {
        if (err) {
            console.error('Error sending zip file:', err);
            res.status(500).send('Error downloading zip file');
        }
    });
});

// POST endpoint to add a certification to a users account
app.post('/addCertification', (req, res) => {
    const { certificationName } = req.body;

    if (!certificationName) {
        return res.status(400).json({ error: 'Certification name is required' });
    }

    // Read the existing certifications from the JSON file
    fs.readFile(certificationsFilePath, 'utf8', (err, data) => {
        if (err && err.code !== 'ENOENT') {
            return res.status(500).json({ error: 'Error reading certifications file' });
        }

        const certifications = data ? JSON.parse(data) : [];

        // Add the new certification
        certifications.push(certificationName);

        // Write the updated certifications back to the file
        fs.writeFile(certificationsFilePath, JSON.stringify(certifications, null, 2), (err) => {
            if (err) {
                return res.status(500).json({ error: 'Error saving certifications' });
            }

            res.status(200).json({ message: 'Certification added successfully', certifications });
        });
    });
});

// GET endpoint to retrive users certifications
app.get('/getCertifications', (req, res) => {
    fs.readFile(certificationsFilePath, 'utf8', (err, data) => {
        if (err && err.code !== 'ENOENT') {
            return res.status(500).json({ error: 'Error reading certifications file' });
        }

        const certifications = data ? JSON.parse(data) : [];
        res.status(200).json(certifications);
    });
});

// Endpoint to serve diploma files
app.get('/diploma-files/:filename', (req, res) => {
    console.log('Received request for diploma file:', req.params.filename);
    
    const filePath = path.join('C:/Users/Miha/Desktop/POLI/licenta/uploadedFilesEduChain', req.params.filename);
    console.log('Looking for diploma file at:', filePath);

    // Check if directory exists
    const dirPath = 'C:/Users/Miha/Desktop/POLI/licenta/uploadedFilesEduChain';
    if (!fs.existsSync(dirPath)) {
        console.log('Directory does not exist:', dirPath);
        return res.status(404).json({ error: 'Directory not found' });
    }

    // List files in directory
    const files = fs.readdirSync(dirPath);
    console.log('Files in directory:', files);

    if (!fs.existsSync(filePath)) {
        console.log('Diploma file not found at path:', filePath);
        return res.status(404).json({ error: 'Diploma file not found' });
    }

    console.log('Diploma file found, sending...');
    // Send the file
    res.sendFile(filePath, (err) => {
        if (err) {
            console.error('Error sending diploma file:', err);
            res.status(500).json({ error: 'Error sending diploma file' });
        } else {
            console.log('Diploma file sent successfully');
        }
    });
});

app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
    console.log('Available endpoints:');
    console.log('- GET /diploma-files/:filename');
    // ... log other endpoints ...
});
