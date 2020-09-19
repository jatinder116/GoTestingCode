
const { check, validationResult } = require('express-validator');

class StudentValidator {

    //handle the validate data
    validateHandler(req, res, next) {
        let errors = validationResult(req);
        if (errors.isEmpty()) {
            next();
        } else {
            res.status(400).send({ message: "Validation Errors", status: 0, errors: errors });
        }
    };

    //handle validation for addStudent validate
    get addStudent() {
        return [
            check("id").not().isEmpty().withMessage("Id is required."),
            check("name").not().isEmpty().withMessage("Name is required.").isLength({ min: 1, max: 100 }).withMessage("username length must be from 1 to 100."),
            check("gender").not().isEmpty().withMessage("Gender is required.").isLength({ min: 1, max: 12 }).withMessage("gender length must be from 1 to 12."),
            check("city").not().isEmpty().withMessage("City is required.").isLength({ min: 1, max: 12 }).withMessage("City length must be from 1 to 12."),
        ]
    }
    //handle validation for Get Student validate
    get getStudent() {
        return [
            check("id").not().isEmpty().withMessage("Id is required."),
        ]
    }
    //handle validation for Update Student validate
    get updateStudent() {
        return [
            check("id").not().isEmpty().withMessage("Id is required."),
        ]
    }
    //handle validation for Delete Student validate
    get deleteStudent() {
        return [
            check("id").not().isEmpty().withMessage("Id is required."),
        ]
    }
}

module.exports = StudentValidator;