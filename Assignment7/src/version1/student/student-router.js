var express = require('express');
var router = express.Router();
const refStudent = require('./student-controller');
const objStudent = new refStudent();

//validator to validate the Student details
const refStudentValidator = require("./student-validator");
const objStudentValidator = new refStudentValidator();

//============= invoke wallet certificates ===============
router.post('/enrollAdmin', async (req, res, next) => {
    try {
        const result = await objStudent.enrollAdminCertificate();
        res.status(200).send({ message: result.message, status: result.status });

    } catch (err) {
        //error handling
        res.status(err.httpStatus || 400).send({ message: err.message, status: err.status || 0 });
    }
});

//============= invoke wallet certificates ===============
router.post('/registerUser', async (req, res, next) => {
    try {
        const result = await objStudent.registerUserCertificate();
        res.status(200).send({ message: result.message, status: result.status });

    } catch (err) {
        //error handling
        res.status(err.httpStatus || 400).send({ message: err.message, status: err.status || 0 });
    }
});

// Add student 
router.post('/add_student', objStudentValidator.addStudent, objStudentValidator.validateHandler, async (req, res, next) => {
    try {
        const result = await objStudent.studentData(req.body, 1);
        res.status(200).send({ message: result.message, status: result.status });

    } catch (err) {
        //error handling
        res.status(err.httpStatus || 400).send({ message: err.message, status: err.status || 0 });
    }
});

// Get Student
router.get('/get_student', objStudentValidator.getStudent, objStudentValidator.validateHandler, async (req, res, next) => {
    try {
        const result = await objStudent.studentData(req.query, 2);
        res.status(200).send({ message: result.message, status: result.status, data: result.data });
    } catch (err) {
        //error handling
        res.status(err.httpStatus || 400).send({ message: err.message, status: err.status || 0 });
    }
});

// Get All Students 
router.get('/get_all_student', async (req, res, next) => {
    try {
        const result = await objStudent.studentData('', 3);
        res.status(200).send({ message: result.message, status: result.status, data: result.data });

    } catch (err) {
        //error handling
        res.status(err.httpStatus || 400).send({ message: err.message, status: err.status || 0 });
    }
});


// Update student 
router.put('/update_student', objStudentValidator.updateStudent, objStudentValidator.validateHandler, async (req, res, next) => {
    try {
        const result = await objStudent.studentData(req.body, 4);
        res.status(200).send({ message: result.message, status: result.status });
    } catch (err) {
        //error handling
        res.status(err.httpStatus || 400).send({ message: err.message, status: err.status || 0 });
    }
});

// Delete student 
router.post('/delete_student', objStudentValidator.deleteStudent, objStudentValidator.validateHandler, async (req, res, next) => {
    try {
        const result = await objStudent.studentData(req.body, 5);
        res.status(200).send({ message: result.message, status: result.status });

    } catch (err) {
        //error handling
        res.status(err.httpStatus || 400).send({ message: err.message, status: err.status || 0 });
    }
});

module.exports = router;