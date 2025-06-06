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
const app = express();
const cors = require('cors');
const multer  = require('multer')
const PORT = 8080; 
const fs = require('fs');
const path = require('path');
const archiver = require('archiver');

app.use(cors());
app.use(express.json());


// set local storage for the uploaded files
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
})

const storageForMultipleFiles = multer.diskStorage({
    destination: function (req, file, cb) {
        // where is the folder located locally
        cb(null, 'C:/Users/Miha/Desktop/POLI/licenta/uploadedFilesEduChain')
    },
    filename: function (req, file, cb) {
        const extractedName = file.originalname.substring(0, file.originalname.lastIndexOf('.'));
        const fileExtension = path.extname(file.originalname); // to keep the original file extesion
        cb(null, `${extractedName}${fileExtension}`);
    }
})

  
const upload = multer({ storage: storage });
const uploadMultipleFiles = multer ({storage: storageForMultipleFiles});
  
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


// limit for uploading in bulk is 100 for now, to modify as you need
app.post('/uploadFiles', uploadMultipleFiles.array('files', 100), (req, res) =>{
    try {
         // try to access uploaded archive
         const files = req.files;

        if (!files || files.length === 0) {
            return res.status(400).json({message: 'Bad request'});
        }
 
        // console.log(`Files uploaded: ${files.map(file => file.filename).join(', ')}`);
        res.status(200).json({message: 'Files uploaded successfully', files});
    } catch (error) {
        console.log("Error uploading archive", error);
        res.status(500).json({message: 'Internal server errror'});
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

app.listen(PORT, () => console.log( `Server running on port ${PORT} `));
